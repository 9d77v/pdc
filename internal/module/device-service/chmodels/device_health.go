package chmodels

import (
	"fmt"
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
		Where(fmt.Sprintf("action_time >= toDateTime64('%s', 3, 'UTC')",
			time.Now().Add(-1*time.Minute).UTC().Format("2006-01-02 15:04:05.999"))).First(m)
	if err != nil {
		return true
	}
	return m.Value == 0
}
