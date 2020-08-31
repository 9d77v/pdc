package dtos

import (
	"github.com/9d77v/pdc/graph/model"
	"github.com/9d77v/pdc/models"
)

//ToDeviceModelDto ...
func ToDeviceModelDto(m *models.DeviceModel) *model.DeviceModel {
	as := make([]*model.AttributeModel, 0, len(m.AttributeModels))
	for _, v := range m.AttributeModels {
		as = append(as, &model.AttributeModel{
			ID:            int64(v.ID),
			DeviceModelID: int64(v.DeviceModelID),
			Key:           v.Key,
			Name:          v.Name,
			CreatedAt:     v.CreatedAt.Unix(),
			UpdatedAt:     v.UpdatedAt.Unix(),
		})
	}
	ts := make([]*model.TelemetryModel, 0, len(m.TelemetryModels))
	for _, v := range m.TelemetryModels {
		ts = append(ts, &model.TelemetryModel{
			ID:            int64(v.ID),
			DeviceModelID: int64(v.DeviceModelID),
			Key:           v.Key,
			Name:          v.Name,
			Factor:        v.Factor,
			Unit:          v.Unit,
			UnitName:      v.UnitName,
			Scale:         int64(v.Scale),
			CreatedAt:     v.CreatedAt.Unix(),
			UpdatedAt:     v.UpdatedAt.Unix(),
		})
	}
	return &model.DeviceModel{
		ID:              int64(m.ID),
		Name:            m.Name,
		Desc:            m.Desc,
		DeviceType:      int64(m.DeviceType),
		AttributeModels: as,
		TelemetryModels: ts,
		CreatedAt:       m.CreatedAt.Unix(),
		UpdatedAt:       m.UpdatedAt.Unix(),
	}
}
