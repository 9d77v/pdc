package models

import (
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
)

//DeviceDashboardTelemetry 仪表盘遥测
type DeviceDashboardTelemetry struct {
	base.DefaultModel
	DeviceDashboardID uint
	TelemetryID       uint
	Telemetry         Telemetry
}

//NewDeviceDashboardTelemetry ..
func NewDeviceDashboardTelemetry() *DeviceDashboardTelemetry {
	m := &DeviceDashboardTelemetry{}
	m.SetDB(db.GetDB())
	return m
}
