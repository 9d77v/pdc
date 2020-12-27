package chmodels

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/9d77v/go-lib/clients/config"
	"github.com/9d77v/pdc/internal/db/clickhouse"
	"github.com/9d77v/pdc/internal/module/base"
)

var migrateTables = []interface{}{
	&DeviceHealth{},
	&DeviceTelemetry{},
}

func TestMain(m *testing.M) {
	initDB()
	m.Run()
	clean()
}

func initDB() {
	config := &config.DBConfig{
		Driver:       "clickhouse",
		Host:         "domain.local",
		Port:         9001,
		User:         "",
		Password:     "",
		Name:         "pdc_test",
		MaxIdleConns: 10,
		MaxOpenConns: 100,
		EnableLog:    false,
	}
	err := clickhouse.GetDB(config).AutoMigrate(migrateTables...)
	if err != nil {
		fmt.Println("auto migrate failed:", err)
	}
}

func clean() {
	err := clickhouse.GetDB().Exec("ALTER TABLE pdc_device_health DELETE WHERE 1=1").Error
	if err != nil {
		log.Println("error:", err)
	}
	err = clickhouse.GetDB().Exec("ALTER TABLE pdc_device_telemetry DELETE WHERE 1=1").Error
	if err != nil {
		log.Println("error:", err)
	}
}

func TestDeviceHealth_IsUnhealthy(t *testing.T) {
	m := NewDeviceHealth()
	m.Create(&DeviceHealth{
		DeviceID:   1,
		ActionTime: time.Now(),
		Value:      1,
	})
	type fields struct {
		Model           base.Model
		ActionTime      time.Time
		ActionTimeNanos uint32
		DeviceID        uint32
		Value           float64
		CreatedAt       time.Time
		CreatedAtNanos  uint32
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"test healthy", fields{DeviceID: 1}, false},
		{"test unhealthy", fields{DeviceID: 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DeviceHealth{
				Model:      tt.fields.Model,
				ActionTime: tt.fields.ActionTime,
				DeviceID:   tt.fields.DeviceID,
				Value:      tt.fields.Value,
				CreatedAt:  tt.fields.CreatedAt,
			}
			m.SetDB(clickhouse.GetDB())
			if got := m.IsUnhealthy(); got != tt.want {
				t.Errorf("DeviceHealth.IsUnhealthy() = %v, want %v", got, tt.want)
			}
		})
	}
}
