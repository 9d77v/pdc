package consumers

import (
	"sync"
	"time"

	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
)

//DeviceUpMsg 设备上传消息
type deviceUpMsg struct {
	batchSize     int
	telemetryChan chan *pb.Telemetry
	healthChan    chan *pb.Health
	duration      time.Duration
}

var globalDeviceUpMsg *deviceUpMsg
var once sync.Once

func getDeviceMsg() *deviceUpMsg {
	once.Do(func() {
		globalDeviceUpMsg = &deviceUpMsg{
			batchSize:     batchSize,
			telemetryChan: make(chan *pb.Telemetry, batchSize),
			healthChan:    make(chan *pb.Health, batchSize),
			duration:      1 * time.Second,
		}
	})
	return globalDeviceUpMsg
}
