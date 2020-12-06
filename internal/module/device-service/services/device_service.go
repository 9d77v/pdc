package services

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/9d77v/go-lib/ptrs"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"

	"github.com/9d77v/pdc/internal/consts"
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/db/mq"
	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/device-service/models"
	"github.com/9d77v/pdc/internal/utils"
	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
)

//DeviceService ..
type DeviceService struct {
	base.Service
}

//CreateDeviceModel  ..
func (s DeviceService) CreateDeviceModel(ctx context.Context,
	input model.NewDeviceModel) (*model.DeviceModel, error) {
	m := &models.DeviceModel{
		Name:          input.Name,
		Desc:          ptrs.String(input.Desc),
		DeviceType:    uint8(input.DeviceType),
		CameraCompany: uint8(input.CameraCompany),
	}
	err := db.GetDB().Create(m).Error
	if err != nil {
		return &model.DeviceModel{}, err
	}
	if input.DeviceType == 1 && input.CameraCompany == 0 {
		attributes := getDefaultHikvisionAttributeModels(m.ID)
		err = db.GetDB().Create(&attributes).Error
		if err != nil {
			return &model.DeviceModel{ID: int64(m.ID)}, err
		}
	}
	return &model.DeviceModel{ID: int64(m.ID)}, err
}

//UpdateDeviceModel ..
func (s DeviceService) UpdateDeviceModel(ctx context.Context,
	input model.NewUpdateDeviceModel) (*model.DeviceModel, error) {
	deviceModel := models.NewDeviceModel()
	fields := s.GetInputFields(ctx)
	if err := deviceModel.GetByID(uint(input.ID), fields); err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"name": input.Name,
		"desc": ptrs.String(input.Desc),
	}
	err := db.GetDB().Model(deviceModel).Updates(updateMap).Error
	return &model.DeviceModel{ID: int64(deviceModel.ID)}, err
}

//CreateAttributeModel  ..
func (s DeviceService) CreateAttributeModel(ctx context.Context, input model.NewAttributeModel) (*model.AttributeModel, error) {
	m := &models.AttributeModel{
		DeviceModelID: uint(input.DeviceModelID),
		Key:           input.Key,
		Name:          input.Name,
	}
	err := db.GetDB().Create(m).Error
	if err != nil {
		return &model.AttributeModel{}, err
	}
	return &model.AttributeModel{ID: int64(m.ID)}, err
}

//UpdateAttributeModel ..
func (s DeviceService) UpdateAttributeModel(ctx context.Context, input model.NewUpdateAttributeModel) (*model.AttributeModel, error) {
	attributeModel := models.NewAttributeModel()
	fields := s.GetInputFields(ctx)
	if err := attributeModel.GetByID(uint(input.ID), fields); err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"name": input.Name,
	}
	err := db.GetDB().Model(attributeModel).Updates(updateMap).Error
	return &model.AttributeModel{ID: int64(attributeModel.ID)}, err
}

//DeleteAttributeModel ..
func (s DeviceService) DeleteAttributeModel(ctx context.Context,
	id int64) (*model.AttributeModel, error) {
	m := new(models.AttributeModel)
	m.ID = uint(id)
	err := db.GetDB().Delete(m).Error
	return &model.AttributeModel{ID: int64(m.ID)}, err
}

//CreateTelemetryModel  ..
func (s DeviceService) CreateTelemetryModel(ctx context.Context,
	input model.NewTelemetryModel) (*model.TelemetryModel, error) {
	m := &models.TelemetryModel{
		DeviceModelID: uint(input.DeviceModelID),
		Key:           input.Key,
		Name:          input.Name,
		Factor:        input.Factor,
		Unit:          ptrs.String(input.Unit),
		UnitName:      ptrs.String(input.UnitName),
		Scale:         uint8(input.Scale),
	}
	err := db.GetDB().Create(m).Error
	if err != nil {
		return &model.TelemetryModel{}, err
	}
	return &model.TelemetryModel{ID: int64(m.ID)}, err
}

//UpdateTelemetryModel ..
func (s DeviceService) UpdateTelemetryModel(ctx context.Context,
	input model.NewUpdateTelemetryModel) (*model.TelemetryModel, error) {
	telemetryModel := models.NewTelemetryModel()
	fields := s.GetInputFields(ctx)
	if err := telemetryModel.GetByID(uint(input.ID), fields); err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"name":      input.Name,
		"factor":    input.Factor,
		"unit":      input.Unit,
		"unit_name": input.UnitName,
		"scale":     input.Scale,
	}
	err := db.GetDB().Model(telemetryModel).Updates(updateMap).Error
	return &model.TelemetryModel{ID: int64(telemetryModel.ID)}, err
}

//DeleteTelemetryModel ..
func (s DeviceService) DeleteTelemetryModel(ctx context.Context,
	id int64) (*model.TelemetryModel, error) {
	m := new(models.TelemetryModel)
	m.ID = uint(id)
	err := db.GetDB().Delete(m).Error
	return &model.TelemetryModel{ID: int64(m.ID)}, err
}

//ListDeviceModel ..
func (s DeviceService) ListDeviceModel(ctx context.Context, searchParam model.SearchParam) (int64, []*model.DeviceModel, error) {
	deviceModel := models.NewDeviceModel()
	deviceModel.FuzzyQuery(searchParam.Keyword, "name")
	replaceFunc := func(edgeFieldMap map[string]bool, edgeFields []string) error {
		if edgeFieldMap["attributeModels"] {
			deviceModel.Preload("AttributeModels", func(db *gorm.DB) *gorm.DB {
				return db.Model(&models.AttributeModel{})
			})
		}
		if edgeFieldMap["telemetryModels"] {
			deviceModel.Preload("TelemetryModels", func(db *gorm.DB) *gorm.DB {
				return db.Model(&models.TelemetryModel{})
			})
		}
		return nil
	}
	data := make([]*models.DeviceModel, 0)
	total, err := s.GetConnection(ctx, deviceModel, searchParam, &data, replaceFunc,
		"attributeModels", "telemetryModels")
	return total, toDeviceModelDtos(data), err
}

//CreateDevice  ..
func (s DeviceService) CreateDevice(ctx context.Context, input model.NewDevice) (*model.Device, error) {
	m := &models.Device{
		DeviceModelID: uint(input.DeviceModelID),
		Name:          input.Name,
		IP:            ptrs.String(input.IP),
		Port:          uint16(ptrs.Int64(input.Port)),
		Username:      ptrs.String(input.Username),
		Password:      ptrs.String(input.Password),
	}
	deviceModel := new(models.DeviceModel)
	if err := db.GetDB().Preload("AttributeModels", func(db *gorm.DB) *gorm.DB {
		return db.Model(&models.AttributeModel{})
	}).Preload("TelemetryModels", func(db *gorm.DB) *gorm.DB {
		return db.Model(&models.TelemetryModel{})
	}).First(deviceModel, "id=?", input.DeviceModelID).Error; err != nil {
		return nil, err
	}
	tx := db.GetDB().Begin()
	err := tx.Create(m).Error
	if err != nil {
		tx.Rollback()
		return &model.Device{}, err
	}
	m.AccessKey = consts.GetDeviceAccessKey(m.ID)
	m.SecretKey = consts.GetDeviceSecretKey()
	err = tx.Save(m).Error
	if err != nil {
		tx.Rollback()
		return &model.Device{}, err
	}
	attributes := make([]*models.Attribute, 0, len(deviceModel.AttributeModels))
	for _, v := range deviceModel.AttributeModels {
		attributes = append(attributes, &models.Attribute{
			DeviceID:         m.ID,
			AttributeModelID: v.ID,
		})
	}
	if len(attributes) > 0 {
		err = tx.Create(&attributes).Error
		if err != nil {
			tx.Rollback()
			return &model.Device{}, err
		}
	}
	telemetries := make([]*models.Telemetry, 0, len(deviceModel.TelemetryModels))
	for _, v := range deviceModel.TelemetryModels {
		telemetries = append(telemetries, &models.Telemetry{
			DeviceID:         m.ID,
			TelemetryModelID: v.ID,
		})

	}
	if len(telemetries) > 0 {
		err = tx.Create(&telemetries).Error
		if err != nil {
			tx.Rollback()
			return &model.Device{}, err
		}
	}
	tx.Commit()
	return &model.Device{ID: int64(m.ID)}, err
}

//UpdateDevice ..
func (s DeviceService) UpdateDevice(ctx context.Context, input model.NewUpdateDevice) (*model.Device, error) {
	m := new(models.Device)
	if err := db.GetDB().Select("id").First(m, "id=?", input.ID).Error; err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"name":     input.Name,
		"ip":       ptrs.String(input.IP),
		"port":     uint(ptrs.Int64(input.Port)),
		"username": ptrs.String(input.Username),
		"password": ptrs.String(input.Password),
	}
	err := db.GetDB().Model(m).Updates(updateMap).Error
	return &model.Device{ID: int64(m.ID)}, err
}

//ListDevice ..
func (s DeviceService) ListDevice(ctx context.Context, searchParam model.SearchParam, deviceType *int64) (int64, []*model.Device, error) {
	device := models.NewDevice()
	device.FuzzyQuery(searchParam.Keyword, "name")
	if deviceType != nil {
		tableDeviceModel := new(models.DeviceModel).TableName()
		device.
			Where(tableDeviceModel+".device_type = ?", ptrs.Int64(deviceType)).
			LeftJoin(tableDeviceModel + " ON " + tableDeviceModel + ".id = " +
				device.TableName() + ".device_model_id")
	}
	omitFields := []string{"attributes", "telemetries",
		"deviceModelName", "deviceModelDesc",
		"deviceModelDeviceType", "deviceModelCameraCompany"}
	replaceFunc := func(edgeFieldMap map[string]bool, edgeFields []string) error {
		if deviceType != nil {
			device.SelectWithPrefix(edgeFields, device.TableName()+".", omitFields...)
		}
		if edgeFieldMap["attributes"] {
			device.Preload("Attributes").Preload("Attributes.AttributeModel")
		}
		if edgeFieldMap["telemetries"] {
			device.Preload("Telemetries").Preload("Telemetries.TelemetryModel")
		}
		if edgeFieldMap["deviceModelName"] || edgeFieldMap["deviceModelDesc"] ||
			edgeFieldMap["deviceModelDeviceType"] || edgeFieldMap["deviceModelCameraCompany"] {
			device.Preload("DeviceModel")
		}
		return nil
	}
	data := make([]*models.Device, 0)
	total, err := s.GetConnection(ctx, device, searchParam, &data, replaceFunc, omitFields...)
	return total, toDeviceDtos(data), err
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

//DeviceLogin ..
func (s DeviceService) DeviceLogin(accessKey, secretKey string) (uint, error) {
	device := new(models.Device)
	err := db.GetDB().Select("id,access_key,secret_key").
		Where("access_key=? and secret_key=?", accessKey, secretKey).First(device).Error
	if err != nil {
		log.Println("get device failed,err", err)
		return 0, err
	}
	return device.ID, err
}

//CreateDeviceDashboard  ..
func (s DeviceService) CreateDeviceDashboard(ctx context.Context,
	input model.NewDeviceDashboard) (*model.DeviceDashboard, error) {
	m := &models.DeviceDashboard{
		Name:       input.Name,
		IsVisible:  input.IsVisible,
		DeviceType: uint8(input.DeviceType),
	}
	err := db.GetDB().Create(m).Error
	if err != nil {
		return &model.DeviceDashboard{}, err
	}
	return &model.DeviceDashboard{ID: int64(m.ID)}, err
}

//UpdateDeviceDashboard ..
func (s DeviceService) UpdateDeviceDashboard(ctx context.Context, input model.NewUpdateDeviceDashboard) (*model.DeviceDashboard, error) {
	deviceDashboard := models.NewDeviceDashboard()
	fields := s.GetInputFields(ctx)
	if err := deviceDashboard.GetByID(uint(input.ID), fields); err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"name":       input.Name,
		"is_visible": input.IsVisible,
	}
	err := db.GetDB().Model(deviceDashboard).Updates(updateMap).Error
	return &model.DeviceDashboard{ID: int64(deviceDashboard.ID)}, err
}

//DeleteDeviceDashboard ..
func (s DeviceService) DeleteDeviceDashboard(ctx context.Context, id int64) (*model.DeviceDashboard, error) {
	m := new(models.DeviceDashboard)
	m.ID = uint(id)
	err := db.GetDB().Delete(m).Error
	return &model.DeviceDashboard{ID: int64(m.ID)}, err
}

//AddDeviceDashboardTelemetry  ..
func (s DeviceService) AddDeviceDashboardTelemetry(ctx context.Context, input model.NewDeviceDashboardTelemetry) (*model.DeviceDashboard, error) {
	if len(input.TelemetryIDs) == 0 {
		return &model.DeviceDashboard{}, nil
	}
	data := make([]*models.DeviceDashboardTelemetry, 0, len(input.TelemetryIDs))
	for _, v := range input.TelemetryIDs {
		m := &models.DeviceDashboardTelemetry{
			DeviceDashboardID: uint(input.DeviceDashboardID),
			TelemetryID:       uint(v),
		}
		data = append(data, m)
	}
	err := db.GetDB().Create(data).Error
	if err != nil {
		return &model.DeviceDashboard{}, err
	}
	return &model.DeviceDashboard{ID: input.DeviceDashboardID}, err
}

//RemoveDeviceDashboardTelemetry ..
func (s DeviceService) RemoveDeviceDashboardTelemetry(ctx context.Context,
	ids []int64) (*model.DeviceDashboard, error) {
	m := new(models.DeviceDashboardTelemetry)
	err := db.GetDB().Delete(m, ids).Error
	return &model.DeviceDashboard{ID: int64(m.ID)}, err
}

//AddDeviceDashboardCamera  ..
func (s DeviceService) AddDeviceDashboardCamera(ctx context.Context,
	input model.NewDeviceDashboardCamera) (*model.DeviceDashboard, error) {
	if len(input.DeviceIDs) == 0 {
		return &model.DeviceDashboard{}, nil
	}
	data := make([]*models.DeviceDashboardCamera, 0, len(input.DeviceIDs))
	for _, v := range input.DeviceIDs {
		m := &models.DeviceDashboardCamera{
			DeviceDashboardID: uint(input.DeviceDashboardID),
			DeviceID:          uint(v),
		}
		data = append(data, m)
	}
	err := db.GetDB().Create(data).Error
	if err != nil {
		return &model.DeviceDashboard{}, err
	}
	return &model.DeviceDashboard{ID: input.DeviceDashboardID}, err
}

//RemoveDeviceDashboardCamera ..
func (s DeviceService) RemoveDeviceDashboardCamera(ctx context.Context,
	ids []int64) (*model.DeviceDashboard, error) {
	m := new(models.DeviceDashboardCamera)
	err := db.GetDB().Delete(m, ids).Error
	return &model.DeviceDashboard{ID: int64(m.ID)}, err
}

//ListDeviceDashboards ..
func (s DeviceService) ListDeviceDashboards(ctx context.Context, searchParam model.SearchParam) (int64, []*model.DeviceDashboard, error) {
	deviceDashboard := models.NewDeviceDashboard()
	deviceDashboard.FuzzyQuery(searchParam.Keyword, "name")
	replaceFunc := func(edgeFieldMap map[string]bool, edgeFields []string) error {
		if edgeFieldMap["telemetries"] {
			deviceDashboard.Preload("Telemetries").
				Preload("Telemetries.Telemetry").
				Preload("Telemetries.Telemetry.TelemetryModel")
			telemetryFieldMap, _ := utils.GetFieldData(ctx, "edges.telemetries.")
			if telemetryFieldMap["deviceName"] {
				deviceDashboard.Preload("Telemetries.Telemetry.Device")
			}
		}
		if edgeFieldMap["cameras"] {
			deviceDashboard.Preload("Cameras")
			cameraFieldMap, _ := utils.GetFieldData(ctx, "edges.cameras.")
			if cameraFieldMap["deviceName"] {
				deviceDashboard.Preload("Cameras.Device")
			}
		}
		return nil
	}
	data := make([]*models.DeviceDashboard, 0)
	total, err := s.GetConnection(ctx, deviceDashboard, searchParam, &data, replaceFunc,
		"telemetries", "cameras")
	return total, toDeviceDashboardDtos(data), err
}

//AppDeviceDashboards ..
func (s DeviceService) AppDeviceDashboards(ctx context.Context,
	deviceType *int64) (int64, []*model.DeviceDashboard, error) {
	result := make([]*model.DeviceDashboard, 0)
	data := make([]*models.DeviceDashboard, 0)
	fieldMap, _ := utils.GetFieldData(ctx, "")
	deviceDashboard := models.NewDeviceDashboard()
	if deviceType != nil {
		deviceDashboard.Where("device_type = ?", ptrs.Int64(deviceType))
	}
	deviceDashboard.Where("is_visible = true")
	var total int64
	if fieldMap["totalCount"] {
		total = int64(len(data))
	}
	if fieldMap["edges"] {
		edgeFieldMap, edgeFields := utils.GetFieldData(ctx, "edges.")
		if edgeFieldMap["telemetries"] {
			deviceDashboard.Preload("Telemetries").
				Preload("Telemetries.Telemetry").
				Preload("Telemetries.Telemetry.TelemetryModel")
			telemetryFieldMap, _ := utils.GetFieldData(ctx, "edges.telemetries.")
			if telemetryFieldMap["deviceName"] {
				deviceDashboard.Preload("Telemetries.Telemetry.Device")
			}
		}
		if edgeFieldMap["cameras"] {
			deviceDashboard.Preload("Cameras")
			cameraFieldMap, _ := utils.GetFieldData(ctx, "edges.cameras.")
			if cameraFieldMap["deviceName"] {
				deviceDashboard.Preload("Cameras.Device")
			}
		}
		err := deviceDashboard.Select(edgeFields, "telemetries", "cameras").Find(&data)
		if err != nil {
			return 0, result, err
		}
	}
	return total, toDeviceDashboardDtos(data), nil
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

//CameraTimeLapseVideos ..
func (s DeviceService) CameraTimeLapseVideos(ctx context.Context,
	deviceID int64, scheme string) (int64, []*model.CameraTimeLapseVideo, error) {
	result := make([]*model.CameraTimeLapseVideo, 0)
	data := make([]*models.CameraTimeLapseVideo, 0)
	fieldMap, _ := utils.GetFieldData(ctx, "")
	camera := models.NewCameraTimeLapseVideo()
	camera.IDQuery(uint(deviceID), "device_id").
		Where("created_at>=?", time.Now().AddDate(0, 0, -7))
	var total int64
	if fieldMap["totalCount"] {
		total = int64(len(data))
	}
	if fieldMap["edges"] {
		_, edgeFields := utils.GetFieldData(ctx, "edges.")
		err := camera.Select(edgeFields).
			Order("id DESC").
			Find(&data)
		if err != nil {
			return 0, result, err
		}
	}
	return total, toCameraTimeLapseVideoDtos(data, scheme), nil
}
