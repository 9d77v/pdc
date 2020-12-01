package consumers

import (
	"log"

	"github.com/9d77v/pdc/internal/db/mq"
	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
	"github.com/golang/protobuf/proto"
)

type cameraCaptureReplyMsg struct {
}

func (m *cameraCaptureReplyMsg) handleMsg(deviceMsg *pb.DeviceUpMsg) {
	data, err := proto.Marshal(deviceMsg)
	if err != nil {
		log.Println("proto marsha failed", err)
		return
	}
	msg := deviceMsg.GetCameraCaptureReplyMsg()
	err = mq.GetClient().NatsConn().Publish(msg.Subject, data)
	if err != nil {
		log.Println("publish DeviceUpMsg_CameraCaptureReplyMsg failed:", err)
	}
}
