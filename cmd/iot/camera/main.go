package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"google.golang.org/protobuf/proto"

	"github.com/9d77v/pdc/pkg/iot/camera"
	"github.com/9d77v/pdc/pkg/iot/sdk"
	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
)

func main() {
	iotSDK := sdk.NewIotSDK()
	works := []func(){
		func() {
			updateDeviceAttributes(iotSDK)
		},
	}
	iotSDK.Run(works, func(msg []byte) {
		deviceMsg := new(pb.DeviceDownMSG)
		err := proto.Unmarshal(msg, deviceMsg)
		if err != nil {
			log.Println("unmarshal data error")
			return
		}
		switch deviceMsg.Payload.(type) {
		case *pb.DeviceDownMSG_CameraCaptureMsg:
			msg := deviceMsg.GetCameraCaptureMsg()
			cameraCapture(iotSDK, msg)
		}
	})
}

func updateDeviceAttributes(iotSDK *sdk.IotSDK) {
	device := iotSDK.DeviceInfo
	cameraDeviceMap := camera.GetDeviceInfo(int(device.CameraCompany), device.Ip, device.Username, device.Password)
	attributeMap := make(map[uint32]string, 0)
	for k, v := range cameraDeviceMap {
		if device.AttributeConfig[k] != 0 {
			attributeMap[device.AttributeConfig[k]] = v
		}
	}
	if len(attributeMap) > 0 {
		iotSDK.SetDeviceAttributes(attributeMap)
	}
}

//cameraCapture get picture from camera,then send to minio
func cameraCapture(iotSDK *sdk.IotSDK, msg *pb.CameraCaptureMsg) {
	device := iotSDK.DeviceInfo
	imageByte := camera.Capture(int(device.CameraCompany), device.Ip, device.Username, device.Password)
	if imageByte != nil {
		requestURL := msg.GetPictureUrl()
		if strings.Contains(sdk.WSURL, "wss") {
			requestURL = strings.Replace(requestURL,
				msg.GetOssPrefix(),
				msg.GetSecureOssPrefix(), 1)
		}
		req, err := http.NewRequest("PUT", requestURL, bytes.NewReader(imageByte))
		if err != nil {
			log.Println("build http request faield,err:", err)
			iotSDK.ReplyCameraCapture(msg.NatsReply, false)
			return
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Println("http request faield,err:", err)
			iotSDK.ReplyCameraCapture(msg.NatsReply, false)
			return
		}
		defer res.Body.Close()
		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println("read body faield,err:", err)
			iotSDK.ReplyCameraCapture(msg.NatsReply, false)
			return
		}
		iotSDK.ReplyCameraCapture(msg.NatsReply, true)
	} else {
		iotSDK.ReplyCameraCapture(msg.NatsReply, false)
	}
}
