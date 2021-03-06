package sdk

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
	"github.com/golang/protobuf/ptypes"
	"github.com/gorilla/websocket"
	"github.com/robfig/cron/v3"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//IotSDK IotSDK
type IotSDK struct {
	conn       *websocket.Conn
	mutex      *sync.Mutex
	DeviceInfo *pb.LoginReplyMsg
}

//NewIotSDK init iot sdk
func NewIotSDK() *IotSDK {
	return &IotSDK{
		conn:  wsConn,
		mutex: new(sync.Mutex),
	}
}

//CronJob ..
type CronJob struct {
	Spec string
	Cmd  func()
}

//Run 运行iot数据采集程序
func (sdk *IotSDK) Run(works []func(), jobs []*CronJob, handleMsg func(iotSDK *IotSDK, msg []byte)) {
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
	if jobs == nil {
		jobs = make([]*CronJob, 0, 0)
	}
	jobs = append(jobs, &CronJob{
		Spec: "*/10 * * * * *",
		Cmd: func() {
			sdk.pingCheck()
		},
	})
	for _, job := range jobs {
		id, err := cr.AddFunc(job.Spec, job.Cmd)
		if err != nil {
			log.Println("cron add func error,entityID:", id, " error:", err)
		}
	}
	cr.Start()
	defer cr.Stop()
	go func() {
		if handleMsg != nil {
			for {
				_, msg, err := sdk.conn.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					break
				}
				handleMsg(sdk, msg)
			}
		}
	}()
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
	request := new(pb.DeviceUpMsg)
	loginMsg := new(pb.LoginMsg)
	loginMsg.AccessKey = accessKey
	loginMsg.SecretKey = secretKey
	request.ActionTime = ptypes.TimestampNow()
	request.Payload = &pb.DeviceUpMsg_LoginMsg{
		LoginMsg: loginMsg,
	}
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
	reply := new(pb.DeviceDownMSG)
	err = proto.Unmarshal(msg, reply)
	if err != nil {
		return err
	}
	fmt.Println(reply)
	if reply.Payload == nil {
		return errors.New("device info is nil")
	}
	sdk.DeviceInfo = reply.GetLoginReplyMsg()
	return nil
}

//SetDeviceAttributes set device attributes
func (sdk *IotSDK) SetDeviceAttributes(attributeMap map[uint32]string) {
	request := new(pb.DeviceUpMsg)
	request.DeviceId = sdk.DeviceInfo.Id
	request.ActionTime = ptypes.TimestampNow()
	request.Payload = &pb.DeviceUpMsg_SetAttributesMsg{
		SetAttributesMsg: &pb.SetAttributesMsg{
			AttributeMap: attributeMap,
		},
	}
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
	request := new(pb.DeviceUpMsg)
	request.DeviceId = sdk.DeviceInfo.Id
	request.ActionTime = now
	request.Payload = &pb.DeviceUpMsg_SetTelemetriesMsg{
		SetTelemetriesMsg: &pb.SetTelemetriesMsg{
			TelemetryMap: telemetryMap,
		},
	}
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
	request := new(pb.DeviceUpMsg)
	request.DeviceId = sdk.DeviceInfo.Id
	request.ActionTime = now
	request.Payload = &pb.DeviceUpMsg_SetHealthMsg{
		SetHealthMsg: &pb.SetHealthMsg{
			DeviceHealth: health,
		},
	}
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
	if sdk.ping(sdk.DeviceInfo.Ip) {
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

//ReplyCameraCapture ..
func (sdk *IotSDK) ReplyCameraCapture(subject string, ok bool) {
	request := new(pb.DeviceUpMsg)
	request.DeviceId = sdk.DeviceInfo.Id
	request.ActionTime = ptypes.TimestampNow()
	request.Payload = &pb.DeviceUpMsg_CameraCaptureReplyMsg{
		CameraCaptureReplyMsg: &pb.CameraCaptureReplyMsg{
			Subject: subject,
			Ok:      ok,
		},
	}
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

//SendPresignedURLMsg ..
func (sdk *IotSDK) SendPresignedURLMsg() {
	request := new(pb.DeviceUpMsg)
	presignedURLMsg := new(pb.PresignedUrlMsg)
	now := time.Now()
	presignedURLMsg.BucketName = "camera"
	presignedURLMsg.ObjectName = "picture/" + strconv.FormatInt(int64(sdk.DeviceInfo.Id), 10) +
		"/" + now.Format("2006-01-02") + "/tmp/" + strconv.FormatInt(now.Unix(), 10) + ".jpg"
	request.DeviceId = sdk.DeviceInfo.Id
	request.ActionTime = ptypes.TimestampNow()
	request.Payload = &pb.DeviceUpMsg_PresignedUrlMsg{
		PresignedUrlMsg: presignedURLMsg,
	}
	requestMsg, err := proto.Marshal(request)
	if err != nil {
		return
	}
	err = sdk.conn.WriteMessage(websocket.BinaryMessage, requestMsg)
	if err != nil {
		log.Println("send login msg:", err)
		return
	}
}

//PictureRequest ..
type PictureRequest struct {
	RequestURL      string
	OssPrefix       string
	SecureOssPrefix string
	Picture         []byte
}

//SavePicture ..
func (sdk *IotSDK) SavePicture(pr PictureRequest) error {
	if pr.Picture == nil {
		return errors.New("image should not be empty")
	}
	if strings.Contains(WSURL, "wss") {
		pr.RequestURL = strings.Replace(pr.RequestURL,
			pr.OssPrefix,
			pr.SecureOssPrefix, 1)
	}
	req, err := http.NewRequest("PUT", pr.RequestURL, bytes.NewReader(pr.Picture))
	if err != nil {
		return err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	return err
}
