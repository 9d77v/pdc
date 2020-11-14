package sdk

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/9d77v/pdc/iot/sdk/pb"
	"github.com/golang/protobuf/ptypes"
	"github.com/gorilla/websocket"
	cron "github.com/robfig/cron/v3"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//IotSDK IotSDK
type IotSDK struct {
	conn       *websocket.Conn
	mutex      *sync.Mutex
	DeviceInfo *pb.DeviceInfo
}

//NewIotSDK init iot sdk
func NewIotSDK() *IotSDK {
	return &IotSDK{
		conn:  wsConn,
		mutex: new(sync.Mutex),
	}
}

//Run 运行iot数据采集程序
func (sdk *IotSDK) Run(works []func()) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	defer sdk.conn.Close()
	done := make(chan struct{})
	err := sdk.login()
	if err != nil {
		return
	}
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for _, v := range works {
		v := v
		go v()
	}
	cr := cron.New(cron.WithSeconds())
	id, err := cr.AddFunc("*/10 * * * * *", func() {
		sdk.pingCheck()
	})
	if err != nil {
		log.Println("cron add func error,entityID:", id, " error:", err)
	}
	cr.Start()
	defer cr.Stop()
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			err := sdk.conn.WriteMessage(websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

//Login get device info by device id
func (sdk *IotSDK) login() error {
	request := new(pb.LoginMSG)
	request.AccessKey = accessKey
	request.SecretKey = secretKey
	requestMsg, err := proto.Marshal(request)
	if err != nil {
		return err
	}
	err = sdk.conn.WriteMessage(websocket.BinaryMessage, requestMsg)
	if err != nil {
		log.Println("send login msg:", err)
		return err
	}
	_, msg, err := sdk.conn.ReadMessage()
	if err != nil {
		log.Println("receive login reply:", err)
		return err
	}
	reply := new(pb.DeviceMSG)
	err = proto.Unmarshal(msg, reply)
	if err != nil {
		return err
	}
	fmt.Println(reply)
	if reply.DeviceInfo == nil {
		return errors.New("device info is nil")
	}
	sdk.DeviceInfo = reply.DeviceInfo
	return nil
}

//SetDeviceAttributes set device attributes
func (sdk *IotSDK) SetDeviceAttributes(attributeMap map[uint32]string) {
	request := new(pb.DeviceMSG)
	request.Action = pb.DeviceAction_SetAttributes
	request.DeviceID = sdk.DeviceInfo.ID
	request.AttributeMap = attributeMap
	request.ActionTime = ptypes.TimestampNow()
	requestMsg, err := proto.Marshal(request)
	if err != nil {
		log.Println("proto marshal error:", err)
		return
	}
	err = sdk.conn.WriteMessage(websocket.BinaryMessage, requestMsg)
	if err != nil {
		log.Printf("SetDeviceAttributes error:%v/n", err)
	}
}

//SetDeviceTelemetries upload device telemetries
func (sdk *IotSDK) SetDeviceTelemetries(telemetryMap map[uint32]float64, now *timestamppb.Timestamp) {
	request := new(pb.DeviceMSG)
	request.Action = pb.DeviceAction_SetTelemetries
	request.DeviceID = sdk.DeviceInfo.ID
	request.TelemetryMap = telemetryMap
	request.ActionTime = now
	requestMsg, err := proto.Marshal(request)
	if err != nil {
		log.Println("proto marshal error:", err)
		return
	}
	err = sdk.conn.WriteMessage(websocket.BinaryMessage, requestMsg)
	if err != nil {
		log.Panicf("SetDeviceTelemetries error:%v/n", err)
	}
}

func (sdk *IotSDK) setDeviceHealth(health uint32, now *timestamppb.Timestamp) {
	request := new(pb.DeviceMSG)
	request.Action = pb.DeviceAction_SetHealth
	request.DeviceID = sdk.DeviceInfo.ID
	request.DeviceHealth = health
	request.ActionTime = now
	requestMsg, err := proto.Marshal(request)
	if err != nil {
		log.Println("proto marshal error:", err)
		return
	}
	err = sdk.conn.WriteMessage(websocket.BinaryMessage, requestMsg)
	if err != nil {
		log.Panicf("SetDeviceHealth error:%v/n", err)
	}
}

func (sdk *IotSDK) pingCheck() {
	if sdk.DeviceInfo == nil {
		return
	}
	now := ptypes.TimestampNow()
	var flag uint32
	if sdk.ping(sdk.DeviceInfo.IP) {
		flag = 1
	}
	sdk.setDeviceHealth(flag, now)
}

func (sdk *IotSDK) ping(ip string) bool {
	var buf bytes.Buffer
	var errorBuf bytes.Buffer
	cmd := exec.Command("ping", "-i", "1", "-c", "3", ip)
	cmd.Stdout = &buf
	cmd.Stderr = &errorBuf
	err := cmd.Run()
	if err != nil {
		log.Println("ping failed:", err)
	}
	if buf.String() != "" {
		data := buf.String()
		dataArr := strings.Split(data, "\n")
		if len(dataArr) != 9 {
			return false
		}
		statisticsStr := dataArr[6]
		statisticsArr := strings.Split(statisticsStr, ",")
		if len(statisticsArr) < 3 {
			return false
		}
		return statisticsArr[2] == " 0% packet loss"
	}
	return false
}
