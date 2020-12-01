package consumers

import (
	"context"
	"fmt"
	"log"

	"github.com/9d77v/pdc/internal/db/mq"
	"github.com/9d77v/pdc/internal/db/redis"
	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
	"github.com/golang/protobuf/proto"
)

type setHealthMsg struct {
}

func (m *setHealthMsg) handleMsg(deviceMsg *pb.DeviceUpMsg) {
	healthMsg := deviceMsg.GetSetHealthMsg()
	health := &pb.Health{
		DeviceID:   deviceMsg.DeviceId,
		ActionTime: deviceMsg.ActionTime,
		Value:      healthMsg.DeviceHealth,
	}
	healthChan <- health
}

func (m *setHealthMsg) sendData(deviceMsg *pb.DeviceUpMsg) {
	healthMsg := deviceMsg.GetSetHealthMsg()
	health := &pb.Health{
		DeviceID:   deviceMsg.DeviceId,
		ActionTime: deviceMsg.ActionTime,
		Value:      healthMsg.DeviceHealth,
	}
	requestMsg, err := proto.Marshal(health)
	if err != nil {
		log.Println("proto marshal error:", err)
		return
	}
	err = redis.GetClient().Publish(context.Background(),
		fmt.Sprintf("%s%d", mq.SubjectDeviceHealthPrefix, deviceMsg.DeviceId), requestMsg).Err()
	if err != nil {
		log.Printf("publish error,err:%v/n", err)
	}
}
