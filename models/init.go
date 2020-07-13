package models

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/9d77v/go-lib/clients/config"
	"github.com/9d77v/pdc/utils"
	redis "github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	minio "github.com/minio/minio-go/v6"
	"golang.org/x/crypto/bcrypt"

	//postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	dbPrefix   = utils.GetEnvStr("DB_PREFIX", "pdc")

	JWTtAccessSecret = utils.GetEnvStr("JWT_ACCESS_SECRET", "JWT_ACCESS_SECRET")
	JWTRefreshSecret = utils.GetEnvStr("JWT_REFRESH_SECRET", "JWT_REFRESH_SECRET")
	JWTIssuer        = utils.GetEnvStr("JWT_ISSUER", "domain.local")

	minioAddress         = utils.GetEnvStr("MINIO_ADDRESS", "oss.domain.local:9000")
	minioAccessKeyID     = utils.GetEnvStr("MINIO_ACCESS_KEY", "minio")
	minioSecretAccessKey = utils.GetEnvStr("MINIO_SECRET_KEY", "minio123")
	minioUseSSL          = utils.GetEnvBool("MINIO_USE_SSL", false)

	OssPrefix = utils.GetEnvStr("OSS_PREFIX", "http://oss.domain.local:9000")

	redisAddress  = utils.GetEnvStr("REDIS_ADDRESS", "domain.local:6379")
	redisPassword = utils.GetEnvStr("REDIS_PASSWORD", "")
)

var (
	//Gorm global  orm
	Gorm *gorm.DB
	//MinioClient S3 OSS
	MinioClient *minio.Client
	//RedisClient ..
	RedisClient *redis.Client
)

//角色
const (
	RoleAdmin   = 1
	RoleManager = 2
	RoleUser    = 3
)

//redis前缀
const (
	PrefixUser = "USER"
)

func init() {
	initDB()
	initDBData()
	initMinio()
	initRedis()
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
		&User{},
		&Thing{},
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
	endpoint := minioAddress
	accessKeyID := minioAccessKeyID
	secretAccessKey := minioSecretAccessKey
	useSSL := minioUseSSL

	var err error
	MinioClient, err = minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	preCreatedBuckets := []string{"image", "video", "vtt"}
	location := "us-east-1"
	for _, bucketName := range preCreatedBuckets {
		err = MinioClient.MakeBucket(bucketName, location)
		if err != nil {
			exists, errBucketExists := MinioClient.BucketExists(bucketName)
			if errBucketExists == nil && exists {
				log.Printf("We already own %s\n", bucketName)
			} else {
				log.Fatalln(err)
			}
		} else {
			log.Printf("Successfully created %s\n", bucketName)
		}
		//mc  policy  set  download  minio/mybucket
		policy := `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS": 
		["*"]},"Action":["s3:GetBucketLocation","s3:ListBucket"],"Resource": 
		["arn:aws:s3:::` + bucketName + `"]},{"Effect":"Allow","Principal":{"AWS":["*"]},"Action": 
		["s3:GetObject"],"Resource":["arn:aws:s3:::` + bucketName + `/*"]}]}`
		err := MinioClient.SetBucketPolicy(bucketName, policy)
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
		return fmt.Sprintf("%s_%s", dbPrefix, defaultTableName)
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
