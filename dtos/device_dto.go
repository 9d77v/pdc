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

//ToDeviceDto ...
func ToDeviceDto(m *models.Device) *model.Device {
	as := make([]*model.Attribute, 0, len(m.Attributes))
	for _, v := range m.Attributes {
		as = append(as, &model.Attribute{
			ID:        int64(v.ID),
			Key:       v.Key,
			Name:      v.AttributeModel.Name,
			Value:     v.Value,
			CreatedAt: v.CreatedAt.Unix(),
			UpdatedAt: v.UpdatedAt.Unix(),
		})
	}
	ts := make([]*model.Telemetry, 0, len(m.Telemetries))
	for _, v := range m.Telemetries {
		ts = append(ts, &model.Telemetry{
			ID:        int64(v.ID),
			Key:       v.Key,
			Name:      v.TelemetryModel.Name,
			Unit:      v.TelemetryModel.Unit,
			UnitName:  v.TelemetryModel.UnitName,
			Factor:    v.TelemetryModel.Factor,
			Scale:     int64(v.TelemetryModel.Scale),
			CreatedAt: v.CreatedAt.Unix(),
			UpdatedAt: v.UpdatedAt.Unix(),
		})
	}
	return &model.Device{
		ID:              int64(m.ID),
		Name:            m.Name,
		IP:              m.IP,
		Port:            int64(m.Port),
		DeviceModelID:   int64(m.DeviceModelID),
		DeviceModelName: m.DeviceModel.Name,
		DeviceModelDesc: m.DeviceModel.Desc,
		Attributes:      as,
		Telemetries:     ts,
		CreatedAt:       m.CreatedAt.Unix(),
		UpdatedAt:       m.UpdatedAt.Unix(),
	}
}
