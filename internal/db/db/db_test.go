package db

import (
	"fmt"
	"log"
	"testing"

	"github.com/9d77v/go-lib/clients/config"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_createDatabaseIfNotExist(t *testing.T) {
	type args struct {
		config *config.DBConfig
		dbName string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test db not exist", args{&config.DBConfig{
			Driver:       "postgres",
			Host:         "domain.local",
			Port:         5432,
			User:         "postgres",
			Password:     "123456",
			Name:         "postgres",
			MaxIdleConns: 10,
			MaxOpenConns: 100,
			EnableLog:    true,
		}, "pdc_test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := GetDB(tt.args.config)
			tt.args.config.Name = tt.args.dbName
			assert.True(t, databaseNotExist(db, tt.args.config))
			createDatabaseIfNotExist(tt.args.config)
			assert.False(t, databaseNotExist(db, tt.args.config))
			deleteDB(db, tt.args.dbName)
		})
	}
}

func deleteDB(db *gorm.DB, dbName string) {
	sql := fmt.Sprintf("DROP DATABASE %s;", dbName)
	err := db.Exec(sql).Error
	if err != nil {
		log.Println("delete database failed:", err)
	}
}
