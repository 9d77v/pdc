package services

import (
	"context"
	"time"

	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/graph/model"
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
	telemetry := models.NewTelemetry()
	ids := s.GetValidIDs(telemetry, telemetry.TableName(), s.ToUintIDs(in.TelemetryIds))
	if len(ids) > 0 {
		telemetries := make([]*models.DeviceDashboardTelemetry, 0, len(ids))
		for _, id := range ids {
			telemetries = append(telemetries, &models.DeviceDashboardTelemetry{
				DeviceDashboardID: uint(in.DeviceDashboardId),
				TelemetryID:       id,
			})
		}
		err = models.NewDeviceDashboard().Create(&telemetries)
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
	device := models.NewDevice()
	device.SelectDeviceType(int64(pb.DeviceType_Camera))
	ids := s.GetValidIDs(device, device.TableName(), s.ToUintIDs(in.DeviceIds))
	if len(ids) > 0 {
		data := make([]*models.DeviceDashboardCamera, 0, len(ids))
		for _, v := range ids {
			m := &models.DeviceDashboardCamera{
				DeviceDashboardID: uint(in.DeviceDashboardId),
				DeviceID:          uint(v),
			}
			data = append(data, m)
		}
		err = db.GetDB().Create(data).Error
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
func (s DeviceDashboardService) ListDeviceDashboards(ctx context.Context, searchParam model.SearchParam) (int64, []*model.DeviceDashboard, error) {
	var deviceDashboard base.Repository = models.NewDeviceDashboard()
	deviceDashboard.FuzzyQuery(searchParam.Keyword, "name")
	replaceFunc := func(edgeField base.GraphQLField) error {
		if edgeField.FieldMap["telemetries"] {
			deviceDashboard.Preload("Telemetries").
				Preload("Telemetries.Telemetry").
				Preload("Telemetries.Telemetry.TelemetryModel")
			telemetryField := base.NewGraphQLField(ctx, "edges.telemetries.")
			if telemetryField.FieldMap["deviceName"] {
				deviceDashboard.Preload("Telemetries.Telemetry.Device")
			}
		}
		if edgeField.FieldMap["cameras"] {
			deviceDashboard.Preload("Cameras")
			cameraField := base.NewGraphQLField(ctx, "edges.cameras.")
			if cameraField.FieldMap["deviceName"] {
				deviceDashboard.Preload("Cameras.Device")
			}
		}
		return nil
	}
	data := make([]*models.DeviceDashboard, 0)
	total, err := s.GetConnection(ctx, deviceDashboard, searchParam, &data, replaceFunc,
		"telemetries", "cameras")
	return total, getDeviceDashboards(data), err
}

//AppDeviceDashboards ..
func (s DeviceDashboardService) AppDeviceDashboards(ctx context.Context,
	deviceType *int64) (int64, []*model.DeviceDashboard, error) {
	result := make([]*model.DeviceDashboard, 0)
	data := make([]*models.DeviceDashboard, 0)
	field := base.NewGraphQLField(ctx, "")
	deviceDashboard := models.NewDeviceDashboard()
	if deviceType != nil {
		deviceDashboard.Where("device_type = ?", ptrs.Int64(deviceType))
	}
	deviceDashboard.Where("is_visible = true")
	var total int64
	if field.FieldMap["totalCount"] {
		total = int64(len(data))
	}
	if field.FieldMap["edges"] {
		edgeField := base.NewGraphQLField(ctx, "edges.")
		if edgeField.FieldMap["telemetries"] {
			deviceDashboard.Preload("Telemetries").
				Preload("Telemetries.Telemetry").
				Preload("Telemetries.Telemetry.TelemetryModel")
			telemetryField := base.NewGraphQLField(ctx, "edges.telemetries.")
			if telemetryField.FieldMap["deviceName"] {
				deviceDashboard.Preload("Telemetries.Telemetry.Device")
			}
		}
		if edgeField.FieldMap["cameras"] {
			deviceDashboard.Preload("Cameras")
			cameraField := base.NewGraphQLField(ctx, "edges.cameras.")
			if cameraField.FieldMap["deviceName"] {
				deviceDashboard.Preload("Cameras.Device")
			}
		}
		err := deviceDashboard.Select(edgeField.Fields, "telemetries", "cameras").Find(&data)
		if err != nil {
			return 0, result, err
		}
	}
	return total, getDeviceDashboards(data), nil
}

//CameraTimeLapseVideos ..
func (s DeviceDashboardService) CameraTimeLapseVideos(ctx context.Context,
	deviceID int64, scheme string) (*model.CameraTimeLapseVideoConnection, error) {
	result := new(model.CameraTimeLapseVideoConnection)
	data := make([]*models.CameraTimeLapseVideo, 0)
	field := base.NewGraphQLField(ctx, "")
	camera := models.NewCameraTimeLapseVideo()
	camera.IDQuery(uint(deviceID), "device_id").
		Where("created_at>=?", time.Now().AddDate(0, 0, -7))
	var total int64
	if field.FieldMap["totalCount"] {
		total = int64(len(data))
	}
	if field.FieldMap["edges"] {
		edgeField := base.NewGraphQLField(ctx, "edges.")
		err := camera.Select(edgeField.Fields).
			Order("id DESC").
			Find(&data)
		if err != nil {
			return result, err
		}
	}
	result.Edges = getCameraTimeLapseVideos(data, scheme)
	result.TotalCount = total
	return result, nil
}
