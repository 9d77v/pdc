package device_dto

import (
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/device-service/pb"
)

func GetDeviceModelConnection(data *pb.ListDeviceModelResponse) *model.DeviceModelConnection {
	return &model.DeviceModelConnection{
		TotalCount: data.TotalCount,
		Edges:      toDeviceModels(data.Edges),
	}
}

func toDeviceModels(data []*pb.DeviceModel) []*model.DeviceModel {
	result := make([]*model.DeviceModel, 0, len(data))
	for _, v := range data {
		r := toDeviceModel(v)
		result = append(result, r)
	}
	return result
}

func toDeviceModel(deviceModel *pb.DeviceModel) *model.DeviceModel {
	as := make([]*model.AttributeModel, 0, len(deviceModel.AttributeModels))
	for _, v := range deviceModel.AttributeModels {
		as = append(as, &model.AttributeModel{
			ID:            v.Id,
			DeviceModelID: v.DeviceModelId,
			Key:           v.Key,
			Name:          v.Name,
			CreatedAt:     v.CreatedAt.GetSeconds(),
			UpdatedAt:     v.UpdatedAt.GetSeconds(),
		})
	}
	ts := make([]*model.TelemetryModel, 0, len(deviceModel.TelemetryModels))
	for _, v := range deviceModel.TelemetryModels {
		ts = append(ts, &model.TelemetryModel{
			ID:            v.Id,
			DeviceModelID: v.DeviceModelId,
			Key:           v.Key,
			Name:          v.Name,
			Factor:        v.Factor,
			Unit:          v.Unit,
			UnitName:      v.UnitName,
			Scale:         int64(v.Scale),
			CreatedAt:     v.CreatedAt.GetSeconds(),
			UpdatedAt:     v.UpdatedAt.GetSeconds(),
		})
	}
	return &model.DeviceModel{
		ID:              deviceModel.Id,
		Name:            deviceModel.Name,
		DeviceType:      int64(deviceModel.DeviceType),
		Desc:            deviceModel.Desc,
		AttributeModels: as,
		TelemetryModels: ts,
		CameraCompany:   int64(deviceModel.CameraCompany),
		CreatedAt:       deviceModel.CreatedAt.GetSeconds(),
		UpdatedAt:       deviceModel.UpdatedAt.GetSeconds(),
	}
}

func GetDeviceConnection(data *pb.ListDeviceResponse) *model.DeviceConnection {
	return &model.DeviceConnection{
		TotalCount: data.TotalCount,
		Edges:      toDevices(data.Edges),
	}
}

func toDevices(data []*pb.Device) []*model.Device {
	result := make([]*model.Device, 0, len(data))
	for _, v := range data {
		r := toDevice(v)
		result = append(result, r)
	}
	return result
}

func toDevice(device *pb.Device) *model.Device {
	as := make([]*model.Attribute, 0, len(device.Attributes))
	for _, v := range device.Attributes {
		as = append(as, &model.Attribute{
			ID:        v.Id,
			Key:       v.Key,
			Name:      v.Name,
			Value:     v.Value,
			CreatedAt: v.CreatedAt.GetSeconds(),
			UpdatedAt: v.UpdatedAt.GetSeconds(),
		})
	}
	ts := make([]*model.Telemetry, 0, len(device.Telemetries))
	for _, v := range device.Telemetries {
		ts = append(ts, &model.Telemetry{
			ID:        v.Id,
			Key:       v.Key,
			Name:      v.Name,
			Unit:      v.Unit,
			UnitName:  v.UnitName,
			Factor:    v.Factor,
			Scale:     v.Scale,
			CreatedAt: v.CreatedAt.GetSeconds(),
			UpdatedAt: v.UpdatedAt.GetSeconds(),
		})
	}
	return &model.Device{
		ID:        device.Id,
		Name:      device.Name,
		IP:        device.Ip,
		Port:      device.Port,
		AccessKey: device.AccessKey,
		SecretKey: device.SecretKey,
		Username:  device.Username,
		Password:  device.Password,
		DeviceModel: &model.DeviceModel{
			ID:            device.DeviceModelId,
			Name:          device.DeviceModel.Name,
			Desc:          device.DeviceModel.Desc,
			DeviceType:    int64(device.DeviceModel.DeviceType),
			CameraCompany: int64(device.DeviceModel.CameraCompany),
		},
		Attributes:  as,
		Telemetries: ts,
		CreatedAt:   device.CreatedAt.GetSeconds(),
		UpdatedAt:   device.UpdatedAt.GetSeconds(),
	}
}

func GetDeviceDashboardConnection(data *pb.ListDeviceDashboardResponse) *model.DeviceDashboardConnection {
	return &model.DeviceDashboardConnection{
		TotalCount: data.TotalCount,
		Edges:      toDeviceDashboards(data.Edges),
	}
}

func toDeviceDashboards(data []*pb.DeviceDashboard) []*model.DeviceDashboard {
	result := make([]*model.DeviceDashboard, 0, len(data))
	for _, v := range data {
		r := toDeviceDashboard(v)
		result = append(result, r)
	}
	return result
}

func toDeviceDashboard(deviceDashboard *pb.DeviceDashboard) *model.DeviceDashboard {
	ts := make([]*model.DeviceDashboardTelemetry, 0, len(deviceDashboard.Telemetries))
	for _, v := range deviceDashboard.Telemetries {
		ts = append(ts, &model.DeviceDashboardTelemetry{
			ID:                v.Id,
			DeviceDashboardID: v.DeviceDashboardId,
			DeviceID:          v.DeviceId,
			DeviceName:        v.Name,
			TelemetryID:       v.TelemetryId,
			Key:               v.Key,
			Name:              v.Name,
			Unit:              v.Unit,
			UnitName:          v.UnitName,
			Factor:            v.Factor,
			Scale:             v.Scale,
			CreatedAt:         v.CreatedAt.GetSeconds(),
			UpdatedAt:         v.UpdatedAt.GetSeconds(),
		})
	}
	cs := make([]*model.DeviceDashboardCamera, 0, len(deviceDashboard.Cameras))
	for _, v := range deviceDashboard.Cameras {
		cs = append(cs, &model.DeviceDashboardCamera{
			ID:                v.Id,
			DeviceDashboardID: v.DeviceDashboardId,
			DeviceID:          v.DeviceId,
			DeviceName:        v.DeviceName,
			CreatedAt:         v.CreatedAt.GetSeconds(),
			UpdatedAt:         v.UpdatedAt.GetSeconds(),
		})
	}
	return &model.DeviceDashboard{
		ID:          deviceDashboard.Id,
		Name:        deviceDashboard.Name,
		IsVisible:   deviceDashboard.IsVisible,
		Telemetries: ts,
		Cameras:     cs,
		DeviceType:  int64(deviceDashboard.DeviceType),
		CreatedAt:   deviceDashboard.CreatedAt.GetSeconds(),
		UpdatedAt:   deviceDashboard.UpdatedAt.GetSeconds(),
	}
}

func GetAppDeviceDashboardConnection(data *pb.ListAppDeviceDashboardResponse) *model.DeviceDashboardConnection {
	return &model.DeviceDashboardConnection{
		TotalCount: data.TotalCount,
		Edges:      toDeviceDashboards(data.Edges),
	}
}

func GetCameraTimeLapseVideoConnection(data *pb.ListCameraTimeLapseVideoResponse) *model.CameraTimeLapseVideoConnection {
	return &model.CameraTimeLapseVideoConnection{
		TotalCount: data.TotalCount,
		Edges:      toCameraTimeLapseVideos(data.Edges),
	}
}

func toCameraTimeLapseVideos(data []*pb.CameraTimeLapseVideo) []*model.CameraTimeLapseVideo {
	result := make([]*model.CameraTimeLapseVideo, 0, len(data))
	for _, v := range data {
		r := toCameraTimeLapseVideo(v)
		result = append(result, r)
	}
	return result
}

func toCameraTimeLapseVideo(video *pb.CameraTimeLapseVideo) *model.CameraTimeLapseVideo {
	return &model.CameraTimeLapseVideo{
		ID:       video.Id,
		DeviceID: video.DeviceId,
		Date:     video.Date,
		VideoURL: video.VideoUrl,
	}
}
