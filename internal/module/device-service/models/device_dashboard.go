package models

import (
	"github.com/9d77v/pdc/internal/module/base"
)

//DeviceDashboard 设备仪表盘
type DeviceDashboard struct {
	base.DefaultModel
	Name        string `gorm:"size:50"`
	IsVisible   bool
	DeviceType  uint8 //设备类型，0:默认设备,1:摄像头
	Telemetries []*DeviceDashboardTelemetry
	Cameras     []*DeviceDashboardCamera
}

//NewDeviceDashboard ..
func NewDeviceDashboard() *DeviceDashboard {
	vs := &DeviceDashboard{}
	vs.DefaultModel = base.NewDefaultModel()
	return vs
}

//DeviceDashboardTelemetry 仪表盘遥测
type DeviceDashboardTelemetry struct {
	*base.DefaultModel
	DeviceDashboardID uint
	TelemetryID       uint
	Telemetry         Telemetry
}

//DeviceDashboardCamera 仪表盘摄像头
type DeviceDashboardCamera struct {
	*base.DefaultModel
	DeviceDashboardID uint
	DeviceID          uint
	Device            Device
}
