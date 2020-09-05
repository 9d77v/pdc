package clickhouse

import "time"

//Telemetry uploaded telemetry data
type Telemetry struct {
	ActionTime  time.Time `db:"action_time"`
	DeviceID    uint      `db:"device_id"`
	TelemetryID uint      `db:"telemetry_id"`
	Value       float64   `db:"value"`
	CreatedAt   time.Time `db:"created_at"`
}
