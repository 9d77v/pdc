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

func (m *setHealthMsg) saveDeviceHealth() {
	var healths []*pb.Health
	timer := time.NewTimer(0)
	if !timer.Stop() {
		<-timer.C
	}
	defer timer.Stop()
	for {
		select {
		case health := <-healthChan:
			healths = append(healths, health)
			if len(healths) != batchSize {
				if len(healths) == 1 {
					timer.Reset(duration)
				}
				break
			}
			batchSaveHealth(healths)
			if !timer.Stop() {
				<-timer.C
			}
			healths = healths[0:0]
		case <-timer.C:
			batchSaveHealth(healths)
			healths = healths[0:0]
		}
	}
}

func batchSaveHealth(healths []*pb.Health) {
	tx, err := clickhouse.GetDB().Begin()
	if err != nil {
		log.Println("clickhouse tx begin error:", err)
		return
	}
	stmt, err := tx.Prepare("INSERT INTO device_health (device_id,action_time,action_time_nanos, value, created_at,created_at_nanos) VALUES (?,?, ?, ?, ?,?,?)")
	if err != nil {
		log.Println("clickhouse tx Prepare error:", err)
		return
	}
	defer stmt.Close()
	now := time.Now()
	for _, v := range healths {
		actionTime, err := ptypes.Timestamp(v.ActionTime)
		if err != nil {
			log.Println("conver actiontime to timestamp error:", err)
		}
		if _, err := stmt.Exec(
			v.DeviceID,
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
