package models

import (
	"github.com/9d77v/pdc/internal/module/base"
	"gorm.io/gorm"
)

//DeviceDashboard 设备仪表盘
type DeviceDashboard struct {
	*base.Model
	Name        string `gorm:"size:50"`
	IsVisible   bool
	DeviceType  uint8 //设备类型，0:默认设备,1:摄像头
	Telemetries []*DeviceDashboardTelemetry
	Cameras     []*DeviceDashboardCamera
}

//NewDeviceDashboard ..
func NewDeviceDashboard() *DeviceDashboard {
	vs := &DeviceDashboard{}
	vs.Model = base.NewModel()
	return vs
}

//GetByID ..
func (m *DeviceDashboard) GetByID(id uint, columns []string) error {
	return m.Select(columns).IDQuery(id).First(m)
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
