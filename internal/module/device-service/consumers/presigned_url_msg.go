package consumers

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/9d77v/pdc/internal/db/mq"
	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

type setPresignedURLSignedMsg struct {
}

func (m *setPresignedURLSignedMsg) sendData(deviceMsg *pb.DeviceUpMsg) {
	presignedURLMsg := deviceMsg.GetPresignedUrlMsg()
	requestURL, err := oss.GetPresignedURL(context.Background(), presignedURLMsg.BucketName, presignedURLMsg.ObjectName, "")
	if err != nil {
		return
	}
	subject := mq.SubjectDevicPrefix + strconv.FormatUint(uint64(deviceMsg.DeviceId), 10)
	request := new(pb.DeviceDownMSG)
	presignedURLReplyMsg := &pb.DeviceDownMSG_PresignedUrlReplyMsg{
		PresignedUrlReplyMsg: &pb.PresignedUrlReplyMsg{
			PictureUrl:      requestURL,
			OssPrefix:       oss.OssPrefix,
			SecureOssPrefix: oss.SecureOssPrerix,
		},
	}
	request.DeviceId = deviceMsg.DeviceId
	request.ActionTime = ptypes.TimestampNow()
	request.Payload = presignedURLReplyMsg
	requestMsg, err := proto.Marshal(request)
	if err != nil {
		log.Println("marshal error:", err)
		return
	}
	_, err = mq.GetClient().NatsConn().Request(subject, requestMsg, 5*time.Second)
	if err != nil {
		log.Println("send data error:", err)
		return
	}
}
