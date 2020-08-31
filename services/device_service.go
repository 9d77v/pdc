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
	"github.com/jinzhu/gorm"
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
		"desc": input.Desc,
	}
	err := models.Gorm.Model(m).Update(updateMap).Error
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
	err := models.Gorm.Model(m).Update(updateMap).Error
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
	err := models.Gorm.Model(m).Update(updateMap).Error
	return &model.TelemetryModel{ID: int64(m.ID)}, err
}

//ListDeviceModel ..
func (s DeviceService) ListDeviceModel(ctx context.Context, keyword *string,
	page, pageSize *int64, ids []int64) (int64, []*model.DeviceModel, error) {
	result := make([]*model.DeviceModel, 0)
	data := make([]*models.DeviceModel, 0)
	offset, limit := GetPageInfo(page, pageSize)
	filedMap, _ := utils.GetFieldData(ctx, "")
	var err error
	builder := models.Gorm
	if keyword != nil && ptrs.String(keyword) != "" {
		builder = builder.Where("name like ?", "%"+ptrs.String(keyword)+"%")
	}
	if filedMap["edges"] {
		edgeFieldMap, edgeFields := utils.GetFieldData(ctx, "edges.")
		builder = builder.Select(utils.ToDBFields(edgeFields, "__typename", "attributeModels", "telemetryModels"))
		if len(ids) > 0 {
			builder = builder.Where("id in (?)", ids)
		}
		subBuilder := builder.Order("id desc")
		if limit > 0 {
			subBuilder = subBuilder.Offset(offset).Limit(limit)
		}
		if edgeFieldMap["attributeModels"] {
			subBuilder = subBuilder.Preload("AttributeModels", func(db *gorm.DB) *gorm.DB {
				return models.Gorm.Model(&models.AttributeModel{})
			})
		}
		if edgeFieldMap["telemetryModels"] {
			subBuilder = subBuilder.Preload("TelemetryModels", func(db *gorm.DB) *gorm.DB {
				return models.Gorm.Model(&models.TelemetryModel{})
			})
		}
		err = subBuilder.Find(&data).Error
		if err != nil {
			return 0, result, err
		}
	}
	var total int64
	if filedMap["totalCount"] {
		if limit == -1 {
			total = int64(len(data))
		} else {
			err = builder.Model(&models.DeviceModel{}).Count(&total).Error
			if err != nil {
				return 0, result, err
			}
		}
	}
	for _, m := range data {
		r := dtos.ToDeviceModelDto(m)
		result = append(result, r)
	}
	return total, result, nil
}
