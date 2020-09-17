package clickhouse

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/9d77v/pdc/utils"
	"github.com/ClickHouse/clickhouse-go"
)

var (
	//Client Clickhouse Client
	Client *sql.DB
	//CHAddr Clickhouse Server Address
	CHAddr = utils.GetEnvStr("CLICKHOUSE_ADDR", "domain.local:9001")
	//DEBUG ..
	DEBUG = utils.GetEnvBool("DEBUG", true)
)

func init() {
	debug := "false"
	if DEBUG {
		debug = "true"
	}
	dbInit, err := sql.Open("clickhouse", fmt.Sprintf("tcp://%s?debug=%s", CHAddr, debug))
	if err != nil {
		log.Fatal(err)
	}
	defer dbInit.Close()
	if err := dbInit.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return
	}
	dbname := "pdc"
	_, err = dbInit.Exec(fmt.Sprintf("CREATE DATABASE  IF NOT EXISTS %s", dbname))
	if err != nil {
		log.Println(err)
	}
	db, err := sql.Open("clickhouse", fmt.Sprintf("tcp://%s?debug=%s&database=%s", CHAddr, debug, "pdc"))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return
	}

	_, err = db.Exec(`
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
	_, err = db.Exec(`
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
	Client = db
}
