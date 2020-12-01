package consumers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/9d77v/pdc/internal/db/clickhouse"
	"github.com/9d77v/pdc/internal/db/mq"
	"github.com/9d77v/pdc/internal/db/redis"
	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
)

type setTelemetriesMsg struct {
	*deviceUpMsg
}

func getSetTelemetriesMsg() *setTelemetriesMsg {
	return &setTelemetriesMsg{
		getDeviceMsg(),
	}
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
		m.telemetryChan <- telemetry
	}
}

func (m *setTelemetriesMsg) publishData(deviceMsg *pb.DeviceUpMsg) {
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

func (m *setTelemetriesMsg) saveDeviceTelemetry() {
	var telemetries []*pb.Telemetry
	timer := time.NewTimer(0)
	if !timer.Stop() {
		<-timer.C
	}
	defer timer.Stop()
	for {
		select {
		case telemetry := <-m.telemetryChan:
			telemetries = append(telemetries, telemetry)
			if len(telemetries) != m.batchSize {
				if len(telemetries) == 1 {
					timer.Reset(m.duration)
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

func batchSaveTelemetry(telemetries []*pb.Telemetry) {
	tx, err := clickhouse.GetDB().Begin()
	if err != nil {
		log.Println("clickhouse tx begin error:", err)
		return
	}
	stmt, err := tx.Prepare("INSERT INTO device_telemetry (device_id,telemetry_id,action_time,action_time_nanos, value, created_at,created_at_nanos) VALUES (?,?, ?, ?, ?,?,?)")
	if err != nil {
		log.Println("clickhouse tx Prepare error:", err)
		return
	}
	defer stmt.Close()
	now := time.Now()
	for _, v := range telemetries {
		actionTime, err := ptypes.Timestamp(v.ActionTime)
		if err != nil {
			log.Println("conver actiontime to timestamp error:", err)
		}
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
			err = tx.Rollback()
			if err != nil {
				log.Println("clickhouse tx rollback error:", err)
			}
			return
		}
	}
	if err := tx.Commit(); err != nil {
		log.Println("commit failed:", err)
	}
}
