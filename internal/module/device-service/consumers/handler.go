package consumers

import (
	"errors"
	"log"

	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
	"github.com/nats-io/stan.go"
	"google.golang.org/protobuf/proto"
)

type msgHandler interface {
	handleMsg(deviceMsg *pb.DeviceUpMsg)
}

func newMsgHandler(deviceMsg *pb.DeviceUpMsg) (msgHandler, error) {
	switch deviceMsg.Payload.(type) {
	case *pb.DeviceUpMsg_CameraCaptureReplyMsg:
		return &cameraCaptureReplyMsg{}, nil
	case *pb.DeviceUpMsg_SetAttributesMsg:
		return &setAttributesMsg{}, nil
	case *pb.DeviceUpMsg_SetTelemetriesMsg:
		return &setTelemetriesMsg{}, nil
	case *pb.DeviceUpMsg_SetHealthMsg:
		return &setHealthMsg{}, nil
	default:
		return nil, errors.New("invalid payload type")
	}
}

type dataHandler interface {
	sendData(deviceMsg *pb.DeviceUpMsg)
}

func newDataHandler(deviceMsg *pb.DeviceUpMsg) (dataHandler, error) {
	switch deviceMsg.Payload.(type) {
	case *pb.DeviceUpMsg_SetTelemetriesMsg:
		return &setTelemetriesMsg{}, nil
	case *pb.DeviceUpMsg_SetHealthMsg:
		return &setHealthMsg{}, nil
	case *pb.DeviceUpMsg_PresignedUrlMsg:
		return &setPresignedURLSignedMsg{}, nil
	default:
		return nil, errors.New("invalid payload type")
	}
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
	handler, err := newMsgHandler(deviceMsg)
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
	handler, err := newDataHandler(deviceMsg)
	if err != nil {
		log.Println("publish data failed:", err)
		return
	}
	handler.sendData(deviceMsg)
}
