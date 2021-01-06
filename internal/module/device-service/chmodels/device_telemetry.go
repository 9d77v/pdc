package chmodels

import "time"

//DeviceTelemetry uploaded telemetry data
type DeviceTelemetry struct {
	ActionTime  time.Time `gorm:"type:DateTime64(3)"`
	DeviceID    uint32
	TelemetryID uint32
	Value       float64
	CreatedAt   time.Time `gorm:"type:DateTime64(3)"`
}
