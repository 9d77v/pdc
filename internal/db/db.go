package db

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/9d77v/go-lib/clients/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

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
			TablePrefix:   TablePrefix + "_", // table name prefix, table for `User` would be `t_users`
			SingularTable: true,
		},
	}
	if DEBUG {
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
