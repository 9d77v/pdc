package services

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/dtos"
	"github.com/9d77v/pdc/graph/model"
	"github.com/9d77v/pdc/models"
	"github.com/9d77v/pdc/utils"
	"gorm.io/gorm"
)

//DeviceService ..
type DeviceService struct {
}

//CreateDeviceModel  ..
func (s DeviceService) CreateDeviceModel(ctx context.Context, input model.NewDeviceModel) (*model.DeviceModel, error) {
	m := &models.DeviceModel{
		Name:       input.Name,
		Desc:       ptrs.String(input.Desc),
		DeviceType: uint8(input.DeviceType),
	}
	err := models.Gorm.Create(m).Error
	if err != nil {
		return &model.DeviceModel{}, err
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
	if err := models.Gorm.Select(utils.ToDBFields(fields)).First(m, "id=?", input.ID).Error; err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"name": input.Name,
		"desc": ptrs.String(input.Desc),
	}
	err := models.Gorm.Model(m).Updates(updateMap).Error
	return &model.DeviceModel{ID: int64(m.ID)}, err
}

//CreateAttributeModel  ..
func (s DeviceService) CreateAttributeModel(ctx context.Context, input model.NewAttributeModel) (*model.AttributeModel, error) {
	m := &models.AttributeModel{
		DeviceModelID: uint(input.DeviceModelID),
		Key:           input.Key,
		Name:          input.Name,
	}
	err := models.Gorm.Create(m).Error
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
	if err := models.Gorm.Select(utils.ToDBFields(fields)).First(m, "id=?", input.ID).Error; err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"name": input.Name,
	}
	fmt.Println(updateMap)
	err := models.Gorm.Model(m).Updates(updateMap).Error
	return &model.AttributeModel{ID: int64(m.ID)}, err
}

//DeleteAttributeModel ..
func (s DeviceService) DeleteAttributeModel(ctx context.Context, id int64) (*model.AttributeModel, error) {
	m := new(models.AttributeModel)
	m.ID = uint(id)
	err := models.Gorm.Delete(m).Error
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
	err := models.Gorm.Create(m).Error
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
	if err := models.Gorm.Select(utils.ToDBFields(fields)).First(m, "id=?", input.ID).Error; err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"name":      input.Name,
		"factor":    input.Factor,
		"unit":      input.Unit,
		"unit_name": input.UnitName,
		"scale":     input.Scale,
	}
	err := models.Gorm.Model(m).Updates(updateMap).Error
	return &model.TelemetryModel{ID: int64(m.ID)}, err
}

//DeleteTelemetryModel ..
func (s DeviceService) DeleteTelemetryModel(ctx context.Context, id int64) (*model.TelemetryModel, error) {
	m := new(models.TelemetryModel)
	m.ID = uint(id)
	err := models.Gorm.Delete(m).Error
	return &model.TelemetryModel{ID: int64(m.ID)}, err
}

//ListDeviceModel ..
func (s DeviceService) ListDeviceModel(ctx context.Context, keyword *string,
	page, pageSize *int64, ids []int64, sorts []*model.Sort) (int64, []*model.DeviceModel, error) {
	result := make([]*model.DeviceModel, 0)
	data := make([]*models.DeviceModel, 0)
	offset, limit := GetPageInfo(page, pageSize)
	fieldMap, _ := utils.GetFieldData(ctx, "")
	var err error
	builder := models.Gorm
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
			if v.IsAsc {
				builder = builder.Order(v.Field + " ASC")
			} else {
				builder = builder.Order(v.Field + " DESC")
			}
		}
		if edgeFieldMap["attributeModels"] {
			builder = builder.Preload("AttributeModels", func(db *gorm.DB) *gorm.DB {
				return models.Gorm.Model(&models.AttributeModel{})
			})
		}
		if edgeFieldMap["telemetryModels"] {
			builder = builder.Preload("TelemetryModels", func(db *gorm.DB) *gorm.DB {
				return models.Gorm.Model(&models.TelemetryModel{})
			})
		}
		err = builder.Find(&data).Error
		if err != nil {
			return 0, result, err
		}
	}
	for _, m := range data {
		r := dtos.ToDeviceModelDto(m)
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
	}
	deviceModel := new(models.DeviceModel)
	if err := models.Gorm.Preload("AttributeModels", func(db *gorm.DB) *gorm.DB {
		return models.Gorm.Model(&models.AttributeModel{})
	}).Preload("TelemetryModels", func(db *gorm.DB) *gorm.DB {
		return models.Gorm.Model(&models.TelemetryModel{})
	}).First(deviceModel, "id=?", input.DeviceModelID).Error; err != nil {
		return nil, err
	}
	tx := models.Gorm.Begin()
	err := tx.Create(m).Error
	if err != nil {
		tx.Rollback()
		return &model.Device{}, err
	}
	attributes := make([]*models.Attribute, 0, len(deviceModel.AttributeModels))
	for _, v := range deviceModel.AttributeModels {
		attributes = append(attributes, &models.Attribute{
			DeviceID:         m.ID,
			AttributeModelID: v.ID,
			Key:              v.Key,
		})
	}
	err = tx.Create(&attributes).Error
	if err != nil {
		tx.Rollback()
		return &model.Device{}, err
	}
	telemetries := make([]*models.Telemetry, 0, len(deviceModel.TelemetryModels))
	for _, v := range deviceModel.TelemetryModels {
		telemetries = append(telemetries, &models.Telemetry{
			DeviceID:         m.ID,
			Key:              v.Key,
			TelemetryModelID: v.ID,
		})

	}
	err = tx.Create(&telemetries).Error
	if err != nil {
		tx.Rollback()
		return &model.Device{}, err
	}
	tx.Commit()
	return &model.Device{ID: int64(m.ID)}, err
}

//UpdateDevice ..
func (s DeviceService) UpdateDevice(ctx context.Context, input model.NewUpdateDevice) (*model.Device, error) {
	m := new(models.Device)
	if err := models.Gorm.Select("id").First(m, "id=?", input.ID).Error; err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"name": input.Name,
		"ip":   ptrs.String(input.IP),
		"port": uint(ptrs.Int64(input.Port)),
	}
	err := models.Gorm.Model(m).Updates(updateMap).Error
	return &model.Device{ID: int64(m.ID)}, err
}

//ListDevice ..
func (s DeviceService) ListDevice(ctx context.Context, keyword *string,
	page, pageSize *int64, ids []int64, sorts []*model.Sort) (int64, []*model.Device, error) {
	result := make([]*model.Device, 0)
	data := make([]*models.Device, 0)
	offset, limit := GetPageInfo(page, pageSize)
	fieldMap, _ := utils.GetFieldData(ctx, "")
	var err error
	builder := models.Gorm
	if keyword != nil && ptrs.String(keyword) != "" {
		builder = builder.Where("name like ?", "%"+ptrs.String(keyword)+"%")
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
		builder = builder.Select(utils.ToDBFields(edgeFields,
			"__typename", "attributes", "telemetries", "deviceModelName", "deviceModelDesc"))
		if len(ids) > 0 {
			builder = builder.Where("id in (?)", ids)
		}
		if limit > 0 {
			builder = builder.Offset(offset).Limit(limit)
		}
		for _, v := range sorts {
			if v.IsAsc {
				builder = builder.Order(v.Field + " ASC")
			} else {
				builder = builder.Order(v.Field + " DESC")
			}
		}
		if edgeFieldMap["attributes"] {
			builder = builder.Preload("Attributes").Preload("Attributes.AttributeModel")
		}
		if edgeFieldMap["telemetries"] {
			builder = builder.Preload("Telemetries").Preload("Telemetries.TelemetryModel")
		}
		if edgeFieldMap["deviceModelName"] || edgeFieldMap["deviceModelDesc"] {
			builder = builder.Preload("DeviceModel")
		}
		err = builder.Find(&data).Error
		if err != nil {
			return 0, result, err
		}
	}
	for _, m := range data {
		r := dtos.ToDeviceDto(m)
		result = append(result, r)
	}
	return total, result, nil
}
