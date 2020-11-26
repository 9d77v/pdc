package consumer

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/nats-io/stan.go"

	"github.com/9d77v/pdc/internal/db/clickhouse"
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/db/mq"
	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/internal/db/redis"
	"github.com/9d77v/pdc/internal/module/device-service/models"
	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
)

var (
	batchSize     = 1000
	telemetryChan = make(chan *pb.Telemetry, batchSize)
	healthChan    = make(chan *pb.Health, batchSize)
	duration      = 1 * time.Second
)

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
		msg := deviceMsg.GetCameraCaptureReplyMsg()
		if msg == nil {
			return
		}
		err = mq.Client.NatsConn().Publish(msg.Subject, m.Data)
		if err != nil {
			log.Println("publish DeviceUpMsg_CameraCaptureReplyMsg failed:", err)
		}
	case *pb.DeviceUpMsg_SetAttributesMsg:
		attributeMsg := deviceMsg.GetSetAttributesMsg()
		if attributeMsg == nil || attributeMsg.AttributeMap == nil {
			return
		}
		for k, v := range attributeMsg.AttributeMap {
			err := db.Gorm.Model(&models.Attribute{}).
				Where("id=?", k).
				Update("value", v).Error
			if err != nil {
				log.Printf("update attribute failed,id:%d,value:%s\n", deviceMsg.DeviceId, v)
			}
		}
	case *pb.DeviceUpMsg_SetTelemetriesMsg:
		telemetryMsg := deviceMsg.GetSetTelemetriesMsg()
		if telemetryMsg == nil || telemetryMsg.TelemetryMap == nil {
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
	case *pb.DeviceUpMsg_SetHealthMsg:
		healthMsg := deviceMsg.GetSetHealthMsg()
		if healthMsg == nil {
			return
		}
		health := &pb.Health{
			DeviceID:   deviceMsg.DeviceId,
			ActionTime: deviceMsg.ActionTime,
			Value:      healthMsg.DeviceHealth,
		}
		healthChan <- health
	}
}

const subjectDeviceTelemetryPrefix = "device.telemetry"
const subjectDeviceHealthPrefix = "device.health"

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
			err = redis.Client.Publish(context.Background(),
				fmt.Sprintf("%s.%d.%d", subjectDeviceTelemetryPrefix, deviceMsg.DeviceId, k), requestMsg).Err()
			if err != nil {
				log.Printf("publish error,err:%v/n", err)
			}
		}
	case *pb.DeviceUpMsg_SetHealthMsg:
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
		err = redis.Client.Publish(context.Background(),
			fmt.Sprintf("%s.%d", subjectDeviceHealthPrefix, deviceMsg.DeviceId), requestMsg).Err()
		if err != nil {
			log.Printf("publish error,err:%v/n", err)
		}
	case *pb.DeviceUpMsg_PresignedUrlMsg:
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
		_, err = mq.Client.NatsConn().Request(subject, requestMsg, 5*time.Second)
		if err != nil {
			log.Println("send data error:", err)
			return
		}
	}
}

//SaveDeviceTelemetry ...
func SaveDeviceTelemetry() {
	var telemetries []*pb.Telemetry
	timer := time.NewTimer(0)
	if !timer.Stop() {
		<-timer.C
	}
	defer timer.Stop()
	for {
		select {
		case telemetry := <-telemetryChan:
			telemetries = append(telemetries, telemetry)
			if len(telemetries) != batchSize {
				if len(telemetries) == 1 {
					timer.Reset(duration)
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

//batchSaveTelemetry ..
func batchSaveTelemetry(telemetries []*pb.Telemetry) {
	tx, err := clickhouse.Client.Begin()
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

//SaveDeviceHealth ...
func SaveDeviceHealth() {
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

//batchSaveHealth ..
func batchSaveHealth(healths []*pb.Health) {
	tx, err := clickhouse.Client.Begin()
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
