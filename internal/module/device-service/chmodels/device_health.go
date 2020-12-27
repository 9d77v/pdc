package chmodels

import (
	"time"

	"github.com/9d77v/pdc/internal/db/clickhouse"
	"github.com/9d77v/pdc/internal/module/base"
)

//DeviceHealth ..
type DeviceHealth struct {
	base.Model
	ActionTime time.Time
	DeviceID   uint32
	Value      float64
	CreatedAt  time.Time
}

//NewDeviceHealth ..
func NewDeviceHealth() *DeviceHealth {
	m := &DeviceHealth{}
	m.SetDB(clickhouse.GetDB())
	return m
}

//IsUnhealthy check if device is unhealthy
func (m *DeviceHealth) IsUnhealthy() bool {
	err := m.IDQuery(uint(m.DeviceID), "device_id").
		Where("action_time >=?", time.Now().Add(-1*time.Minute)).First(m)
	if err != nil {
		return true
	}
	return m.Value == 0
}
