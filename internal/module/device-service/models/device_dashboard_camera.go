package models

import (
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
)

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
