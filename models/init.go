package models

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"git.9d77v.me/9d77v/pdc/utils"
	"github.com/9d77v/go-lib/clients/config"
	"github.com/jinzhu/gorm"
	minio "github.com/minio/minio-go/v6"

	//postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//环境变量
var (
	DEBUG = utils.GetEnvBool("DEBUG", true)

	DBHost     = utils.GetEnvStr("POSTGRES_HOST", "127.0.0.1")
	DBPort     = utils.GetEnvInt("POSTGRES_PORT", 5432)
	DBUser     = utils.GetEnvStr("POSTGRES_USER", "postgres")
	DBPassword = utils.GetEnvStr("POSTGRES_PASSWORD", "123456")
	DBName     = utils.GetEnvStr("POSTGRES_NAME", "pdc")

	MinioAddress         = utils.GetEnvStr("MINIO_ADDRESS", "127.0.0.1:9500")
	MinioAccessKeyID     = utils.GetEnvStr("MINIO_ACCESS_KEY", "minio")
	MinioSecretAccessKey = utils.GetEnvStr("MINIO_SECRET_KEY", "minio123")
	MinioUseSSL          = utils.GetEnvBool("MINIO_USE_SSL", false)
)

var (
	//Gorm global  orm
	Gorm *gorm.DB
	//MinioClient S3 OSS
	MinioClient *minio.Client
)

func init() {
	initDB()
	initMinio()
}

func initDB() {
	dbConfig := &config.DBConfig{
		Driver:       "postgres",
		Host:         DBHost,
		Port:         uint(DBPort),
		User:         DBUser,
		Password:     DBPassword,
		Name:         DBName,
		MaxIdleConns: 10,
		MaxOpenConns: 100,
		EnableLog:    DEBUG,
	}
	var err error
	Gorm, err = NewClient(dbConfig)
	if err != nil {
		log.Printf("Could not initialize gorm: %s", err.Error())
	}
	log.Println(Gorm, err)
	Gorm.AutoMigrate(
		&Media{},
		&Episode{},
	)
}

func initMinio() {
	endpoint := MinioAddress
	accessKeyID := MinioAccessKeyID
	secretAccessKey := MinioSecretAccessKey
	useSSL := MinioUseSSL

	var err error
	MinioClient, err = minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

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
