package services

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/9d77v/go-lib/ptrs"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/protobuf/proto"
	"gorm.io/gorm"

	"github.com/9d77v/pdc/internal/consts"
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/db/mq"
	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/device-service/models"
	"github.com/9d77v/pdc/internal/utils"
	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
)

//DeviceService ..
type DeviceService struct {
}

//CreateDeviceModel  ..
func (s DeviceService) CreateDeviceModel(ctx context.Context, input model.NewDeviceModel) (*model.DeviceModel, error) {
	m := &models.DeviceModel{
		Name:          input.Name,
		Desc:          ptrs.String(input.Desc),
		DeviceType:    uint8(input.DeviceType),
		CameraCompany: uint8(input.CameraCompany),
	}
	err := db.Gorm.Create(m).Error
	if err != nil {
		return &model.DeviceModel{}, err
	}
	if input.DeviceType == 1 && input.CameraCompany == 0 {
		attributes := []*models.AttributeModel{
			{
				DeviceModelID: uint(m.ID),
				Key:           "device_name",
				Name:          "设备名称",
			},
			{
				DeviceModelID: uint(m.ID),
				Key:           "device_id",
				Name:          "设备ID",
			},
			{
				DeviceModelID: uint(m.ID),
				Key:           "device_description",
				Name:          "设备描述",
			},
			{
				DeviceModelID: uint(m.ID),
				Key:           "device_location",
				Name:          "设备位置",
			},
			{
				DeviceModelID: uint(m.ID),
				Key:           "system_contact",
				Name:          "系统联系方",
			},
			{
				DeviceModelID: uint(m.ID),
				Key:           "model",
				Name:          "类型",
			},
			{
				DeviceModelID: uint(m.ID),
				Key:           "serial_number",
				Name:          "序列号",
			},
			{
				DeviceModelID: uint(m.ID),
				Key:           "mac_address",
				Name:          "MAC地址",
			},
			{
				DeviceModelID: uint(m.ID),
				Key:           "firmware_version",
				Name:          "固件版本",
			},
			{
				DeviceModelID: uint(m.ID),
				Key:           "firmware_releasedDate",
				Name:          "固件发布日期",
			},
			{
				DeviceModelID: uint(m.ID),
				Key:           "encoder_version",
				Name:          "编码器版本",
			},
			{
				DeviceModelID: uint(m.ID),
				Key:           "encoder_released_date",
				Name:          "编码器发布日期",
			},
			{
				DeviceModelID: uint(m.ID),
				Key:           "boot_version",
				Name:          "引导版本",
			},
			{
				DeviceModelID: uint(m.ID),
				Key:           "boot_released_date",
				Name:          "引导发布日期",
			},
			{
				DeviceModelID: uint(m.ID),
				Key:           "hardware_version",
				Name:          "硬件版本",
			},
			{
				DeviceModelID: uint(m.ID),
				Key:           "device_type",
				Name:          "设备类型",
			},
			{
				DeviceModelID: uint(m.ID),
				Key:           "telecontrol_id",
				Name:          "远程 ID",
			},
			{
				DeviceModelID: uint(m.ID),
				Key:           "support beep",
				Name:          "支持蜂鸣音",
			},
			{
				DeviceModelID: uint(m.ID),
				Key:           "support_video_loss",
				Name:          "支持视频丢失",
			},
			{
				DeviceModelID: uint(m.ID),
				Key:           "firmware_version_info",
				Name:          "固件版本信息",
			},
		}
		err = db.Gorm.Create(&attributes).Error
		if err != nil {
			return &model.DeviceModel{ID: int64(m.ID)}, err
		}
	}
	return &model.DeviceModel{ID: int64(m.ID)}, err
}

//UpdateDeviceModel ..
func (s DeviceService) UpdateDeviceModel(ctx context.Context, input model.NewUpdateDeviceModel) (*model.DeviceModel, error) {
	m := new(models.DeviceModel)
	varibales := graphql.GetRequestContext(ctx).Variables
	fields := make([]string, 0)
	for k := range varibales["input"].(map[string]interface{}) {
		fields = append(fields, k)
	}
	if err := db.Gorm.Select(utils.ToDBFields(fields)).First(m, "id=?", input.ID).Error; err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"name": input.Name,
		"desc": ptrs.String(input.Desc),
	}
	err := db.Gorm.Model(m).Updates(updateMap).Error
	return &model.DeviceModel{ID: int64(m.ID)}, err
}

//CreateAttributeModel  ..
func (s DeviceService) CreateAttributeModel(ctx context.Context, input model.NewAttributeModel) (*model.AttributeModel, error) {
	m := &models.AttributeModel{
		DeviceModelID: uint(input.DeviceModelID),
		Key:           input.Key,
		Name:          input.Name,
	}
	err := db.Gorm.Create(m).Error
	if err != nil {
		return &model.AttributeModel{}, err
	}
	return &model.AttributeModel{ID: int64(m.ID)}, err
}

//UpdateAttributeModel ..
func (s DeviceService) UpdateAttributeModel(ctx context.Context, input model.NewUpdateAttributeModel) (*model.AttributeModel, error) {
	m := new(models.AttributeModel)
	varibales := graphql.GetRequestContext(ctx).Variables
	fields := make([]string, 0)
	for k := range varibales["input"].(map[string]interface{}) {
		fields = append(fields, k)
	}
	if err := db.Gorm.Select(utils.ToDBFields(fields)).First(m, "id=?", input.ID).Error; err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"name": input.Name,
	}
	err := db.Gorm.Model(m).Updates(updateMap).Error
	return &model.AttributeModel{ID: int64(m.ID)}, err
}

//DeleteAttributeModel ..
func (s DeviceService) DeleteAttributeModel(ctx context.Context, id int64) (*model.AttributeModel, error) {
	m := new(models.AttributeModel)
	m.ID = uint(id)
	err := db.Gorm.Delete(m).Error
	return &model.AttributeModel{ID: int64(m.ID)}, err
}

//CreateTelemetryModel  ..
func (s DeviceService) CreateTelemetryModel(ctx context.Context, input model.NewTelemetryModel) (*model.TelemetryModel, error) {
	m := &models.TelemetryModel{
		DeviceModelID: uint(input.DeviceModelID),
		Key:           input.Key,
		Name:          input.Name,
		Factor:        input.Factor,
		Unit:          ptrs.String(input.Unit),
		UnitName:      ptrs.String(input.UnitName),
		Scale:         uint8(input.Scale),
	}
	err := db.Gorm.Create(m).Error
	if err != nil {
		return &model.TelemetryModel{}, err
	}
	return &model.TelemetryModel{ID: int64(m.ID)}, err
}

//UpdateTelemetryModel ..
func (s DeviceService) UpdateTelemetryModel(ctx context.Context, input model.NewUpdateTelemetryModel) (*model.TelemetryModel, error) {
	m := new(models.TelemetryModel)
	varibales := graphql.GetRequestContext(ctx).Variables
	fields := make([]string, 0)
	for k := range varibales["input"].(map[string]interface{}) {
		fields = append(fields, k)
	}
	if err := db.Gorm.Select(utils.ToDBFields(fields)).First(m, "id=?", input.ID).Error; err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"name":      input.Name,
		"factor":    input.Factor,
		"unit":      input.Unit,
		"unit_name": input.UnitName,
		"scale":     input.Scale,
	}
	err := db.Gorm.Model(m).Updates(updateMap).Error
	return &model.TelemetryModel{ID: int64(m.ID)}, err
}

//DeleteTelemetryModel ..
func (s DeviceService) DeleteTelemetryModel(ctx context.Context, id int64) (*model.TelemetryModel, error) {
	m := new(models.TelemetryModel)
	m.ID = uint(id)
	err := db.Gorm.Delete(m).Error
	return &model.TelemetryModel{ID: int64(m.ID)}, err
}

//ListDeviceModel ..
func (s DeviceService) ListDeviceModel(ctx context.Context, keyword *string,
	page, pageSize *int64, ids []int64, sorts []*model.Sort) (int64, []*model.DeviceModel, error) {
	result := make([]*model.DeviceModel, 0)
	data := make([]*models.DeviceModel, 0)
	offset, limit := utils.GetPageInfo(page, pageSize)
	fieldMap, _ := utils.GetFieldData(ctx, "")
	var err error
	builder := db.Gorm
	if keyword != nil && ptrs.String(keyword) != "" {
		builder = builder.Where("name like ?", "%"+ptrs.String(keyword)+"%")
	}
	var total int64
	if fieldMap["totalCount"] {
		if limit == -1 {
			total = int64(len(data))
		} else {
			err = builder.Model(&models.DeviceModel{}).Count(&total).Error
			if err != nil {
				return 0, result, err
			}
		}
	}
	if fieldMap["edges"] {
		edgeFieldMap, edgeFields := utils.GetFieldData(ctx, "edges.")
		builder = builder.Select(utils.ToDBFields(edgeFields, "__typename", "attributeModels", "telemetryModels"))
		if len(ids) > 0 {
			builder = builder.Where("id in (?)", ids)
		}
		if limit > 0 {
			builder = builder.Offset(offset).Limit(limit)
		}
		for _, v := range sorts {
			sort := " DESC"
			if v.IsAsc {
				sort = " ASC"
			}
			builder = builder.Order(v.Field + sort)
		}
		if edgeFieldMap["attributeModels"] {
			builder = builder.Preload("AttributeModels", func(db *gorm.DB) *gorm.DB {
				return db.Model(&models.AttributeModel{})
			})
		}
		if edgeFieldMap["telemetryModels"] {
			builder = builder.Preload("TelemetryModels", func(db *gorm.DB) *gorm.DB {
				return db.Model(&models.TelemetryModel{})
			})
		}
		err = builder.Find(&data).Error
		if err != nil {
			return 0, result, err
		}
	}
	for _, m := range data {
		r := toDeviceModelDto(m)
		result = append(result, r)
	}
	return total, result, nil
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
	if err := db.Gorm.Preload("AttributeModels", func(db *gorm.DB) *gorm.DB {
		return db.Model(&models.AttributeModel{})
	}).Preload("TelemetryModels", func(db *gorm.DB) *gorm.DB {
		return db.Model(&models.TelemetryModel{})
	}).First(deviceModel, "id=?", input.DeviceModelID).Error; err != nil {
		return nil, err
	}
	tx := db.Gorm.Begin()
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
	if err := db.Gorm.Select("id").First(m, "id=?", input.ID).Error; err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"name":     input.Name,
		"ip":       ptrs.String(input.IP),
		"port":     uint(ptrs.Int64(input.Port)),
		"username": ptrs.String(input.Username),
		"password": ptrs.String(input.Password),
	}
	err := db.Gorm.Model(m).Updates(updateMap).Error
	return &model.Device{ID: int64(m.ID)}, err
}

//ListDevice ..
func (s DeviceService) ListDevice(ctx context.Context, keyword *string,
	page, pageSize *int64, ids []int64, sorts []*model.Sort, deviceType *int64) (int64, []*model.Device, error) {
	result := make([]*model.Device, 0)
	data := make([]*models.Device, 0)
	offset, limit := utils.GetPageInfo(page, pageSize)
	fieldMap, _ := utils.GetFieldData(ctx, "")
	var err error
	builder := db.Gorm
	if keyword != nil && ptrs.String(keyword) != "" {
		builder = builder.Where("name like ?", "%"+ptrs.String(keyword)+"%")
	}
	if deviceType != nil {
		builder = builder.
			Where(db.TablePrefix+"_device_model.device_type = ?", ptrs.Int64(deviceType)).
			Joins("JOIN " + db.TablePrefix + "_device_model ON " + db.TablePrefix + "_device_model.id = " +
				db.TablePrefix + "_device.device_model_id")
	}
	var total int64
	if fieldMap["totalCount"] {
		if limit == -1 {
			total = int64(len(data))
		} else {
			err = builder.Model(&models.Device{}).Count(&total).Error
			if err != nil {
				return 0, result, err
			}
		}
	}
	if fieldMap["edges"] {
		edgeFieldMap, edgeFields := utils.GetFieldData(ctx, "edges.")
		fields := utils.ToDBFields(edgeFields,
			"__typename", "attributes", "telemetries",
			"deviceModelName", "deviceModelDesc",
			"deviceModelDeviceType", "deviceModelCameraCompany")
		if deviceType != nil {
			for i, v := range fields {
				fields[i] = db.TablePrefix + "_device." + v
			}
		}
		builder = builder.Select(fields)
		if len(ids) > 0 {
			builder = builder.Where("id in (?)", ids)
		}
		if limit > 0 {
			builder = builder.Offset(offset).Limit(limit)
		}
		for _, v := range sorts {
			sort := " DESC"
			if v.IsAsc {
				sort = " ASC"
			}
			builder = builder.Order(v.Field + sort)
		}
		if edgeFieldMap["attributes"] {
			builder = builder.Preload("Attributes").Preload("Attributes.AttributeModel")
		}
		if edgeFieldMap["telemetries"] {
			builder = builder.Preload("Telemetries").Preload("Telemetries.TelemetryModel")
		}
		if edgeFieldMap["deviceModelName"] || edgeFieldMap["deviceModelDesc"] ||
			edgeFieldMap["deviceModelDeviceType"] || edgeFieldMap["deviceModelCameraCompany"] {
			builder = builder.Preload("DeviceModel")
		}
		err = builder.Find(&data).Error
		if err != nil {
			return 0, result, err
		}
	}
	for _, m := range data {
		r := toDeviceDto(m)
		result = append(result, r)
	}
	return total, result, nil
}

//GetDeviceInfo ..
func (s DeviceService) GetDeviceInfo(deviceID uint32) (*pb.DeviceDownMSG, error) {
	replyDeviceMsg := &pb.DeviceDownMSG{
		DeviceId: deviceID,
	}
	device := new(models.Device)
	err := db.Gorm.Preload("Attributes").Preload("Attributes.AttributeModel").
		Preload("Telemetries").Preload("Telemetries.TelemetryModel").Preload("DeviceModel").
		Where("id=?", replyDeviceMsg.DeviceId).First(device).Error
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
	err := db.Gorm.Select("id,access_key,secret_key").
		Where("access_key=? and secret_key=?", accessKey, secretKey).First(device).Error
	if err != nil {
		log.Println("get device failed,err", err)
		return 0, err
	}
	return device.ID, err
}

//CreateDeviceDashboard  ..
func (s DeviceService) CreateDeviceDashboard(ctx context.Context, input model.NewDeviceDashboard) (*model.DeviceDashboard, error) {
	m := &models.DeviceDashboard{
		Name:       input.Name,
		IsVisible:  input.IsVisible,
		DeviceType: uint8(input.DeviceType),
	}
	err := db.Gorm.Create(m).Error
	if err != nil {
		return &model.DeviceDashboard{}, err
	}
	return &model.DeviceDashboard{ID: int64(m.ID)}, err
}

//UpdateDeviceDashboard ..
func (s DeviceService) UpdateDeviceDashboard(ctx context.Context, input model.NewUpdateDeviceDashboard) (*model.DeviceDashboard, error) {
	m := new(models.DeviceDashboard)
	varibales := graphql.GetRequestContext(ctx).Variables
	fields := make([]string, 0)
	for k := range varibales["input"].(map[string]interface{}) {
		fields = append(fields, k)
	}
	if err := db.Gorm.Select(utils.ToDBFields(fields)).First(m, "id=?", input.ID).Error; err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"name":       input.Name,
		"is_visible": input.IsVisible,
	}
	err := db.Gorm.Model(m).Updates(updateMap).Error
	return &model.DeviceDashboard{ID: int64(m.ID)}, err
}

//DeleteDeviceDashboard ..
func (s DeviceService) DeleteDeviceDashboard(ctx context.Context, id int64) (*model.DeviceDashboard, error) {
	m := new(models.DeviceDashboard)
	m.ID = uint(id)
	err := db.Gorm.Delete(m).Error
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
	err := db.Gorm.Create(data).Error
	if err != nil {
		return &model.DeviceDashboard{}, err
	}
	return &model.DeviceDashboard{ID: input.DeviceDashboardID}, err
}

//RemoveDeviceDashboardTelemetry ..
func (s DeviceService) RemoveDeviceDashboardTelemetry(ctx context.Context, ids []int64) (*model.DeviceDashboard, error) {
	m := new(models.DeviceDashboardTelemetry)
	err := db.Gorm.Delete(m, ids).Error
	return &model.DeviceDashboard{ID: int64(m.ID)}, err
}

//AddDeviceDashboardCamera  ..
func (s DeviceService) AddDeviceDashboardCamera(ctx context.Context, input model.NewDeviceDashboardCamera) (*model.DeviceDashboard, error) {
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
	err := db.Gorm.Create(data).Error
	if err != nil {
		return &model.DeviceDashboard{}, err
	}
	return &model.DeviceDashboard{ID: input.DeviceDashboardID}, err
}

//RemoveDeviceDashboardCamera ..
func (s DeviceService) RemoveDeviceDashboardCamera(ctx context.Context, ids []int64) (*model.DeviceDashboard, error) {
	m := new(models.DeviceDashboardCamera)
	err := db.Gorm.Delete(m, ids).Error
	return &model.DeviceDashboard{ID: int64(m.ID)}, err
}

//ListDeviceDashboards ..
func (s DeviceService) ListDeviceDashboards(ctx context.Context, keyword *string,
	page, pageSize *int64, ids []int64, sorts []*model.Sort) (int64, []*model.DeviceDashboard, error) {
	result := make([]*model.DeviceDashboard, 0)
	data := make([]*models.DeviceDashboard, 0)
	offset, limit := utils.GetPageInfo(page, pageSize)
	fieldMap, _ := utils.GetFieldData(ctx, "")
	var err error
	builder := db.Gorm
	if keyword != nil && ptrs.String(keyword) != "" {
		builder = builder.Where("name like ?", "%"+ptrs.String(keyword)+"%")
	}
	var total int64
	if fieldMap["totalCount"] {
		if limit == -1 {
			total = int64(len(data))
		} else {
			err = builder.Model(&models.DeviceDashboard{}).Count(&total).Error
			if err != nil {
				return 0, result, err
			}
		}
	}
	if fieldMap["edges"] {
		edgeFieldMap, edgeFields := utils.GetFieldData(ctx, "edges.")
		builder = builder.Select(utils.ToDBFields(edgeFields,
			"__typename", "telemetries", "cameras"))
		if len(ids) > 0 {
			builder = builder.Where("id in (?)", ids)
		}
		if limit > 0 {
			builder = builder.Offset(offset).Limit(limit)
		}
		for _, v := range sorts {
			sort := " DESC"
			if v.IsAsc {
				sort = " ASC"
			}
			builder = builder.Order(v.Field + sort)
		}
		if edgeFieldMap["telemetries"] {
			builder = builder.Preload("Telemetries").
				Preload("Telemetries.Telemetry").
				Preload("Telemetries.Telemetry.TelemetryModel")
			telemetryFieldMap, _ := utils.GetFieldData(ctx, "edges.telemetries.")
			if telemetryFieldMap["deviceName"] {
				builder = builder.Preload("Telemetries.Telemetry.Device")
			}
		}
		if edgeFieldMap["cameras"] {
			builder = builder.Preload("Cameras")
			cameraFieldMap, _ := utils.GetFieldData(ctx, "edges.cameras.")
			if cameraFieldMap["deviceName"] {
				builder = builder.Preload("Cameras.Device")
			}
		}
		err = builder.Find(&data).Error
		if err != nil {
			return 0, result, err
		}
	}
	for _, m := range data {
		r := toDeviceDashboardDto(m)
		result = append(result, r)
	}
	return total, result, nil
}

//AppDeviceDashboards ..
func (s DeviceService) AppDeviceDashboards(ctx context.Context, deviceType *int64) (int64, []*model.DeviceDashboard, error) {
	result := make([]*model.DeviceDashboard, 0)
	data := make([]*models.DeviceDashboard, 0)
	fieldMap, _ := utils.GetFieldData(ctx, "")
	var err error
	builder := db.Gorm
	if deviceType != nil {
		builder = builder.Where("device_type = ?", ptrs.Int64(deviceType))
	}
	builder = builder.Where("is_visible = true")
	var total int64
	if fieldMap["totalCount"] {
		total = int64(len(data))
	}
	if fieldMap["edges"] {
		edgeFieldMap, edgeFields := utils.GetFieldData(ctx, "edges.")
		builder = builder.Select(utils.ToDBFields(edgeFields,
			"__typename", "telemetries", "cameras"))
		if edgeFieldMap["telemetries"] {
			builder = builder.Preload("Telemetries").
				Preload("Telemetries.Telemetry").
				Preload("Telemetries.Telemetry.TelemetryModel")
			telemetryFieldMap, _ := utils.GetFieldData(ctx, "edges.telemetries.")
			if telemetryFieldMap["deviceName"] {
				builder = builder.Preload("Telemetries.Telemetry.Device")
			}
		}
		if edgeFieldMap["cameras"] {
			builder = builder.Preload("Cameras")
			cameraFieldMap, _ := utils.GetFieldData(ctx, "edges.cameras.")
			if cameraFieldMap["deviceName"] {
				builder = builder.Preload("Cameras.Device")
			}
		}
		err = builder.Find(&data).Error
		if err != nil {
			return 0, result, err
		}
	}
	for _, m := range data {
		r := toDeviceDashboardDto(m)
		result = append(result, r)
	}
	return total, result, nil
}

//CameraCapture ..
func (s DeviceService) CameraCapture(ctx context.Context, deviceID int64, scheme string) (string, error) {
	bucketName := "camera"
	now := time.Now()
	objectName := "picture/" + strconv.FormatInt(deviceID, 10) + "/" + now.Format("2006-01-02") + "/" + strconv.FormatInt(now.Unix(), 10) + ".jpg"
	minioClient := oss.MinioClient
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
			OssPrefix:       oss.OssPrefix,
			SecureOssPrefix: oss.SecureOssPrerix,
		},
	}
	requestMsg, err := proto.Marshal(request)
	if err != nil {
		log.Println("proto marshal error:", err)
		return "", nil
	}
	subject := mq.SubjectDevicPrefix + strconv.FormatUint(uint64(deviceID), 10)
	log.Println("发送数据到主题", subject)
	msg, err := mq.Client.NatsConn().Request(subject, requestMsg, 5*time.Second)
	if err != nil {
		log.Println("send data error:", err)
		return oss.GetOSSPrefix(scheme) + u.Path, nil
	}
	deviceMsg := new(pb.DeviceUpMsg)
	err = proto.Unmarshal(msg.Data, deviceMsg)
	if err != nil {
		log.Println("unmarshal data error")
	}
	return oss.GetOSSPrefix(scheme) + u.Path, nil
}
