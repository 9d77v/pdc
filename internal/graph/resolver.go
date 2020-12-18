package graph

import (
	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/device-service/pb"
	device "github.com/9d77v/pdc/internal/module/device-service/services"
	history "github.com/9d77v/pdc/internal/module/history-service/services"
	thing "github.com/9d77v/pdc/internal/module/thing-service/services"
	user "github.com/9d77v/pdc/internal/module/user-service/services"
	video "github.com/9d77v/pdc/internal/module/video-service/services"
)

//go:generate go run github.com/99designs/gqlgen
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//Resolver ..
type Resolver struct{}

var (
	videoService           = video.VideoService{}
	thingService           = thing.ThingService{}
	userService            = user.UserService{}
	historyService         = history.HistoryService{}
	deviceService          = device.DeviceService{}
	deviceModelService     = device.DeviceModelService{}
	deviceDashboardService = device.DeviceDashboardService{}
)

func getCreateDeviceModelRequest(input model.NewDeviceModel) *pb.CreateDeviceModelRequest {
	return &pb.CreateDeviceModelRequest{
		Name:          input.Name,
		Desc:          ptrs.String(input.Desc),
		DeviceType:    pb.DeviceType(input.DeviceType),
		CameraCompany: pb.CameraCompany(input.CameraCompany),
	}
}

func getUpdateDeviceModelRequest(input model.NewUpdateDeviceModel) *pb.UpdateDeviceModelRequest {
	return &pb.UpdateDeviceModelRequest{
		Id:   input.ID,
		Name: input.Name,
		Desc: ptrs.String(input.Desc),
	}
}

func getCreateAttributeModelRequest(input model.NewAttributeModel) *pb.CreateAttributeModelRequest {
	return &pb.CreateAttributeModelRequest{
		DeviceModelId: input.DeviceModelID,
		Key:           input.Key,
		Name:          input.Name,
	}
}

func getUpdateAttributeModelRequest(input model.NewUpdateAttributeModel) *pb.UpdateAttributeModelRequest {
	return &pb.UpdateAttributeModelRequest{
		Id:   input.ID,
		Name: input.Name,
	}
}

func getCreateTelemetryModelRequest(input model.NewTelemetryModel) *pb.CreateTelemetryModelRequest {
	return &pb.CreateTelemetryModelRequest{
		DeviceModelId: input.DeviceModelID,
		Key:           input.Key,
		Name:          input.Name,
		Factor:        input.Factor,
		Unit:          ptrs.String(input.Unit),
		UnitName:      ptrs.String(input.UnitName),
		Scale:         input.Scale,
	}
}

func getUpdateTelemetryModelRequest(input model.NewUpdateTelemetryModel) *pb.UpdateTelemetryModelRequest {
	return &pb.UpdateTelemetryModelRequest{
		Id:       input.ID,
		Name:     input.Name,
		Factor:   input.Factor,
		Unit:     input.Unit,
		UnitName: input.UnitName,
		Scale:    input.Scale,
	}
}

func getCreateDeviceRequest(input model.NewDevice) *pb.CreateDeviceRequest {
	return &pb.CreateDeviceRequest{
		DeviceModelId: input.DeviceModelID,
		Name:          input.Name,
		Ip:            ptrs.String(input.IP),
		Port:          uint64(ptrs.Int64(input.Port)),
		Username:      ptrs.String(input.Username),
		Password:      ptrs.String(input.Password),
	}
}

func getUpdateDeviceRequest(input model.NewUpdateDevice) *pb.UpdateDeviceRequest {
	return &pb.UpdateDeviceRequest{
		Id:       input.ID,
		Name:     input.Name,
		Ip:       ptrs.String(input.IP),
		Port:     uint64(ptrs.Int64(input.Port)),
		Username: ptrs.String(input.Username),
		Password: ptrs.String(input.Password),
	}
}

func getCreateDeviceDashboardRequest(input model.NewDeviceDashboard) *pb.CreateDeviceDashboardRequest {
	return &pb.CreateDeviceDashboardRequest{
		Name:       input.Name,
		DeviceType: pb.DeviceType(input.DeviceType),
		IsVisible:  input.IsVisible,
	}
}

func getUpdateDeviceDashboardRequest(input model.NewUpdateDeviceDashboard) *pb.UpdateDeviceDashboardRequest {
	return &pb.UpdateDeviceDashboardRequest{
		Id:        input.ID,
		Name:      input.Name,
		IsVisible: input.IsVisible,
	}
}

func getAddTelemetriesRequest(input model.NewDeviceDashboardTelemetry) *pb.AddTelemetriesRequest {
	return &pb.AddTelemetriesRequest{
		DeviceDashboardId: input.DeviceDashboardID,
		TelemetryIds:      input.TelemetryIDs,
	}
}

func getAddCamerasRequest(input model.NewDeviceDashboardCamera) *pb.AddCamerasRequest {
	return &pb.AddCamerasRequest{
		DeviceDashboardId: input.DeviceDashboardID,
		DeviceIds:         input.DeviceIDs,
	}
}
