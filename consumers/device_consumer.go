package consumers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/nats-io/stan.go"

	"github.com/9d77v/pdc/iot/sdk/pb"
	"github.com/9d77v/pdc/models"
	"github.com/9d77v/pdc/models/clickhouse"
)

var (
	batchSize     = 1000
	telemetryChan = make(chan *pb.Telemetry, batchSize)
	healthChan    = make(chan *pb.Health, batchSize)
	duration      = 1 * time.Second
)

//HandleDeviceMSG ..
func HandleDeviceMSG(m *stan.Msg) {
	deviceMsg := new(pb.DeviceMSG)
	err := proto.Unmarshal(m.Data, deviceMsg)
	if err != nil {
		log.Println("unmarshal data error")
		return
	}
	switch deviceMsg.Action {
	case pb.DeviceAction_SetAttributes:
		setAttributes(deviceMsg)
	case pb.DeviceAction_SetTelemetries:
		for k, v := range deviceMsg.TelemetryMap {
			telemetry := &pb.Telemetry{
				DeviceID:   deviceMsg.DeviceID,
				ActionTime: deviceMsg.ActionTime,
				ID:         k,
				Value:      v,
			}
			telemetryChan <- telemetry
		}
	case pb.DeviceAction_SetHealth:
		health := &pb.Health{
			DeviceID:   deviceMsg.DeviceID,
			ActionTime: deviceMsg.ActionTime,
			Value:      deviceMsg.DeviceHealth,
		}
		healthChan <- health
	}
}

const subjectDeviceTelemetryPrefix = "device.telemetry"
const subjectDeviceHealthPrefix = "device.health"

//PublishDeviceData push device data to redis
func PublishDeviceData(m *stan.Msg) {
	deviceMsg := new(pb.DeviceMSG)
	err := proto.Unmarshal(m.Data, deviceMsg)
	if err != nil {
		log.Println("unmarshal data error")
		return
	}
	switch deviceMsg.Action {
	case pb.DeviceAction_SetTelemetries:
		for k, v := range deviceMsg.TelemetryMap {
			telemetry := &pb.Telemetry{
				DeviceID:   deviceMsg.DeviceID,
				ActionTime: deviceMsg.ActionTime,
				ID:         k,
				Value:      v,
			}
			requestMsg, err := proto.Marshal(telemetry)
			if err != nil {
				log.Println("proto marshal error:", err)
				return
			}
			err = models.RedisClient.Publish(context.Background(),
				fmt.Sprintf("%s.%d.%d", subjectDeviceTelemetryPrefix, deviceMsg.DeviceID, k), requestMsg).Err()
			if err != nil {
				log.Printf("publish error,err:%v/n", err)
			}
		}
	case pb.DeviceAction_SetHealth:
		health := &pb.Health{
			DeviceID:   deviceMsg.DeviceID,
			ActionTime: deviceMsg.ActionTime,
			Value:      deviceMsg.DeviceHealth,
		}
		requestMsg, err := proto.Marshal(health)
		if err != nil {
			log.Println("proto marshal error:", err)
			return
		}
		err = models.RedisClient.Publish(context.Background(),
			fmt.Sprintf("%s.%d", subjectDeviceHealthPrefix, deviceMsg.DeviceID), requestMsg).Err()
		if err != nil {
			log.Printf("publish error,err:%v/n", err)
		}
	}
}

func setAttributes(deviceMsg *pb.DeviceMSG) {
	for k, v := range deviceMsg.AttributeMap {
		err := models.Gorm.Model(&models.Attribute{}).
			Where("id=?", k).
			Update("value", v).Error
		if err != nil {
			log.Printf("update attribute failed,id:%d,value:%s\n", deviceMsg.DeviceID, v)
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
	tx, _ := clickhouse.Client.Begin()
	stmt, _ := tx.Prepare("INSERT INTO device_telemetry (device_id,telemetry_id,action_time,action_time_nanos, value, created_at,created_at_nanos) VALUES (?,?, ?, ?, ?,?,?)")
	defer stmt.Close()
	now := time.Now()
	for _, v := range telemetries {
		actionTime, _ := ptypes.Timestamp(v.ActionTime)
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
			tx.Rollback()
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
	tx, _ := clickhouse.Client.Begin()
	stmt, _ := tx.Prepare("INSERT INTO device_health (device_id,action_time,action_time_nanos, value, created_at,created_at_nanos) VALUES (?,?, ?, ?, ?,?,?)")
	defer stmt.Close()
	now := time.Now()
	for _, v := range healths {
		actionTime, _ := ptypes.Timestamp(v.ActionTime)
		if _, err := stmt.Exec(
			v.DeviceID,
			actionTime,
			actionTime.Nanosecond(),
			v.Value,
			now,
			now.Nanosecond(),
		); err != nil {
			log.Println(err)
			tx.Rollback()
			return
		}
	}
	if err := tx.Commit(); err != nil {
		log.Println("commit failed:", err)
	}
}
