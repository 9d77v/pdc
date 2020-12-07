package consumers

import (
	"errors"

	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
)

type msgHandler interface {
	handleMsg(deviceMsg *pb.DeviceUpMsg)
}

func createMsgHandler(deviceMsg *pb.DeviceUpMsg) (msgHandler, error) {
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

func createDataHandler(deviceMsg *pb.DeviceUpMsg) (dataHandler, error) {
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
