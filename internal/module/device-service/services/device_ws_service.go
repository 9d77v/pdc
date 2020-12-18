package services

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/9d77v/pdc/internal/db/mq"
	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/internal/module/device-service/models"
	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
	"github.com/golang/protobuf/ptypes"
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
	replyDeviceMsg := &pb.DeviceDownMSG{
		DeviceId: deviceID,
	}
	device := models.NewDevice()
	err := device.Preload("Attributes").
		Preload("Attributes.AttributeModel").
		Preload("Telemetries").
		Preload("Telemetries.TelemetryModel").
		Preload("DeviceModel").
		IDQuery(uint(replyDeviceMsg.DeviceId)).First(device)
	if err != nil {
		log.Println("get device failed,err", err)
		return replyDeviceMsg, err
	}
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
func (s DeviceService) CameraCapture(ctx context.Context, deviceID int64, scheme string) (string, error) {
	bucketName := "camera"
	now := time.Now()
	objectName := "picture/" + strconv.FormatInt(deviceID, 10) + "/" + now.Format("2006-01-02") + "/" +
		strconv.FormatInt(now.Unix(), 10) + ".jpg"
	minioClient := oss.GetMinioClient()
	u, err := minioClient.PresignedPutObject(ctx, bucketName, objectName, 10*time.Minute)
	if err != nil {
		return "", err
	}
	request := new(pb.DeviceDownMSG)
	request.DeviceId = uint32(deviceID)
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
		log.Println("proto marshal error:", err)
		return "", nil
	}
	subject := mq.SubjectDevicPrefix + strconv.FormatUint(uint64(deviceID), 10)
	log.Println("发送数据到主题", subject)
	msg, err := mq.GetClient().NatsConn().Request(subject, requestMsg, 5*time.Second)
	if err != nil {
		log.Println("send data error:", err)
		return oss.GetOSSPrefixByScheme(scheme) + u.Path, nil
	}
	deviceMsg := new(pb.DeviceUpMsg)
	err = proto.Unmarshal(msg.Data, deviceMsg)
	if err != nil {
		log.Println("unmarshal data error")
	}
	return oss.GetOSSPrefixByScheme(scheme) + u.Path, nil
}
