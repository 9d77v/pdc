package chmodels

import "time"

//DeviceTelemetry uploaded telemetry data
type DeviceTelemetry struct {
	ActionTime      time.Time
	ActionTimeNanos uint32
	DeviceID        uint32
	TelemetryID     uint32
	Value           float64
	CreatedAt       time.Time
	CreatedAtNanos  uint32
}
