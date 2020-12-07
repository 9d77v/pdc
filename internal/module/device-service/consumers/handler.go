package consumers

import (
	"errors"
	"log"
	"time"

	"github.com/9d77v/pdc/internal/db/clickhouse"
	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
	"github.com/golang/protobuf/ptypes"
	"github.com/nats-io/stan.go"
	"google.golang.org/protobuf/proto"
)

var (
	batchSize     = 1000
	telemetryChan = make(chan *pb.Telemetry, batchSize)
	healthChan    = make(chan *pb.Health, batchSize)
	duration      = 1 * time.Second
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
	handler, err := createMsgHandler(deviceMsg)
	if err != nil {
		log.Println("handle device msg failed:", err)
		return
	}
	handler.handleMsg(deviceMsg)
}

//PublishDeviceData ..
func PublishDeviceData(m *stan.Msg) {
	deviceMsg := new(pb.DeviceUpMsg)
	err := proto.Unmarshal(m.Data, deviceMsg)
	if err != nil {
		log.Println("unmarshal data error")
		return
	}
	handler, err := createDataHandler(deviceMsg)
	if err != nil {
		log.Println("publish data failed:", err)
		return
	}
	handler.sendData(deviceMsg)
}

//SaveDeviceTelemetry 。。
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

//SaveDeviceHealth 。。
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
