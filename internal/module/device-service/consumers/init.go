package consumers

import (
	"log"

	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
	"github.com/nats-io/stan.go"
	"google.golang.org/protobuf/proto"
)

var batchSize = 1000

func init() {
	go getSetTelemetriesMsg().saveDeviceTelemetry()
	go getSetHealthMsg().saveDeviceHealth()
}

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
	switch deviceMsg.Payload.(type) {
	case *pb.DeviceUpMsg_CameraCaptureReplyMsg:
		getCameraCaptureReplyMsg().handleMsg(deviceMsg, m.Data)
	case *pb.DeviceUpMsg_SetAttributesMsg:
		getSetAttributesMsg().handleMsg(deviceMsg)
	case *pb.DeviceUpMsg_SetTelemetriesMsg:
		getSetTelemetriesMsg().handleMsg(deviceMsg)
	case *pb.DeviceUpMsg_SetHealthMsg:
		getSetHealthMsg().handleMsg(deviceMsg)
	}
}

//PublishDeviceData push device data to redis
func PublishDeviceData(m *stan.Msg) {
	deviceMsg := new(pb.DeviceUpMsg)
	err := proto.Unmarshal(m.Data, deviceMsg)
	if err != nil {
		log.Println("unmarshal data error")
		return
	}
	switch deviceMsg.Payload.(type) {
	case *pb.DeviceUpMsg_SetTelemetriesMsg:
		getSetTelemetriesMsg().publishData(deviceMsg)
	case *pb.DeviceUpMsg_SetHealthMsg:
		getSetHealthMsg().publishData(deviceMsg)
	case *pb.DeviceUpMsg_PresignedUrlMsg:
		getSetPresignedURLSignedMsg().sendData(deviceMsg)
	}
}
