package clickhouse

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/ClickHouse/clickhouse-go"

	"github.com/9d77v/pdc/internal/consts"
	"github.com/9d77v/pdc/internal/utils"
)

var (
	client *sql.DB
	chAddr = utils.GetEnvStr("CLICKHOUSE_ADDR", "domain.local:9001")
	once   sync.Once
)

//GetDB get clickhouse connection
func GetDB() *sql.DB {
	once.Do(func() {
		client = initClient()
	})
	return client
}

func initClient() *sql.DB {
	debug := "false"
	if consts.DEBUG {
		debug = "true"
	}
	dbInit, err := sql.Open("clickhouse", fmt.Sprintf("tcp://%s?debug=%s", chAddr, debug))
	if err != nil {
		log.Fatal(err)
	}
	defer dbInit.Close()
	if err := dbInit.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			log.Panicf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			log.Panicln(err)
		}
		return nil
	}
	dbname := "pdc"
	_, err = dbInit.Exec(fmt.Sprintf("CREATE DATABASE  IF NOT EXISTS %s", dbname))
	if err != nil {
		log.Println(err)
	}
	db, err := sql.Open("clickhouse", fmt.Sprintf("tcp://%s?debug=%s&database=%s", chAddr, debug, "pdc"))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			log.Panicf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			log.Panicln(err)
		}
		return nil
	}

	return db
}

//CreateTables create tables
func CreateTables() {
	_, err := GetDB().Exec(`
		CREATE TABLE IF NOT EXISTS device_telemetry (
			device_id UInt32,
			telemetry_id UInt32,
			action_time  DateTime CODEC(DoubleDelta),
			action_time_nanos UInt32,
			value        Float64 CODEC(Gorilla),
			created_at   DateTime CODEC(DoubleDelta),
			created_at_nanos UInt32
		) engine=MergeTree()
		ORDER BY (device_id,telemetry_id,action_time)
		PARTITION BY (device_id)
	`)
	if err != nil {
		log.Panicln("create table error:", err)
	}
	_, err = GetDB().Exec(`
		CREATE TABLE IF NOT EXISTS device_health (
			device_id UInt32,
			action_time  DateTime CODEC(DoubleDelta),
			action_time_nanos UInt32,
			value        UInt32,
			created_at   DateTime CODEC(DoubleDelta),
			created_at_nanos UInt32
		) engine=MergeTree()
		ORDER BY (device_id,action_time)
		PARTITION BY (device_id)
	`)
	if err != nil {
		log.Panicln("create table error:", err)
	}
}
