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

type setTelemetriesMsg struct {
}

func (m *setTelemetriesMsg) handleMsg(deviceMsg *pb.DeviceUpMsg) {
	telemetryMsg := deviceMsg.GetSetTelemetriesMsg()
	if telemetryMsg.TelemetryMap == nil {
		return
	}
	for k, v := range telemetryMsg.TelemetryMap {
		telemetry := &pb.Telemetry{
			DeviceID:   deviceMsg.DeviceId,
			ActionTime: deviceMsg.ActionTime,
			ID:         k,
			Value:      v,
		}
		telemetryChan <- telemetry
	}
}

func (m *setTelemetriesMsg) sendData(deviceMsg *pb.DeviceUpMsg) {
	telemetryMsg := deviceMsg.GetSetTelemetriesMsg()
	if telemetryMsg.TelemetryMap == nil {
		return
	}
	for k, v := range telemetryMsg.TelemetryMap {
		telemetry := &pb.Telemetry{
			DeviceID:   deviceMsg.DeviceId,
			ActionTime: deviceMsg.ActionTime,
			ID:         k,
			Value:      v,
		}
		requestMsg, err := proto.Marshal(telemetry)
		if err != nil {
			log.Println("proto marshal error:", err)
			return
		}
		err = redis.GetClient().Publish(context.Background(),
			fmt.Sprintf("%s%d.%d", mq.SubjectDeviceTelemetryPrefix, deviceMsg.DeviceId, k), requestMsg).Err()
		if err != nil {
			log.Printf("publish error,err:%v/n", err)
		}
	}
}
