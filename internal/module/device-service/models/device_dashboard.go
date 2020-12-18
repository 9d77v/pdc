package models

import (
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/device-service/pb"
	"gorm.io/gorm"
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
	m := &DeviceDashboard{}
	m.SetDB(db.GetDB())
	return m
}

//NewDeviceDashboardFromPB ..
func NewDeviceDashboardFromPB(in *pb.CreateDeviceDashboardRequest) *DeviceDashboard {
	m := &DeviceDashboard{
		Name:       in.Name,
		IsVisible:  in.IsVisible,
		DeviceType: uint8(in.DeviceType),
	}
	m.SetDB(db.GetDB())
	return m
}

//GetByID ..
func (m *DeviceDashboard) GetByID(id uint) error {
	return m.IDQuery(id).Preload("Cameras", func(db *gorm.DB) *gorm.DB {
		return db.Model(&DeviceDashboardCamera{})
	}).Preload("Telemetries", func(db *gorm.DB) *gorm.DB {
		return db.Model(&DeviceDashboardTelemetry{})
	}).First(m)
}

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

//DeviceDashboardCamera 仪表盘摄像头
type DeviceDashboardCamera struct {
	base.DefaultModel
	DeviceDashboardID uint
	DeviceID          uint
	Device            Device
}

//NewDeviceDashboardCamera ..
func NewDeviceDashboardCamera() *DeviceDashboardCamera {
	m := &DeviceDashboardCamera{}
	m.SetDB(db.GetDB())
	return m
}
