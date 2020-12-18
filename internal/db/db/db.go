package db

import (
	"fmt"
	"log"
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
	dbHost        = utils.GetEnvStr("DB_HOST", "domain.local")
	dbPort        = utils.GetEnvInt("DB_PORT", 5432)
	dbUser        = utils.GetEnvStr("DB_USER", "postgres")
	dbPassword    = utils.GetEnvStr("DB_PASSWORD", "123456")
	dbName        = utils.GetEnvStr("DB_NAME", "pdc")
	dbTablePrefix = utils.GetEnvStr("DB_TABLE_PREFIX", "pdc")
)

var (
	client *gorm.DB
	once   sync.Once
)

//TablePrefix 获取表前缀
func TablePrefix() string {
	return dbTablePrefix + "_"
}

//GetDB get db connection
func GetDB(config ...*config.DBConfig) *gorm.DB {
	once.Do(func() {
		dbConfig := defaultConfig()
		if config != nil && len(config) == 1 {
			dbConfig = config[0]
		}
		createDatabaseIfNotExist(dbConfig)
		var err error
		client, err = newClient(dbConfig)
		if err != nil {
			log.Panicf("Could not initialize gorm: %s\n", err.Error())
		}
	})
	return client
}

func defaultConfig() *config.DBConfig {
	return &config.DBConfig{
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
}

func createDatabaseIfNotExist(config *config.DBConfig) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s sslmode=disable password=%s",
		config.Host, config.Port, config.User, config.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("connect to postgres failed:", err)
	}
	if databaseNotExist(db, config) {
		createDatabase(db, config)
	}
	sqlDBInit, err := db.DB()
	sqlDBInit.Close()
}

func databaseNotExist(db *gorm.DB, config *config.DBConfig) bool {
	var total int64
	err := db.Raw("SELECT 1 FROM pg_database WHERE datname = ?", config.Name).Scan(&total).Error
	if err != nil {
		log.Println("check db failed", err)
	}
	return total == 0
}

func createDatabase(db *gorm.DB, config *config.DBConfig) {
	initSQL := fmt.Sprintf("CREATE DATABASE \"%s\" WITH  OWNER =%s ENCODING = 'UTF8' CONNECTION LIMIT=-1;",
		config.Name, config.User)
	err := db.Exec(initSQL).Error
	if err != nil {
		log.Println("create db failed:", err)
	} else {
		log.Printf("create db '%s' succeed\n", config.Name)
	}
}

func newClient(config *config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		config.Host, config.Port, config.User, config.Name, config.Password)
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   TablePrefix(),
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
