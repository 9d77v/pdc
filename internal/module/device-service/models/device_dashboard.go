package models

import "gorm.io/gorm"

//DeviceDashboard 设备仪表盘
type DeviceDashboard struct {
	gorm.Model
	Name        string `gorm:"size:50"`
	IsVisible   bool
	DeviceType  uint8 //设备类型，0:默认设备,1:摄像头
	Telemetries []*DeviceDashboardTelemetry
	Cameras     []*DeviceDashboardCamera
}

//DeviceDashboardTelemetry 仪表盘遥测
type DeviceDashboardTelemetry struct {
	gorm.Model
	DeviceDashboardID uint
	TelemetryID       uint
	Telemetry         Telemetry
}

//DeviceDashboardCamera 仪表盘摄像头
type DeviceDashboardCamera struct {
	gorm.Model
	DeviceDashboardID uint
	DeviceID          uint
	Device            Device
}
