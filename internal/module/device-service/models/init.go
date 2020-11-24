package models

import (
	"log"

	"github.com/9d77v/pdc/internal/db"
)

func init() {
	err := db.Gorm.AutoMigrate(
		&DeviceModel{},
		&TelemetryModel{},
		&AttributeModel{},
		&Device{},
		&Attribute{},
		&Telemetry{},
		&DeviceDashboard{},
		&DeviceDashboardTelemetry{},
		&DeviceDashboardCamera{},
	)
	if err != nil {
		log.Println("auto migrate error:", err)
	}
}
