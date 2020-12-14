package clickhouse

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/9d77v/go-lib/clients/config"
	"github.com/9d77v/pdc/internal/consts"
	"github.com/9d77v/pdc/internal/utils"
)

var (
	client        *gorm.DB
	dbHost        = utils.GetEnvStr("CLICKHOUSE_HOST", "domain.local")
	dbPort        = utils.GetEnvInt("CLICKHOUSE_PORT", 9001)
	dbName        = utils.GetEnvStr("CLICKHOUSE_NAME", "pdc")
	dbTablePrefix = utils.GetEnvStr("CLICKHOUSE_TABLE_PREFIX", "pdc")
	once          sync.Once
)

//TablePrefix 获取表前缀
func TablePrefix() string {
	return dbTablePrefix + "_"
}

//GetDB get db connection
func GetDB(config ...*config.DBConfig) *gorm.DB {
	once.Do(func() {
		dbConfig := defualtConfig()
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

func defualtConfig() *config.DBConfig {
	return &config.DBConfig{
		Driver:       "clickhouse",
		Host:         dbHost,
		Port:         uint(dbPort),
		User:         "",
		Password:     "",
		Name:         dbName,
		MaxIdleConns: 10,
		MaxOpenConns: 100,
		EnableLog:    consts.DEBUG,
	}
}

func createDatabaseIfNotExist(config *config.DBConfig) {
	dsn := fmt.Sprintf("tcp://%s:%d?debug=true", config.Host, config.Port)
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("connect to postgres failed:", err)
	}
	sql := fmt.Sprintf("CREATE DATABASE  IF NOT EXISTS %s", config.Name)
	err = db.Exec(sql).Error
	if err != nil {
		log.Println("create db failed:", err)
	} else {
		log.Printf("create db '%s' succeed\n", config.Name)
	}
	dbInit, err := db.DB()
	dbInit.Close()
}

func newClient(config *config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("tcp://%s:%d?debug=%t&database=%s&read_timeout=10&write_timeout=20",
		config.Host, config.Port, consts.DEBUG, config.Name)
	gormConfig := &gorm.Config{
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
	db, err := gorm.Open(clickhouse.New(clickhouse.Config{
		DSN:                       dsn,
		DisableDatetimePrecision:  false,    // disable datetime64 precision, not supported before clickhouse 20.4
		DontSupportRenameColumn:   true,     // rename column not supported before clickhouse 20.4
		SkipInitializeWithVersion: false,    // smart configure based on used version
		DefaultGranularity:        3,        // 1 granule = 8192 rows
		DefaultCompression:        "LZ4",    // default compression algorithm. LZ4 is lossless
		DefaultIndexType:          "minmax", // index stores extremes of the expression
	}), gormConfig)
	return db, err
}

func databaseNotExist(db *gorm.DB, config *config.DBConfig) bool {
	var total int64
	err := db.Raw("select count(1) from system.databases where name = ?;", config.Name).Scan(&total).Error
	if err != nil {
		log.Println("check db failed", err)
	}
	return total == 0
}
