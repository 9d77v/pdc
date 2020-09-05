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
	switch deviceMsg.DeviceAction {
	case pb.DeviceAction_GET_DEVICE_INFO:
		replyDeviceMsg := &pb.DeviceMSG{
			DeviceID: deviceMsg.DeviceID,
		}
		device := new(models.Device)
		err := models.Gorm.Select("ip,port").Where("id=?", replyDeviceMsg.DeviceID).First(device).Error
		if err != nil {
			iotsdk.ReplyDeviceInfo(m.Reply, replyDeviceMsg)
			return
		}
		deviceInfo := &pb.DeviceInfo{
			DeviceID: replyDeviceMsg.DeviceID,
			IP:       device.IP,
			Port:     uint32(device.Port),
		}
		replyDeviceMsg.DeviceInfo = deviceInfo
		iotsdk.ReplyDeviceInfo(m.Reply, replyDeviceMsg)
	}
}

//FilterDeviceMSG ...
func FilterDeviceMSG(m *stan.Msg) {
	deviceMsg := new(pb.DeviceMSG)
	err := proto.Unmarshal(m.Data, deviceMsg)
	if err != nil {
		log.Println("unmarshal data error")
		return
	}
	switch deviceMsg.DeviceAction {
	case pb.DeviceAction_SET_ATTRIBUTES:
		saveAttributes(deviceMsg)
	case pb.DeviceAction_UPLOAD_TELEMETRIES:
		sendDeviceData(deviceMsg)
	}
}

//SaveDeviceData ...
func SaveDeviceData(m *stan.Msg) {
	deviceMsg := new(pb.DeviceMSG)
	err := proto.Unmarshal(m.Data, deviceMsg)
	if err != nil {
		log.Println("unmarshal data error")
		return
	}
	tx, _ := clickhouse.Client.Begin()
	stmt, _ := tx.Prepare("INSERT INTO telemetry (device_id,telemetry_id,action_time,action_time_nano_second, value, created_at,created_at_nano_second) VALUES (?,?, ?, ?, ?,?,?)")
	defer stmt.Close()
	now := time.Now()
	for _, v := range deviceMsg.Telemetries {
		actionTime, _ := ptypes.Timestamp(v.ActionTime)
		if _, err := stmt.Exec(
			deviceMsg.DeviceID,
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
	return
}

//PublishDeviceData ...
func PublishDeviceData(m *stan.Msg) {
	deviceMsg := new(pb.DeviceMSG)
	err := proto.Unmarshal(m.Data, deviceMsg)
	if err != nil {
		log.Println("unmarshal data error")
		return
	}
	iotsdk := new(sdk.IotSDK)
	for _, v := range deviceMsg.Telemetries {
		iotsdk.PublishTelemetry(v)
	}
	return
}

func saveAttributes(deviceMsg *pb.DeviceMSG) {
	attributes := make([]*models.Attribute, 0, len(deviceMsg.AttributeMap))
	keys := make([]string, 0, len(deviceMsg.AttributeMap))
	for k := range deviceMsg.AttributeMap {
		keys = append(keys, k)
	}
	if len(keys) == 0 {
		return
	}
	err := models.Gorm.Where("device_id=? and key in (?)", deviceMsg.DeviceID, keys).Find(&attributes).Error
	if err != nil {
		log.Println("get attribute schema faield, error:", err)
		return
	}
	for _, v := range attributes {
		err := models.Gorm.Model(&models.Attribute{}).
			Where("id=?", v.ID).
			Update("value", deviceMsg.AttributeMap[v.Key]).Error
		if err != nil {
			log.Printf("update attribute failed,id:%d,key:%s\n", deviceMsg.DeviceID, v.Key)
		}
	}
	return
}

func sendDeviceData(deviceMsg *pb.DeviceMSG) {
	telemetries := make([]*models.Telemetry, 0, len(deviceMsg.Telemetries))
	keys := make([]string, 0, len(deviceMsg.Telemetries))
	for k := range deviceMsg.TelemetryMap {
		keys = append(keys, k)
	}
	if len(keys) == 0 {
		return
	}
	err := models.Gorm.Where("device_id=? and key in (?)", deviceMsg.DeviceID, keys).Find(&telemetries).Error
	if err != nil {
		log.Println("get telemetry schema faield, error:", err)
		return
	}
	sendDeviceMsg := new(pb.DeviceMSG)
	sendDeviceMsg.DeviceID = deviceMsg.DeviceID
	publishTelemetries := make([]*pb.Telemetry, 0, len(telemetries))
	for _, v := range telemetries {
		publishTelemetries = append(publishTelemetries, &pb.Telemetry{
			ActionTime: deviceMsg.ActionTime,
			ID:         uint32(v.ID),
			Value:      deviceMsg.TelemetryMap[v.Key],
		})
	}
	sendDeviceMsg.Telemetries = publishTelemetries
	iotsdk := sdk.NewIotSDK()
	iotsdk.PublishDeviceData(sendDeviceMsg)
}
