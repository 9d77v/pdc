package models

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/9d77v/go-lib/clients/config"
	"github.com/9d77v/pdc/utils"
	redis "github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"golang.org/x/crypto/bcrypt"

	//postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"

	stan "github.com/nats-io/stan.go"
)

//环境变量
var (
	DEBUG = utils.GetEnvBool("DEBUG", true)

	ownerName     = utils.GetEnvStr("ADMIN_NAME", "admin")
	ownerPassword = utils.GetEnvStr("ADMIN_PASSWORD", "123456")

	dbHost     = utils.GetEnvStr("DB_HOST", "domain.local")
	dbPort     = utils.GetEnvInt("DB_PORT", 5432)
	dbUser     = utils.GetEnvStr("DB_USER", "postgres")
	dbPassword = utils.GetEnvStr("DB_PASSWORD", "123456")
	dbName     = utils.GetEnvStr("DB_NAME", "pdc")
	DBPrefix   = utils.GetEnvStr("DB_PREFIX", "pdc")

	JWTtAccessSecret = utils.GetEnvStr("JWT_ACCESS_SECRET", "JWT_ACCESS_SECRET")
	JWTRefreshSecret = utils.GetEnvStr("JWT_REFRESH_SECRET", "JWT_REFRESH_SECRET")
	JWTIssuer        = utils.GetEnvStr("JWT_ISSUER", "domain.local")

	minioAddress         = utils.GetEnvStr("MINIO_ADDRESS", "oss.domain.local:9000")
	secureMinioAddress   = utils.GetEnvStr("SECURE_MINIO_ADDRESS", "oss.domain.local")
	minioAccessKeyID     = utils.GetEnvStr("MINIO_ACCESS_KEY", "minio")
	minioSecretAccessKey = utils.GetEnvStr("MINIO_SECRET_KEY", "minio123")

	OssPrefix       = ""
	SecureOssPrerix = ""

	redisAddress  = utils.GetEnvStr("REDIS_ADDRESS", "domain.local:6379")
	redisPassword = utils.GetEnvStr("REDIS_PASSWORD", "")
)

var (
	//Gorm global  orm
	Gorm *gorm.DB
	//MinioClient S3 OSS by http
	MinioClient *minio.Client
	//SecureMinioClient S3 OSS by https
	SecureMinioClient *minio.Client
	//RedisClient ..
	RedisClient *redis.Client
	//NatsClient ..
	NatsClient stan.Conn
)

//角色
const (
	RoleAdmin   = 1
	RoleManager = 2
	RoleUser    = 3
	RoleGuest   = 4
)

//redis前缀
const (
	PrefixUser = "USER"
)

//mq constants
const (
	SubjectVideo = "video"
	GroupVideo   = "video-group"
)

func init() {
	initDB()
	initDBData()
	initMinio()
	initRedis()
	initNatsStreaming()
}

func initDB() {
	dbConfig := &config.DBConfig{
		Driver:       "postgres",
		Host:         dbHost,
		Port:         uint(dbPort),
		User:         dbUser,
		Password:     dbPassword,
		Name:         dbName,
		MaxIdleConns: 10,
		MaxOpenConns: 100,
		EnableLog:    DEBUG,
	}
	var err error
	Gorm, err = NewClient(dbConfig)
	if err != nil {
		log.Printf("Could not initialize gorm: %s", err.Error())
	}
	Gorm.AutoMigrate(
		&Video{},
		&Episode{},
		&VideoSeries{},
		&VideoSeriesItem{},
		&User{},
		&Thing{},
		&History{},
	)
}

func initDBData() {
	var total int64
	err := Gorm.Model(&User{}).Count(&total).Error
	if err != nil {
		log.Panicf("Get User total failed:%v/n", err)
	}
	if total == 0 {
		bytes, err := bcrypt.GenerateFromPassword([]byte(ownerPassword), 12)
		if err != nil {
			log.Panicf("generate password failed:%v/n", err)
		}
		user := &User{
			Name:     ownerName,
			Password: string(bytes),
			RoleID:   1,
		}
		err = Gorm.Create(user).Error
		if err != nil {
			log.Panicf("create owner failed:%v/n", err)
		}
	}
}

func initMinio() {
	accessKeyID := minioAccessKeyID
	secretAccessKey := minioSecretAccessKey

	OssPrefix = fmt.Sprintf("http://%s", minioAddress)
	SecureOssPrerix = fmt.Sprintf("https://%s", secureMinioAddress)
	var err error
	MinioClient, err = minio.New(minioAddress, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Fatalln(err)
	}
	SecureMinioClient, err = minio.New(secureMinioAddress, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: true,
	})
	preCreatedBuckets := []string{"image", "video", "vtt"}
	location := "us-east-1"
	for _, bucketName := range preCreatedBuckets {
		err = MinioClient.MakeBucket(context.Background(), bucketName,
			minio.MakeBucketOptions{Region: location})
		if err != nil {
			exists, errBucketExists := MinioClient.BucketExists(context.Background(), bucketName)
			if errBucketExists == nil && exists {
				log.Printf("We already own %s\n", bucketName)
			} else {
				log.Fatalln(err)
			}
		} else {
			log.Printf("Successfully created %s\n", bucketName)
		}
		//mc  policy  set  download  minio/mybucket
		policy := `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action": 
		["s3:GetObject"],"Resource":["arn:aws:s3:::` + bucketName + `/*"]}]}`
		err := MinioClient.SetBucketPolicy(context.Background(), bucketName, policy)
		if err != nil {
			log.Printf("Set bucket:%s policy faield:%v\n", bucketName, err)
		}
	}
}

func initRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       0, // use default DB
	})
}

//NewClient gorm client
func NewClient(config *config.DBConfig) (*gorm.DB, error) {
	if config == nil {
		return nil, errors.New("db config is not exist")
	}
	//support postgres
	if config.Driver != "postgres" {
		return nil, errors.New("unsupport driver,now only support postgres")
	}
	//auto create database
	dbURL := fmt.Sprintf("host=%s port=%d user=%s sslmode=disable password=%s",
		config.Host, config.Port, config.User, config.Password)
	dbInit, err := gorm.Open(config.Driver, dbURL)
	if err != nil {
		return nil, err
	}
	defer dbInit.Close()
	initSQL := fmt.Sprintf("CREATE DATABASE \"%s\" WITH  OWNER =%s ENCODING = 'UTF8' CONNECTION LIMIT=-1;",
		config.Name, config.User)
	err = dbInit.Exec(initSQL).Error
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return nil, err
	}
	dbWithNameURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		config.Host, config.Port, config.User, config.Name, config.Password)
	//global database connection
	db, err := gorm.Open(config.Driver, dbWithNameURL)
	if err != nil {
		return nil, err
	}

	//设置表名称的前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return fmt.Sprintf("%s_%s", DBPrefix, defaultTableName)
	}

	db.SingularTable(true)
	db.LogMode(config.EnableLog)
	useHstoreSQL := fmt.Sprintf("CREATE EXTENSION hstore;")
	err = db.Exec(useHstoreSQL).Error
	if err != nil {
		log.Println("create extension hstore failed:", err.Error())
	}
	db.DB().SetMaxIdleConns(int(config.MaxIdleConns))
	db.DB().SetMaxOpenConns(int(config.MaxOpenConns))
	return db, err
}

func initNatsStreaming() {
	NatsClient, _ = stan.Connect("test-cluster",
		fmt.Sprintf("client-%d", time.Now().Unix()),
		stan.Pings(10, 5),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
}
