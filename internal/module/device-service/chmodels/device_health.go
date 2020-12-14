package chmodels

import "time"

//DeviceHealth ..
type DeviceHealth struct {
	ActionTime      time.Time
	ActionTimeNanos uint32
	DeviceID        uint32
	Value           float64
	CreatedAt       time.Time
	CreatedAtNanos  uint32
}
