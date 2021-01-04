package graph

import (
	"context"
	"time"

	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/device-service/pb"
	device "github.com/9d77v/pdc/internal/module/device-service/services"
	history "github.com/9d77v/pdc/internal/module/history-service/services"
	notePB "github.com/9d77v/pdc/internal/module/note-service/pb"
	note "github.com/9d77v/pdc/internal/module/note-service/services"
	thing "github.com/9d77v/pdc/internal/module/thing-service/services"
	user "github.com/9d77v/pdc/internal/module/user-service/services"
	"github.com/golang/protobuf/ptypes"

	video "github.com/9d77v/pdc/internal/module/video-service/services"
	"github.com/9d77v/pdc/internal/utils"
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
	noteService            = note.NoteService{}
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

func getSearchParam(ctx context.Context, input model.SearchParam) *base.SearchParam {
	sorts := make([]*base.Sort, 0, len(input.Sorts))
	for _, v := range input.Sorts {
		sorts = append(sorts, &base.Sort{
			Field: v.Field,
			IsAsc: v.IsAsc,
		})
	}
	return &base.SearchParam{
		QueryFields: utils.GetPreloads(ctx),
		Keyword:     ptrs.String(input.Keyword),
		Page:        ptrs.Int64(input.Page),
		PageSize:    ptrs.Int64(input.PageSize),
		Ids:         input.Ids,
		Tags:        input.Tags,
		Sorts:       sorts,
		IsRandom:    ptrs.Bool(input.IsRandom),
		IsMobile:    ptrs.Bool(input.IsMobile),
	}
}

func getDeviceModelConnection(data *pb.ListDeviceModelResponse) *model.DeviceModelConnection {
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

func getDeviceConnection(data *pb.ListDeviceResponse) *model.DeviceConnection {
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

func getDeviceDashboardConnection(data *pb.ListDeviceDashboardResponse) *model.DeviceDashboardConnection {
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

func getAppDeviceDashboardConnection(data *pb.ListAppDeviceDashboardResponse) *model.DeviceDashboardConnection {
	return &model.DeviceDashboardConnection{
		TotalCount: data.TotalCount,
		Edges:      toDeviceDashboards(data.Edges),
	}
}

func getCameraTimeLapseVideoConnection(data *pb.ListCameraTimeLapseVideoResponse) *model.CameraTimeLapseVideoConnection {
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

func getNotes(data []*model.NewNote) []*notePB.Note {
	result := make([]*notePB.Note, 0, len(data))
	for _, v := range data {
		r := getNote(v)
		result = append(result, r)
	}
	return result
}
func getNote(m *model.NewNote) *notePB.Note {
	createdAt, _ := ptypes.TimestampProto(time.Unix(m.CreatedAt, 0))
	updatedAt, _ := ptypes.TimestampProto(time.Unix(m.UpdatedAt, 0))
	return &notePB.Note{
		Id:        m.ID,
		ParentId:  m.ParentID,
		Uid:       m.UID,
		NoteType:  notePB.NoteType(m.NoteType),
		Level:     int32(m.Level),
		Title:     m.Title,
		Color:     m.Color,
		State:     notePB.NoteState(m.State),
		Version:   int32(m.Version),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Content:   m.Content,
		Tags:      m.Tags,
		Sha1:      m.Sha1,
	}
}

func toNotes(data []*notePB.Note) []*model.Note {
	result := make([]*model.Note, 0, len(data))
	for _, v := range data {
		r := toNote(v)
		result = append(result, r)
	}
	return result
}

func toNote(m *notePB.Note) *model.Note {
	return &model.Note{
		ID:        m.Id,
		ParentID:  m.ParentId,
		UID:       m.Uid,
		NoteType:  int64(m.NoteType),
		Level:     int64(m.Level),
		Title:     m.Title,
		Color:     m.Color,
		State:     int64(m.State),
		Version:   int64(m.Version),
		CreatedAt: m.CreatedAt.Seconds,
		UpdatedAt: m.UpdatedAt.Seconds,
		Content:   m.Content,
		Tags:      m.Tags,
		Sha1:      m.Sha1,
	}
}
