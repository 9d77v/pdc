package consumers

import (
	"time"

	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
)

var (
	batchSize     = 1000
	telemetryChan = make(chan *pb.Telemetry, batchSize)
	healthChan    = make(chan *pb.Health, batchSize)
	duration      = 1 * time.Second
)

func init() {
	go new(setTelemetriesMsg).saveDeviceTelemetry()
	go new(setHealthMsg).saveDeviceHealth()
}
