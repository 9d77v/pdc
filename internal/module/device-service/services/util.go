package services

import (
	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/device-service/models"
)

func toDeviceModelDtos(data []*models.DeviceModel) []*model.DeviceModel {
	result := make([]*model.DeviceModel, 0, len(data))
	for _, m := range data {
		r := toDeviceModelDto(m)
		result = append(result, r)
	}
	return result
}

func toDeviceModelDto(m *models.DeviceModel) *model.DeviceModel {
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

func toDeviceDtos(data []*models.Device) []*model.Device {
	result := make([]*model.Device, 0, len(data))
	for _, m := range data {
		r := toDeviceDto(m)
		result = append(result, r)
	}
	return result
}

func toDeviceDto(m *models.Device) *model.Device {
	as := make([]*model.Attribute, 0, len(m.Attributes))
	for _, v := range m.Attributes {
		as = append(as, &model.Attribute{
			ID:        int64(v.ID),
			Key:       v.AttributeModel.Key,
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
			Key:       v.TelemetryModel.Key,
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
		ID:                       int64(m.ID),
		Name:                     m.Name,
		IP:                       m.IP,
		Port:                     int64(m.Port),
		AccessKey:                m.AccessKey,
		SecretKey:                m.SecretKey,
		Username:                 m.Username,
		Password:                 m.Password,
		DeviceModelID:            int64(m.DeviceModelID),
		DeviceModelName:          m.DeviceModel.Name,
		DeviceModelDesc:          m.DeviceModel.Desc,
		DeviceModelDeviceType:    int64(m.DeviceModel.DeviceType),
		DeviceModelCameraCompany: int64(m.DeviceModel.CameraCompany),
		Attributes:               as,
		Telemetries:              ts,
		CreatedAt:                m.CreatedAt.Unix(),
		UpdatedAt:                m.UpdatedAt.Unix(),
	}
}

func toDeviceDashboardDtos(data []*models.DeviceDashboard) []*model.DeviceDashboard {
	result := make([]*model.DeviceDashboard, 0, len(data))
	for _, m := range data {
		r := toDeviceDashboardDto(m)
		result = append(result, r)
	}
	return result
}
func toDeviceDashboardDto(m *models.DeviceDashboard) *model.DeviceDashboard {
	ts := make([]*model.DeviceDashboardTelemetry, 0, len(m.Telemetries))
	for _, v := range m.Telemetries {
		ts = append(ts, &model.DeviceDashboardTelemetry{
			ID:                int64(v.ID),
			DeviceDashboardID: int64(v.DeviceDashboardID),
			DeviceID:          int64(v.Telemetry.DeviceID),
			DeviceName:        v.Telemetry.Device.Name,
			TelemetryID:       int64(v.TelemetryID),
			Key:               v.Telemetry.TelemetryModel.Key,
			Name:              v.Telemetry.TelemetryModel.Name,
			Unit:              v.Telemetry.TelemetryModel.Unit,
			UnitName:          v.Telemetry.TelemetryModel.UnitName,
			Factor:            v.Telemetry.TelemetryModel.Factor,
			Scale:             int64(v.Telemetry.TelemetryModel.Scale),
			CreatedAt:         v.CreatedAt.Unix(),
			UpdatedAt:         v.UpdatedAt.Unix(),
		})
	}
	cs := make([]*model.DeviceDashboardCamera, 0, len(m.Cameras))
	for _, v := range m.Cameras {
		cs = append(cs, &model.DeviceDashboardCamera{
			ID:                int64(v.ID),
			DeviceDashboardID: int64(v.DeviceDashboardID),
			DeviceID:          int64(v.DeviceID),
			DeviceName:        v.Device.Name,
			CreatedAt:         v.CreatedAt.Unix(),
			UpdatedAt:         v.UpdatedAt.Unix(),
		})
	}
	return &model.DeviceDashboard{
		ID:          int64(m.ID),
		Name:        m.Name,
		IsVisible:   m.IsVisible,
		Telemetries: ts,
		Cameras:     cs,
		DeviceType:  int64(m.DeviceType),
		CreatedAt:   m.CreatedAt.Unix(),
		UpdatedAt:   m.UpdatedAt.Unix(),
	}
}

func toCameraTimeLapseVideoDtos(data []*models.CameraTimeLapseVideo, scheme string) []*model.CameraTimeLapseVideo {
	result := make([]*model.CameraTimeLapseVideo, 0, len(data))
	for _, m := range data {
		r := toCameraTimeLapseVideoDto(m, scheme)
		result = append(result, r)
	}
	return result
}

func toCameraTimeLapseVideoDto(m *models.CameraTimeLapseVideo, scheme string) *model.CameraTimeLapseVideo {
	return &model.CameraTimeLapseVideo{
		ID:       int64(m.ID),
		DeviceID: int64(m.DeviceID),
		Date:     m.Date,
		VideoURL: oss.GetOSSPrefix(scheme) + m.VideoURL,
	}
}

func getDefaultHikvisionAttributeModels(id uint) []*models.AttributeModel {
	return []*models.AttributeModel{
		{
			DeviceModelID: id,
			Key:           "device_name",
			Name:          "设备名称",
		},
		{
			DeviceModelID: id,
			Key:           "device_id",
			Name:          "设备ID",
		},
		{
			DeviceModelID: id,
			Key:           "device_description",
			Name:          "设备描述",
		},
		{
			DeviceModelID: id,
			Key:           "device_location",
			Name:          "设备位置",
		},
		{
			DeviceModelID: id,
			Key:           "system_contact",
			Name:          "系统联系方",
		},
		{
			DeviceModelID: id,
			Key:           "model",
			Name:          "类型",
		},
		{
			DeviceModelID: id,
			Key:           "serial_number",
			Name:          "序列号",
		},
		{
			DeviceModelID: id,
			Key:           "mac_address",
			Name:          "MAC地址",
		},
		{
			DeviceModelID: id,
			Key:           "firmware_version",
			Name:          "固件版本",
		},
		{
			DeviceModelID: id,
			Key:           "firmware_releasedDate",
			Name:          "固件发布日期",
		},
		{
			DeviceModelID: id,
			Key:           "encoder_version",
			Name:          "编码器版本",
		},
		{
			DeviceModelID: id,
			Key:           "encoder_released_date",
			Name:          "编码器发布日期",
		},
		{
			DeviceModelID: id,
			Key:           "boot_version",
			Name:          "引导版本",
		},
		{
			DeviceModelID: id,
			Key:           "boot_released_date",
			Name:          "引导发布日期",
		},
		{
			DeviceModelID: id,
			Key:           "hardware_version",
			Name:          "硬件版本",
		},
		{
			DeviceModelID: id,
			Key:           "device_type",
			Name:          "设备类型",
		},
		{
			DeviceModelID: id,
			Key:           "telecontrol_id",
			Name:          "远程 ID",
		},
		{
			DeviceModelID: id,
			Key:           "support beep",
			Name:          "支持蜂鸣音",
		},
		{
			DeviceModelID: id,
			Key:           "support_video_loss",
			Name:          "支持视频丢失",
		},
		{
			DeviceModelID: id,
			Key:           "firmware_version_info",
			Name:          "固件版本信息",
		},
	}
}
