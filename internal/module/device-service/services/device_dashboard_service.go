package services

import (
	"context"
	"time"

	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/device-service/models"
	"github.com/9d77v/pdc/internal/module/device-service/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//DeviceDashboardService ..
type DeviceDashboardService struct {
	base.Service
}

//CreateDeviceDashboard  ..
func (s DeviceDashboardService) CreateDeviceDashboard(ctx context.Context,
	in *pb.CreateDeviceDashboardRequest) (*pb.CreateDeviceDashboardResponse, error) {
	resp := new(pb.CreateDeviceDashboardResponse)
	m := models.NewDeviceDashboardFromPB(in)
	err := m.Create(m)
	resp.Id = int64(m.ID)
	return resp, err
}

//UpdateDeviceDashboard ..
func (s DeviceDashboardService) UpdateDeviceDashboard(ctx context.Context,
	in *pb.UpdateDeviceDashboardRequest) (*pb.UpdateDeviceDashboardResponse, error) {
	resp := &pb.UpdateDeviceDashboardResponse{
		Id: in.Id,
	}
	m := models.NewDeviceDashboard()
	if s.RecordNotExist(m, uint(in.Id)) {
		return resp, status.Error(codes.NotFound, "数据不存在")
	}
	return resp, m.Updates(map[string]interface{}{
		"name":       in.Name,
		"is_visible": in.IsVisible,
	})
}

//DeleteDeviceDashboard ..
func (s DeviceDashboardService) DeleteDeviceDashboard(ctx context.Context,
	in *pb.DeleteDeviceDashboardRequest) (*pb.DeleteDeviceDashboardResponse, error) {
	resp := &pb.DeleteDeviceDashboardResponse{
		Ids: in.Ids,
	}
	m := models.NewDeviceDashboard()
	return resp, m.Delete(m, in.Ids)
}

//AddTelemetries  ..
func (s DeviceDashboardService) AddTelemetries(ctx context.Context,
	in *pb.AddTelemetriesRequest) (resp *pb.AddTelemetriesResponse, err error) {
	resp = new(pb.AddTelemetriesResponse)
	m := models.NewTelemetry()
	ids := s.GetValidIDs(m, m.TableName(), s.ToUintIDs(in.TelemetryIds))
	if len(ids) > 0 {
		telemetries := make([]*models.DeviceDashboardTelemetry, 0, len(ids))
		for _, id := range ids {
			telemetries = append(telemetries, &models.DeviceDashboardTelemetry{
				DeviceDashboardID: uint(in.DeviceDashboardId),
				TelemetryID:       id,
			})
		}
		err = models.NewDeviceDashboardTelemetry().Create(&telemetries)
	}
	resp.Ids = s.ToInt64(ids)
	return resp, err
}

//RemoveTelemetries ..
func (s DeviceDashboardService) RemoveTelemetries(ctx context.Context,
	in *pb.RemoveTelemetriesRequest) (*pb.RemoveTelemetriesResponse, error) {
	resp := &pb.RemoveTelemetriesResponse{
		Ids: in.Ids,
	}
	m := models.NewDeviceDashboardTelemetry()
	return resp, m.Delete(m, in.Ids)
}

//AddCameras  ..
func (s DeviceDashboardService) AddCameras(ctx context.Context,
	in *pb.AddCamerasRequest) (resp *pb.AddCamerasResponse, err error) {
	resp = new(pb.AddCamerasResponse)
	m := models.NewDevice()
	m.SelectDeviceType(int64(pb.DeviceType_Camera))
	ids := s.GetValidIDs(m, m.TableName(), s.ToUintIDs(in.DeviceIds))
	if len(ids) > 0 {
		data := make([]*models.DeviceDashboardCamera, 0, len(ids))
		for _, v := range ids {
			data = append(data, &models.DeviceDashboardCamera{
				DeviceDashboardID: uint(in.DeviceDashboardId),
				DeviceID:          uint(v),
			})
		}
		err = models.NewDeviceDashboardCamera().Create(&data)
	}
	resp.Ids = s.ToInt64(ids)
	return resp, err
}

//RemoveCameras ..
func (s DeviceDashboardService) RemoveCameras(ctx context.Context,
	in *pb.RemoveCamerasRequest) (*pb.RemoveCamerasResponse, error) {
	resp := &pb.RemoveCamerasResponse{
		Ids: in.Ids,
	}
	m := models.NewDeviceDashboardCamera()
	return resp, m.Delete(m, in.Ids)
}

//ListDeviceDashboards ..
func (s DeviceDashboardService) ListDeviceDashboards(ctx context.Context,
	in *pb.ListDeviceDashboardRequest) (*pb.ListDeviceDashboardResponse, error) {
	resp := new(pb.ListDeviceDashboardResponse)
	m := models.NewDeviceDashboard()
	var deviceDashboard base.Repository = m
	deviceDashboard.FuzzyQuery(in.SearchParam.Keyword, "name")
	replaceFunc := func(edgeField base.GraphQLField) error {
		if edgeField.FieldMap["telemetries"] {
			deviceDashboard.Preload("Telemetries").
				Preload("Telemetries.Telemetry").
				Preload("Telemetries.Telemetry.TelemetryModel")
			telemetryField := base.ForGraphQLField(in.SearchParam.QueryFields, "edges.telemetries.")
			if telemetryField.FieldMap["deviceName"] {
				deviceDashboard.Preload("Telemetries.Telemetry.Device")
			}
		}
		if edgeField.FieldMap["cameras"] {
			deviceDashboard.Preload("Cameras")
			cameraField := base.ForGraphQLField(in.SearchParam.QueryFields, "edges.cameras.")
			if cameraField.FieldMap["deviceName"] {
				deviceDashboard.Preload("Cameras.Device")
			}
		}
		return nil
	}
	data := make([]*models.DeviceDashboard, 0)
	total, err := s.GetNewConnection(deviceDashboard, in.SearchParam, &data, replaceFunc, "telemetries", "cameras")
	resp.TotalCount = total
	resp.Edges = m.ToDeviceDashboardPBs(data)
	return resp, err
}

//ListAppDeviceDashboards ..
func (s DeviceDashboardService) ListAppDeviceDashboards(ctx context.Context,
	in *pb.ListAppDeviceDashboardRequest) (resp *pb.ListAppDeviceDashboardResponse, err error) {
	resp = new(pb.ListAppDeviceDashboardResponse)
	data := make([]*models.DeviceDashboard, 0)
	field := base.ForGraphQLField(in.QueryFields, "")
	m := models.NewDeviceDashboard()
	if in.DeviceType != nil {
		m.Where("device_type = ?", ptrs.Int64(in.DeviceType))
	}
	m.Where("is_visible = true")
	if field.FieldMap["edges"] {
		edgeField := base.ForGraphQLField(in.QueryFields, "edges.")
		if edgeField.FieldMap["telemetries"] {
			m.Preload("Telemetries").
				Preload("Telemetries.Telemetry").
				Preload("Telemetries.Telemetry.TelemetryModel")
			telemetryField := base.ForGraphQLField(in.QueryFields, "edges.telemetries.")
			if telemetryField.FieldMap["deviceName"] {
				m.Preload("Telemetries.Telemetry.Device")
			}
		}
		if edgeField.FieldMap["cameras"] {
			m.Preload("Cameras")
			cameraField := base.ForGraphQLField(in.QueryFields, "edges.cameras.")
			if cameraField.FieldMap["deviceName"] {
				m.Preload("Cameras.Device")
			}
		}
		err = m.Select(edgeField.Fields, "telemetries", "cameras").Find(&data)
	}
	resp.TotalCount = int64(len(data))
	resp.Edges = m.ToDeviceDashboardPBs(data)
	return resp, err
}

//ListCameraTimeLapseVideos ..
func (s DeviceDashboardService) ListCameraTimeLapseVideos(ctx context.Context,
	in *pb.ListCameraTimeLapseVideoRequest) (resp *pb.ListCameraTimeLapseVideoResponse, err error) {
	resp = new(pb.ListCameraTimeLapseVideoResponse)
	data := make([]*models.CameraTimeLapseVideo, 0)
	field := base.ForGraphQLField(in.QueryFields, "")
	m := models.NewCameraTimeLapseVideo()
	m.IDQuery(uint(in.DeviceID), "device_id").
		Where("created_at>=?", time.Now().AddDate(0, 0, -7))
	if field.FieldMap["edges"] {
		edgeField := base.ForGraphQLField(in.QueryFields, "edges.")
		err = m.Select(edgeField.Fields).
			Order("id DESC").
			Find(&data)
	}
	resp.Edges = m.ToCameraTimeLapseVideoPBs(data, in.Scheme)
	resp.TotalCount = int64(len(resp.Edges))
	return resp, nil
}
