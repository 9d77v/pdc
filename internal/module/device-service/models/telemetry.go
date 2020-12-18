package models

import (
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
)

//Telemetry 遥测
type Telemetry struct {
	base.DefaultModel
	DeviceID         uint `gorm:"uniqueIndex:telemetry_uix"`
	Device           Device
	TelemetryModelID uint `gorm:"uniqueIndex:telemetry_uix"`
	TelemetryModel   TelemetryModel
}

//NewTelemetry ..
func NewTelemetry() *Telemetry {
	m := &Telemetry{}
	m.SetDB(db.GetDB())
	return m
}

//TableName ..
func (m *Telemetry) TableName() string {
	return db.TablePrefix() + "telemetry"
}
