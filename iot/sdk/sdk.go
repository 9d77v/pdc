package sdk

import (
	"fmt"
	"log"
	"time"

	"github.com/9d77v/pdc/iot/sdk/pb"
	"github.com/golang/protobuf/ptypes"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"google.golang.org/protobuf/proto"
)

//IotSDK IotSDK
type IotSDK struct {
	client stan.Conn
}

//NewIotSDK init iot sdk
func NewIotSDK() *IotSDK {
	return &IotSDK{
		client: client,
	}
}

//GetDeviceInfo get device info by device id
func (sdk *IotSDK) GetDeviceInfo(deviceID int) (*pb.DeviceInfo, error) {
	request := new(pb.DeviceMSG)
	request.DeviceAction = pb.DeviceAction_GET_DEVICE_INFO
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
func (sdk *IotSDK) SetDeviceAttributes(deviceID int, attributes map[string]string) {
	request := new(pb.DeviceMSG)
	request.DeviceAction = pb.DeviceAction_SET_ATTRIBUTES
	request.DeviceID = uint32(deviceID)
	request.AttributeMap = attributes
	request.ActionTime = ptypes.TimestampNow()
	requestMsg, err := proto.Marshal(request)
	if err != nil {
		log.Println("proto marshal error:", err)
		return
	}
	msg, err := sdk.client.PublishAsync(subjectDeviceFilter, requestMsg, ackHandler)
	if err != nil {
		log.Printf("publish error,id:%s,err:%v/n", msg, err)
	}
}

//UploadDeviceTelemetries upload device telemetries
func (sdk *IotSDK) UploadDeviceTelemetries(deviceID int, telemetries map[string]float64) {
	request := new(pb.DeviceMSG)
	request.DeviceAction = pb.DeviceAction_UPLOAD_TELEMETRIES
	request.DeviceID = uint32(deviceID)
	request.ActionTime = ptypes.TimestampNow()
	request.TelemetryMap = telemetries
	requestMsg, err := proto.Marshal(request)
	if err != nil {
		log.Println("proto marshal error:", err)
		return
	}
	msg, err := sdk.client.PublishAsync(subjectDeviceFilter, requestMsg, ackHandler)
	if err != nil {
		log.Printf("publish error,id:%s,err:%v/n", msg, err)
	}
}

//PublishDeviceData publish telemtry
func (sdk *IotSDK) PublishDeviceData(deviceMsg *pb.DeviceMSG) {
	data, marshalError := proto.Marshal(deviceMsg)
	if marshalError != nil {
		log.Println("unmarshal telemtry error:", marshalError)
	}
	msg, err := sdk.client.PublishAsync(subjectDeviceData, data, ackHandler)
	if err != nil {
		log.Printf("publish error,id:%s,err:%v/n", msg, err)
	}
}

//PublishTelemetry publish telemtry
func (sdk *IotSDK) PublishTelemetry(telemtry *pb.Telemetry) {
	data, marshalError := proto.Marshal(telemtry)
	if marshalError != nil {
		log.Println("unmarshal telemtry error:", marshalError)
	}
	msg, err := sdk.client.PublishAsync(fmt.Sprintf("%s:%d", subjectDeviceData, telemtry.ID), data, ackHandler)
	if err != nil {
		log.Printf("publish error,id:%s,err:%v/n", msg, err)
	}
}

//NatsSubscribe ..
func (sdk *IotSDK) NatsSubscribe(handler func(m *nats.Msg)) (*nats.Subscription, error) {
	return sdk.client.NatsConn().QueueSubscribe(subjectDevice, groupDeviceFilter, handler)
}

//SubscribeDeviceFilter ..
func (sdk *IotSDK) SubscribeDeviceFilter(handler func(m *stan.Msg)) (stan.Subscription, error) {
	return sdk.client.QueueSubscribe(subjectDeviceFilter, groupDevice, handler)
}

//SubscribeSaveDeviceData ..
func (sdk *IotSDK) SubscribeSaveDeviceData(handler func(m *stan.Msg)) (stan.Subscription, error) {
	return sdk.client.QueueSubscribe(subjectDeviceData, groupDeviceSave, handler)
}

//SubscribePublishDeviceData ..
func (sdk *IotSDK) SubscribePublishDeviceData(handler func(m *stan.Msg)) (stan.Subscription, error) {
	return sdk.client.QueueSubscribe(subjectDeviceData, groupDevicePublish, handler)
}
