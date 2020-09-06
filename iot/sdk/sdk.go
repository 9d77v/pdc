package sdk

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/9d77v/pdc/iot/sdk/pb"
	"github.com/golang/protobuf/ptypes"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//IotSDK IotSDK
type IotSDK struct {
	client          stan.Conn
	telemetryConfig map[string]*pb.Telemetry
	mutex           *sync.Mutex
}

//NewIotSDK init iot sdk
func NewIotSDK() *IotSDK {
	return &IotSDK{
		client:          client,
		telemetryConfig: make(map[string]*pb.Telemetry),
		mutex:           new(sync.Mutex),
	}
}

//GetTelemetryConfig ..
func (sdk *IotSDK) GetTelemetryConfig() map[string]*pb.Telemetry {
	sdk.mutex.Lock()
	defer sdk.mutex.Unlock()
	return sdk.telemetryConfig
}

//SetTelemetryConfig ..
func (sdk *IotSDK) SetTelemetryConfig(telemetryMap map[string]*pb.Telemetry) {
	sdk.mutex.Lock()
	defer sdk.mutex.Unlock()
	if telemetryMap != nil {
		sdk.telemetryConfig = telemetryMap
	}
}

//GetDeviceInfo get device info by device id
func (sdk *IotSDK) GetDeviceInfo(deviceID uint32) (*pb.DeviceInfo, error) {
	request := new(pb.DeviceMSG)
	request.DeviceID = uint32(deviceID)
	requestMsg, err := proto.Marshal(request)
	if err != nil {
		return nil, err
	}
	msg, err := sdk.client.NatsConn().Request(subjectDevice, requestMsg, 5*time.Second)
	if err != nil {
		return nil, err
	}
	reply := new(pb.DeviceMSG)
	err = proto.Unmarshal(msg.Data, reply)
	if err != nil {
		return nil, err
	}
	return reply.DeviceInfo, nil
}

//ReplyDeviceInfo reply device msg
func (sdk *IotSDK) ReplyDeviceInfo(replySubject string, deviceMsg *pb.DeviceMSG) {
	reply, marshalError := proto.Marshal(deviceMsg)
	if marshalError != nil {
		log.Println("unmarshal deviceMSG error:", marshalError)
	}
	err := sdk.client.NatsConn().Publish(replySubject, reply)
	if err != nil {
		log.Println("publish error:", err)
	}
}

//SetDeviceAttributes set device attributes
func (sdk *IotSDK) SetDeviceAttributes(deviceID uint32, attributes map[uint32]string) {
	request := new(pb.DeviceMSG)
	request.Action = pb.DeviceAction_SetAttributes
	request.DeviceID = deviceID
	request.AttributeMap = attributes
	request.ActionTime = ptypes.TimestampNow()
	requestMsg, err := proto.Marshal(request)
	if err != nil {
		log.Println("proto marshal error:", err)
		return
	}
	msg, err := sdk.client.PublishAsync(subjectDeviceData, requestMsg, ackHandler)
	if err != nil {
		log.Printf("publish error,id:%s,err:%v/n", msg, err)
	}
}

//UploadDeviceTelemetries upload device telemetries
func (sdk *IotSDK) UploadDeviceTelemetries(deviceID uint32, telemetries map[uint32]float64, now *timestamppb.Timestamp) {
	request := new(pb.DeviceMSG)
	request.Action = pb.DeviceAction_SetTelemetries
	request.DeviceID = deviceID
	request.TelemetryMap = telemetries
	request.ActionTime = now
	requestMsg, err := proto.Marshal(request)
	if err != nil {
		log.Println("proto marshal error:", err)
		return
	}
	msg, err := sdk.client.PublishAsync(subjectDeviceData, requestMsg, ackHandler)
	if err != nil {
		log.Printf("publish error,id:%s,err:%v/n", msg, err)
	}
}

//PublishDeviceTelemetries Publish device telemetries
func (sdk *IotSDK) PublishDeviceTelemetries(deviceID uint32, telemetries map[uint32]float64, now *timestamppb.Timestamp) {
	for k, v := range telemetries {
		request := &pb.Telemetry{
			DeviceID:   deviceID,
			ActionTime: now,
			ID:         k,
			Value:      v,
		}
		requestMsg, err := proto.Marshal(request)
		if err != nil {
			log.Println("proto marshal error:", err)
			return
		}
		err = sdk.client.NatsConn().Publish(fmt.Sprintf("%s.%d.%d", subjectDeviceTelemetryPrefix, deviceID, k), requestMsg)
		if err != nil {
			log.Printf("publish error,err:%v/n", err)
		}
	}
}

//NatsSubscribe ..
func (sdk *IotSDK) NatsSubscribe(handler func(m *nats.Msg)) (*nats.Subscription, error) {
	return sdk.client.NatsConn().QueueSubscribe(subjectDevice, groupDevice, handler)
}

//SubscribeDeviceInfo ..
func (sdk *IotSDK) SubscribeDeviceInfo(deviceID uint32) (stan.Subscription, error) {
	return sdk.client.Subscribe(fmt.Sprintf("%s.%d", subjectDevice, deviceID), func(m *stan.Msg) {
		deviceMsg := new(pb.DeviceMSG)
		err := proto.Unmarshal(m.Data, deviceMsg)
		if err != nil {
			log.Println("unmarshal data error")
			return
		}
		sdk.SetTelemetryConfig(deviceMsg.DeviceInfo.TelemetryConfig)
	})
}

//SubscribeDeviceData ..
func (sdk *IotSDK) SubscribeDeviceData(handler func(m *stan.Msg)) (stan.Subscription, error) {
	return sdk.client.QueueSubscribe(subjectDeviceData, groupDeviceData, handler, stan.DurableName("dur"))
}

//SubscribeDeviceTelemetry subscribe telemetry data
func (sdk *IotSDK) SubscribeDeviceTelemetry(deviceID uint32, telemetryID uint, handler func(m *nats.Msg)) (*nats.Subscription, error) {
	subject := ""
	if deviceID == 0 {
		subject = fmt.Sprintf("%s.*", subjectDeviceTelemetryPrefix)
	} else if telemetryID == 0 {
		subject = fmt.Sprintf("%s.%d.*", subjectDeviceTelemetryPrefix, deviceID)
	} else {
		subject = fmt.Sprintf("%s.%d.%d", subjectDeviceTelemetryPrefix, deviceID, telemetryID)

	}
	return sdk.client.NatsConn().Subscribe(subject, handler)
}
