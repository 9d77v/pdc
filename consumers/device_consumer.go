package consumers

import (
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"

	"github.com/9d77v/pdc/iot/sdk"
	"github.com/9d77v/pdc/iot/sdk/pb"
	"github.com/9d77v/pdc/models"
	"github.com/9d77v/pdc/models/clickhouse"
)

//ReplyDeviceMSG ...
func ReplyDeviceMSG(m *nats.Msg) {
	deviceMsg := new(pb.DeviceMSG)
	err := proto.Unmarshal(m.Data, deviceMsg)
	if err != nil {
		log.Println("unmarshal data error")
		return
	}
	iotsdk := sdk.NewIotSDK()
	replyDeviceMsg := &pb.DeviceMSG{
		DeviceID: deviceMsg.DeviceID,
	}
	device := new(models.Device)
	err = models.Gorm.Preload("Attributes").Preload("Telemetries").Where("id=?", replyDeviceMsg.DeviceID).First(device).Error
	if err != nil {
		iotsdk.ReplyDeviceInfo(m.Reply, replyDeviceMsg)
		return
	}
	attributeConfig := make(map[string]uint32, 0)
	for _, v := range device.Attributes {
		attributeConfig[v.Key] = uint32(v.ID)
	}
	telemetryConfig := make(map[string]*pb.Telemetry)
	for _, v := range device.Telemetries {
		telemetryConfig[v.Key] = &pb.Telemetry{
			ID:     uint32(v.ID),
			Factor: v.TelemetryModel.Factor,
			Scale:  uint32(v.TelemetryModel.Scale),
		}
	}
	deviceInfo := &pb.DeviceInfo{
		IP:              device.IP,
		Port:            uint32(device.Port),
		AttributeConfig: attributeConfig,
		TelemetryConfig: telemetryConfig,
	}
	replyDeviceMsg.DeviceInfo = deviceInfo
	iotsdk.ReplyDeviceInfo(m.Reply, replyDeviceMsg)
}

var (
	batchSize     = 1000
	telemetryChan = make(chan *pb.Telemetry, batchSize)
	duration      = 1 * time.Second
)

//HandleDeviceMSG ..
func HandleDeviceMSG(m *stan.Msg) {
	deviceMsg := new(pb.DeviceMSG)
	err := proto.Unmarshal(m.Data, deviceMsg)
	if err != nil {
		log.Println("unmarshal data error")
		return
	}
	switch deviceMsg.Action {
	case pb.DeviceAction_SetAttributes:
		setAttributes(deviceMsg)
	case pb.DeviceAction_SetTelemetries:
		for k, v := range deviceMsg.TelemetryMap {
			telemetry := &pb.Telemetry{
				DeviceID:   deviceMsg.DeviceID,
				ActionTime: deviceMsg.ActionTime,
				ID:         k,
				Value:      v,
			}
			telemetryChan <- telemetry
		}
	}
}

func setAttributes(deviceMsg *pb.DeviceMSG) {
	for k, v := range deviceMsg.AttributeMap {
		err := models.Gorm.Model(&models.Attribute{}).
			Where("id=?", k).
			Update("value", v).Error
		if err != nil {
			log.Printf("update attribute failed,id:%d,value:%s\n", deviceMsg.DeviceID, v)
		}
	}
}

//SaveTelemetry ...
func SaveTelemetry() {
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

//batchSaveTelemetry ..
func batchSaveTelemetry(telemetries []*pb.Telemetry) {
	tx, _ := clickhouse.Client.Begin()
	stmt, _ := tx.Prepare("INSERT INTO telemetry (device_id,telemetry_id,action_time,action_time_nanos, value, created_at,created_at_nanos) VALUES (?,?, ?, ?, ?,?,?)")
	defer stmt.Close()
	now := time.Now()
	for _, v := range telemetries {
		actionTime, _ := ptypes.Timestamp(v.ActionTime)
		if _, err := stmt.Exec(
			v.DeviceID,
			v.ID,
			actionTime,
			actionTime.Nanosecond(),
			v.Value,
			now,
			now.Nanosecond(),
		); err != nil {
			log.Println(err)
			tx.Rollback()
			return
		}
	}
	if err := tx.Commit(); err != nil {
		log.Println("commit failed:", err)
	}
}
