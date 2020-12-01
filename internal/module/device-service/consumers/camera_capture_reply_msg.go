package consumers

import (
	"log"

	"github.com/9d77v/pdc/internal/db/mq"
	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
)

type cameraCaptureReplyMsg struct {
	*deviceUpMsg
}

func getCameraCaptureReplyMsg() *cameraCaptureReplyMsg {
	return &cameraCaptureReplyMsg{
		getDeviceMsg(),
	}
}

func (m *cameraCaptureReplyMsg) handleMsg(deviceMsg *pb.DeviceUpMsg, data []byte) {
	msg := deviceMsg.GetCameraCaptureReplyMsg()
	err := mq.GetClient().NatsConn().Publish(msg.Subject, data)
	if err != nil {
		log.Println("publish DeviceUpMsg_CameraCaptureReplyMsg failed:", err)
	}
}
