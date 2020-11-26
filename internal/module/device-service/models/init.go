package models

import (
	"log"

	"github.com/9d77v/pdc/internal/db/db"
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
		&CameraTimeLapseVideo{},
	)
	if err != nil {
		log.Println("auto migrate error:", err)
	}
}
