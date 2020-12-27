package consumers

import (
	"log"
	"time"

	"github.com/9d77v/pdc/internal/db/clickhouse"
	"github.com/9d77v/pdc/internal/module/device-service/chmodels"
	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
	"github.com/golang/protobuf/ptypes"
	"github.com/nats-io/stan.go"
	"google.golang.org/protobuf/proto"
)

var (
	batchSize     = 1000
	telemetryChan = make(chan *pb.Telemetry, batchSize)
	healthChan    = make(chan *pb.Health, batchSize)
	duration      = 1 * time.Second
)

//HandleDeviceMsg ..
func HandleDeviceMsg(m *stan.Msg) {
	deviceMsg := new(pb.DeviceUpMsg)
	err := proto.Unmarshal(m.Data, deviceMsg)
	if err != nil {
		log.Println("unmarshal data error")
		return
	}
	if deviceMsg.Payload == nil {
		return
	}
	handler, err := createMsgHandler(deviceMsg)
	if err != nil {
		log.Println("handle device msg failed:", err)
		return
	}
	handler.handleMsg(deviceMsg)
}

//PublishDeviceData ..
func PublishDeviceData(m *stan.Msg) {
	deviceMsg := new(pb.DeviceUpMsg)
	err := proto.Unmarshal(m.Data, deviceMsg)
	if err != nil {
		log.Println("unmarshal data error")
		return
	}
	handler, err := createDataHandler(deviceMsg)
	if err != nil {
		log.Println("publish data failed:", err)
		return
	}
	handler.sendData(deviceMsg)
}

//SaveDeviceTelemetry 。。
func SaveDeviceTelemetry() {
	var telemetries []*pb.Telemetry
	timer := time.NewTimer(0)
	if !timer.Stop() {
		<-timer.C
	}
	defer timer.Stop()
	for {
		select {
		case telemetry := <-telemetryChan:
			telemetries = append(telemetries, telemetry)
			if len(telemetries) != batchSize {
				if len(telemetries) == 1 {
					timer.Reset(duration)
				}
				break
			}
			batchSaveTelemetry(telemetries)
			if !timer.Stop() {
				<-timer.C
			}
			telemetries = telemetries[0:0]
		case <-timer.C:
			batchSaveTelemetry(telemetries)
			telemetries = telemetries[0:0]
		}
	}
}

func batchSaveTelemetry(telemetries []*pb.Telemetry) {
	if telemetries == nil || len(telemetries) == 0 {
		return
	}
	deviceTelemetries := make([]chmodels.DeviceTelemetry, 0, len(telemetries))
	now := time.Now()
	for _, v := range telemetries {
		actionTime, err := ptypes.Timestamp(v.ActionTime)
		if err != nil {
			log.Println("conver actiontime to timestamp error:", err)
		}
		deviceTelemetries = append(deviceTelemetries,
			chmodels.DeviceTelemetry{
				ActionTime:  actionTime,
				DeviceID:    v.DeviceID,
				TelemetryID: v.ID,
				Value:       v.Value,
				CreatedAt:   now,
			})
	}
	err := clickhouse.GetDB().Create(&deviceTelemetries).Error
	if err != nil {
		log.Println("batchSaveTelemetry failed:", err)
	}
}

//SaveDeviceHealth 。。
func SaveDeviceHealth() {
	var healths []*pb.Health
	timer := time.NewTimer(0)
	if !timer.Stop() {
		<-timer.C
	}
	defer timer.Stop()
	for {
		select {
		case health := <-healthChan:
			healths = append(healths, health)
			if len(healths) != batchSize {
				if len(healths) == 1 {
					timer.Reset(duration)
				}
				break
			}
			batchSaveHealth(healths)
			if !timer.Stop() {
				<-timer.C
			}
			healths = healths[0:0]
		case <-timer.C:
			batchSaveHealth(healths)
			healths = healths[0:0]
		}
	}
}

func batchSaveHealth(healths []*pb.Health) {
	if healths == nil || len(healths) == 0 {
		return
	}
	deviceHealths := make([]chmodels.DeviceHealth, 0, len(healths))
	now := time.Now()
	for _, v := range healths {
		actionTime, err := ptypes.Timestamp(v.ActionTime)
		if err != nil {
			log.Println("conver actiontime to timestamp error:", err)
		}
		deviceHealths = append(deviceHealths,
			chmodels.DeviceHealth{
				ActionTime: actionTime,
				DeviceID:   v.DeviceID,
				Value:      float64(v.Value),
				CreatedAt:  now,
			})
	}
	err := clickhouse.GetDB().Create(&deviceHealths).Error
	if err != nil {
		log.Println("batchSaveHealth failed:", err)
	}
}
