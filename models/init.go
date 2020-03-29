package models

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/9d77v/go-lib/clients/config"
	"github.com/jinzhu/gorm"

	//postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	//Gorm global  orm
	Gorm *gorm.DB
)

func init() {
	dbConfig := &config.DBConfig{
		Driver:       "postgres",
		Host:         "192.168.1.234",
		Port:         5432,
		User:         "postgres",
		Password:     "123456",
		Name:         "pdc",
		MaxIdleConns: 10,
		MaxOpenConns: 100,
		EnableLog:    true,
	}
	dbConfig.Driver = "postgres"
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
