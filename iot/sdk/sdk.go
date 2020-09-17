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
	conn  stan.Conn
	mutex *sync.Mutex
}

//NewIotSDK init iot sdk
func NewIotSDK() *IotSDK {
	return &IotSDK{
		conn:  natsConn,
		mutex: new(sync.Mutex),
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
	msg, err := sdk.conn.NatsConn().Request(subjectDevice, requestMsg, 5*time.Second)
	if err != nil {
		return nil, err
	}
	reply := new(pb.DeviceMSG)
	err = proto.Unmarshal(msg.Data, reply)
	if err != nil {
		return nil, err
	}
	fmt.Println(reply)
	return reply.DeviceInfo, nil
}

//ReplyDeviceInfo reply device msg
func (sdk *IotSDK) ReplyDeviceInfo(replySubject string, deviceMsg *pb.DeviceMSG) {
	reply, marshalError := proto.Marshal(deviceMsg)
	if marshalError != nil {
		log.Println("unmarshal deviceMSG error:", marshalError)
	}
	err := sdk.conn.NatsConn().Publish(replySubject, reply)
	if err != nil {
		log.Println("publish error:", err)
	}
}

//SetDeviceAttributes set device attributes
func (sdk *IotSDK) SetDeviceAttributes(deviceID uint32, attributeMap map[uint32]string) {
	request := new(pb.DeviceMSG)
	request.Action = pb.DeviceAction_SetAttributes
	request.DeviceID = deviceID
	request.AttributeMap = attributeMap
	request.ActionTime = ptypes.TimestampNow()
	requestMsg, err := proto.Marshal(request)
	if err != nil {
		log.Println("proto marshal error:", err)
		return
	}
	msg, err := sdk.conn.PublishAsync(subjectDeviceData, requestMsg, AckHandler)
	if err != nil {
		log.Printf("publish error,id:%s,err:%v/n", msg, err)
	}
}

//SetDeviceTelemetries upload device telemetries
func (sdk *IotSDK) SetDeviceTelemetries(deviceID uint32, telemetryMap map[uint32]float64, now *timestamppb.Timestamp) {
	request := new(pb.DeviceMSG)
	request.Action = pb.DeviceAction_SetTelemetries
	request.DeviceID = deviceID
	request.TelemetryMap = telemetryMap
	request.ActionTime = now
	requestMsg, err := proto.Marshal(request)
	if err != nil {
		log.Println("proto marshal error:", err)
		return
	}
	msg, err := sdk.conn.PublishAsync(subjectDeviceData, requestMsg, AckHandler)
	if err != nil {
		log.Printf("publish error,id:%s,err:%v/n", msg, err)
	}
}

//SetDeviceHealth upload device health
func (sdk *IotSDK) SetDeviceHealth(deviceID uint32, health uint32, now *timestamppb.Timestamp) {
	request := new(pb.DeviceMSG)
	request.Action = pb.DeviceAction_SetHealth
	request.DeviceID = deviceID
	request.DeviceHealth = health
	request.ActionTime = now
	requestMsg, err := proto.Marshal(request)
	if err != nil {
		log.Println("proto marshal error:", err)
		return
	}
	msg, err := sdk.conn.PublishAsync(subjectDeviceData, requestMsg, AckHandler)
	if err != nil {
		log.Printf("publish error,id:%s,err:%v/n", msg, err)
	}
}

//NatsSubscribe ..
func (sdk *IotSDK) NatsSubscribe(handler func(m *nats.Msg)) (*nats.Subscription, error) {
	return sdk.conn.NatsConn().QueueSubscribe(subjectDevice, groupDevice, handler)
}

//SubscribeSaveDeviceData ..
func (sdk *IotSDK) SubscribeSaveDeviceData(handler func(m *stan.Msg)) (stan.Subscription, error) {
	return sdk.conn.QueueSubscribe(subjectDeviceData, groupSaveDeviceData, handler, stan.DurableName("dur"))
}

//SubscribePublishDeviceData ..
func (sdk *IotSDK) SubscribePublishDeviceData(handler func(m *stan.Msg)) (stan.Subscription, error) {
	return sdk.conn.QueueSubscribe(subjectDeviceData, groupPublishDeviceData, handler, stan.DurableName("dur"))
}
