package db

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/9d77v/go-lib/clients/config"
	"github.com/9d77v/pdc/internal/consts"
	"github.com/9d77v/pdc/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

//环境变量
var (
	dbHost     = utils.GetEnvStr("DB_HOST", "domain.local")
	dbPort     = utils.GetEnvInt("DB_PORT", 5432)
	dbUser     = utils.GetEnvStr("DB_USER", "postgres")
	dbPassword = utils.GetEnvStr("DB_PASSWORD", "123456")
	dbName     = utils.GetEnvStr("DB_NAME", "pdc")
)

var (
	client *gorm.DB
	once   sync.Once
)

//TablePrefix 获取表前缀
func TablePrefix() string {
	return dbName + "_"
}

//GetDB get db connection
func GetDB() *gorm.DB {
	once.Do(func() {
		client = initClient()
	})
	return client
}

func initClient() *gorm.DB {
	dbConfig := &config.DBConfig{
		Driver:       "postgres",
		Host:         dbHost,
		Port:         uint(dbPort),
		User:         dbUser,
		Password:     dbPassword,
		Name:         dbName,
		MaxIdleConns: 10,
		MaxOpenConns: 100,
		EnableLog:    consts.DEBUG,
	}
	var err error
	db, err := newClient(dbConfig)
	if err != nil {
		log.Printf("Could not initialize gorm: %s", err.Error())
	}
	return db
}

func newClient(config *config.DBConfig) (*gorm.DB, error) {
	if config == nil {
		return nil, errors.New("db config is not exist")
	}
	//support postgres
	if config.Driver != "postgres" {
		return nil, errors.New("unsupport driver,now only support postgres")
	}
	//auto create database
	dsnInit := fmt.Sprintf("host=%s port=%d user=%s sslmode=disable password=%s",
		config.Host, config.Port, config.User, config.Password)
	dbInit, err := gorm.Open(postgres.Open(dsnInit), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	initSQL := fmt.Sprintf("CREATE DATABASE \"%s\" WITH  OWNER =%s ENCODING = 'UTF8' CONNECTION LIMIT=-1;",
		config.Name, config.User)
	err = dbInit.Exec(initSQL).Error
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		return nil, err
	}
	sqlDBInit, err := dbInit.DB()
	defer sqlDBInit.Close()

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		config.Host, config.Port, config.User, config.Name, config.Password)
	//global database connection
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   TablePrefix(), // table name prefix, table for `User` would be `t_users`
			SingularTable: true,
		},
	}
	if consts.DEBUG {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gormConfig.DisableForeignKeyConstraintWhenMigrating = true
	}
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(int(config.MaxIdleConns))
	sqlDB.SetMaxOpenConns(int(config.MaxOpenConns))
	return db, err
}
