package services

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/9d77v/pdc/internal/db/mq"
	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/internal/module/device-service/chmodels"
	"github.com/9d77v/pdc/internal/module/device-service/models"
	devicePB "github.com/9d77v/pdc/internal/module/device-service/pb"
	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

//DeviceLogin ..
func (s DeviceService) DeviceLogin(accessKey, secretKey string) (uint, error) {
	m := models.NewDevice()
	err := m.Select([]string{"id", "access_key", "secret_key"}).
		Where("access_key=? and secret_key=?", accessKey, secretKey).First(m)
	return m.ID, err
}

//GetDeviceInfo ..
func (s DeviceService) GetDeviceInfo(deviceID uint32) (*pb.DeviceDownMSG, error) {
	replyDeviceMsg := new(pb.DeviceDownMSG)
	device := models.NewDevice()
	err := device.Preload("Attributes").
		Preload("Attributes.AttributeModel").
		Preload("Telemetries").
		Preload("Telemetries.TelemetryModel").
		Preload("DeviceModel").
		IDQuery(uint(deviceID)).First(device)
	if err != nil {
		log.Println("get device failed,err", err)
		return replyDeviceMsg, err
	}
	replyDeviceMsg.DeviceId = uint32(device.ID)
	attributeConfig := make(map[string]uint32, 0)
	for _, v := range device.Attributes {
		attributeConfig[v.AttributeModel.Key] = uint32(v.ID)
	}
	telemetryConfig := make(map[string]uint32)
	for _, v := range device.Telemetries {
		telemetryConfig[v.TelemetryModel.Key] = uint32(v.ID)
	}
	replyDeviceMsg.Payload = &pb.DeviceDownMSG_LoginReplyMsg{
		LoginReplyMsg: &pb.LoginReplyMsg{
			Id:              deviceID,
			Ip:              device.IP,
			Port:            uint32(device.Port),
			AttributeConfig: attributeConfig,
			TelemetryConfig: telemetryConfig,
			Username:        device.Username,
			Password:        device.Password,
			CameraCompany:   uint32(device.DeviceModel.CameraCompany),
		},
	}
	return replyDeviceMsg, nil
}

//CameraCapture ..
func (s DeviceService) CameraCapture(ctx context.Context,
	in *devicePB.CameraCaptureRequest) (*devicePB.CameraCaptureResponse, error) {
	resp := new(devicePB.CameraCaptureResponse)
	m := models.NewDevice()
	if s.RecordNotExist(m, uint(in.DeviceId)) {
		return resp, status.Error(codes.NotFound, "数据不存在")
	}
	deviceHealth := chmodels.NewDeviceHealth()
	deviceHealth.DeviceID = in.DeviceId
	if deviceHealth.IsUnhealthy() {
		return resp, nil
	}
	bucketName := "camera"
	now := time.Now()
	objectName := "picture/" + strconv.FormatUint(uint64(in.DeviceId), 10) + "/" + now.Format("2006-01-02") + "/" +
		strconv.FormatInt(now.Unix(), 10) + ".jpg"
	minioClient := oss.GetMinioClient()
	u, err := minioClient.PresignedPutObject(ctx, bucketName, objectName, 10*time.Minute)
	if err != nil {
		return resp, err
	}
	request := new(pb.DeviceDownMSG)
	request.DeviceId = in.DeviceId
	request.ActionTime = ptypes.TimestampNow()
	request.Payload = &pb.DeviceDownMSG_CameraCaptureMsg{
		CameraCaptureMsg: &pb.CameraCaptureMsg{
			PictureUrl:      u.String(),
			OssPrefix:       oss.OSSPrefix(),
			SecureOssPrefix: oss.SecureOSSPrefix(),
		},
	}
	requestMsg, err := proto.Marshal(request)
	if err != nil {
		return resp, err
	}
	subject := mq.SubjectDevicPrefix + strconv.FormatUint(uint64(in.DeviceId), 10)

	msg, err := mq.GetClient().Request(subject, requestMsg, 5*time.Second)
	if err != nil {
		return resp, err
	}
	deviceMsg := new(pb.DeviceUpMsg)
	err = proto.Unmarshal(msg.Data, deviceMsg)
	resp.ImageUrl = oss.GetOSSPrefixByScheme(in.Scheme) + u.Path
	return resp, err
}
