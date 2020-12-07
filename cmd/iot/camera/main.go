package main

import (
	"log"

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
	cronJobs := make([]*sdk.CronJob, 0)
	cronJobs = append(cronJobs, &sdk.CronJob{
		Spec: "*/15 * * * * *",
		Cmd: func() {
			iotSDK.SendPresignedURLMsg()
		},
	})
	iotSDK.Run(works, cronJobs, cameraMsgHandler)
}

func updateDeviceAttributes(iotSDK *sdk.IotSDK) {
	device := iotSDK.DeviceInfo
	cameraDeviceMap := make(map[string]string, 0)

	c := createCameraer(device)
	cameraDeviceMap = c.GetDeviceInfo()
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

func cameraMsgHandler(iotSDK *sdk.IotSDK, msg []byte) {
	deviceMsg := new(pb.DeviceDownMSG)
	err := proto.Unmarshal(msg, deviceMsg)
	if err != nil {
		log.Println("unmarshal data error")
		return
	}
	switch deviceMsg.Payload.(type) {
	case *pb.DeviceDownMSG_CameraCaptureMsg:
		msg := deviceMsg.GetCameraCaptureMsg()
		handleCameraCaptureMsg(iotSDK, msg)
	case *pb.DeviceDownMSG_PresignedUrlReplyMsg:
		msg := deviceMsg.GetPresignedUrlReplyMsg()
		handlePresignedURLReplyMsg(iotSDK, msg)
	}
}

func handleCameraCaptureMsg(iotSDK *sdk.IotSDK, msg *pb.CameraCaptureMsg) {
	device := iotSDK.DeviceInfo
	c := createCameraer(device)
	picture := c.Capture()
	captureOk := false
	err := iotSDK.SavePicture(sdk.PictureRequest{
		RequestURL:      msg.GetPictureUrl(),
		OssPrefix:       msg.GetOssPrefix(),
		SecureOssPrefix: msg.GetSecureOssPrefix(),
		Picture:         picture,
	})
	if err != nil {
		log.Println("Save Picture failed, ", err)
	} else {
		captureOk = true
	}
	iotSDK.ReplyCameraCapture(msg.NatsReply, captureOk)
}

func handlePresignedURLReplyMsg(iotSDK *sdk.IotSDK, msg *pb.PresignedUrlReplyMsg) {
	device := iotSDK.DeviceInfo
	c := createCameraer(device)
	picture := c.Capture()
	err := iotSDK.SavePicture(sdk.PictureRequest{
		RequestURL:      msg.GetPictureUrl(),
		OssPrefix:       msg.GetOssPrefix(),
		SecureOssPrefix: msg.GetSecureOssPrefix(),
		Picture:         picture,
	})
	if err != nil {
		log.Println("Save Picture failed, ", err)
	}
}

func createCameraer(device *pb.LoginReplyMsg) camera.Cameraer {
	var c camera.Cameraer
	switch device.CameraCompany {
	case camera.CameraCompanyHikvision:
		c = camera.NewHikvision(device.Ip, device.Username, device.Password)
	case camera.CameraCompanyDaHua:
		c = camera.NewDahua(device.Ip, device.Username, device.Password)
	}
	return c
}
