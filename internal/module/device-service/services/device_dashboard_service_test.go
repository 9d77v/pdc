package services

import (
	"context"
	"testing"

	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/device-service/models"
	"github.com/9d77v/pdc/internal/module/device-service/pb"
	"github.com/stretchr/testify/assert"
)

var (
	deviceDashboardService = DeviceDashboardService{}
	testDeviceDashboard    = &pb.CreateDeviceDashboardRequest{
		Name:       "测试面板1",
		IsVisible:  true,
		DeviceType: pb.DeviceType_Camera,
	}
)

func TestDeviceService_CreateDeviceDashboard(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pb.CreateDeviceDashboardRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test CreateDeviceDashboard", args{ctx, testDeviceDashboard}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deviceDashboardService.CreateDeviceDashboard(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.NotZero(t, got.Id)
		})
	}
}

func TestDeviceService_UpdateDeviceDashboard(t *testing.T) {
	deviceDashboard, _ := deviceDashboardService.CreateDeviceDashboard(ctx, testDeviceDashboard)
	type args struct {
		ctx context.Context
		in  *pb.UpdateDeviceDashboardRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test deviceDashboard exist", args{ctx, &pb.UpdateDeviceDashboardRequest{
			Id:        deviceDashboard.Id,
			Name:      "测试面板2",
			IsVisible: false,
		}}, false},
		{"test deviceDashboard not exist", args{ctx, &pb.UpdateDeviceDashboardRequest{
			Id:        deviceDashboard.Id + 1,
			Name:      "测试面板2",
			IsVisible: false,
		}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deviceDashboardService.UpdateDeviceDashboard(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.NotZero(t, got.Id)
		})
	}
}

func TestDeviceService_DeleteDeviceDashboard(t *testing.T) {
	deviceDashboard, _ := deviceDashboardService.CreateDeviceDashboard(ctx, testDeviceDashboard)
	type args struct {
		ctx context.Context
		in  *pb.DeleteDeviceDashboardRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test ids length is zero", args{ctx, &pb.DeleteDeviceDashboardRequest{}}, false},
		{"test ids length is  not zero and record exist", args{ctx, &pb.DeleteDeviceDashboardRequest{
			Ids: []int64{int64(deviceDashboard.Id)},
		}}, false},
		{"test ids length is  not zero and record not exist", args{ctx, &pb.DeleteDeviceDashboardRequest{
			Ids: []int64{int64(deviceDashboard.Id + 1)},
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := deviceDashboardService.DeleteDeviceDashboard(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestDeviceDashboardService_AddTelemetries(t *testing.T) {
	deviceDashboard, _ := deviceDashboardService.CreateDeviceDashboard(ctx, testDeviceDashboard)
	wrongTelemetryRequest := &pb.AddTelemetriesRequest{
		DeviceDashboardId: deviceDashboard.Id,
		TelemetryIds:      []int64{1100, 2100, 3100},
	}
	deviceResp := createDevice()
	device := models.NewDevice()
	device.GetByID(uint(deviceResp.Id))
	ids := make([]int64, 0, len(device.Telemetries))
	for _, v := range device.Telemetries {
		ids = append(ids, int64(v.ID))
	}
	rightTelemetryRequest := &pb.AddTelemetriesRequest{
		DeviceDashboardId: deviceDashboard.Id,
		TelemetryIds:      ids,
	}
	type args struct {
		ctx context.Context
		in  *pb.AddTelemetriesRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.AddTelemetriesResponse
		wantErr bool
	}{
		{"test AddTelemetries wrong telemetryIDs", args{ctx, wrongTelemetryRequest},
			&pb.AddTelemetriesResponse{
				Ids: []int64{},
			}, false},
		{"test AddTelemetries right telemetryIDs", args{ctx, rightTelemetryRequest}, &pb.AddTelemetriesResponse{
			Ids: ids,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deviceDashboardService.AddTelemetries(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestDeviceDashboardService_RemoveTelemetries(t *testing.T) {
	deviceDashboard, _ := deviceDashboardService.CreateDeviceDashboard(ctx, testDeviceDashboard)
	dashboard := models.NewDeviceDashboard()
	dashboard.GetByID(uint(deviceDashboard.Id))
	removeIds := make([]int64, 0)
	for _, v := range dashboard.Telemetries {
		removeIds = append(removeIds, int64(v.ID))
	}
	type args struct {
		ctx context.Context
		in  *pb.RemoveTelemetriesRequest
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test ids length is zero", args{ctx, &pb.RemoveTelemetriesRequest{}}, false},
		{"test ids length is  not zero and record exist", args{ctx, &pb.RemoveTelemetriesRequest{
			Ids: removeIds,
		}}, false},
		{"test ids length is  not zero and record not exist", args{ctx, &pb.RemoveTelemetriesRequest{
			Ids: []int64{1000, 2000, 3000},
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := deviceDashboardService.RemoveTelemetries(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestDeviceDashboardService_AddCameras(t *testing.T) {
	deviceDashboard, _ := deviceDashboardService.CreateDeviceDashboard(ctx, testDeviceDashboard)
	wrongCameraRequest := &pb.AddCamerasRequest{
		DeviceDashboardId: deviceDashboard.Id,
		DeviceIds:         []int64{1100, 2100, 3100},
	}
	deviceResp := createDevice()
	rightCameraRequest := &pb.AddCamerasRequest{
		DeviceDashboardId: deviceDashboard.Id,
		DeviceIds:         []int64{deviceResp.Id},
	}
	type args struct {
		ctx context.Context
		in  *pb.AddCamerasRequest
	}
	tests := []struct {
		name     string
		args     args
		wantResp *pb.AddCamerasResponse
		wantErr  bool
	}{
		{"test AddCameras wrong deviceIDs", args{ctx, wrongCameraRequest},
			&pb.AddCamerasResponse{
				Ids: []int64{},
			}, false},
		{"test AddCameras right deviceIDs", args{ctx, rightCameraRequest},
			&pb.AddCamerasResponse{
				Ids: []int64{deviceResp.Id},
			}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deviceDashboardService.AddCameras(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantResp, got)
		})
	}
}

func TestDeviceDashboardService_RemoveCameras(t *testing.T) {
	deviceDashboard, _ := deviceDashboardService.CreateDeviceDashboard(ctx, testDeviceDashboard)
	deviceResp := createDevice()
	rightCameraRequest := &pb.AddCamerasRequest{
		DeviceDashboardId: deviceDashboard.Id,
		DeviceIds:         []int64{deviceResp.Id},
	}
	deviceDashboardService.AddCameras(ctx, rightCameraRequest)
	dashboard := models.NewDeviceDashboard()
	dashboard.GetByID(uint(deviceDashboard.Id))
	removeIds := make([]int64, 0)
	for _, v := range dashboard.Cameras {
		removeIds = append(removeIds, int64(v.ID))
	}
	type args struct {
		ctx context.Context
		in  *pb.RemoveCamerasRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test ids length is zero", args{ctx, &pb.RemoveCamerasRequest{}}, false},
		{"test ids length is  not zero and record exist", args{ctx, &pb.RemoveCamerasRequest{
			Ids: removeIds,
		}}, false},
		{"test ids length is  not zero and record not exist", args{ctx, &pb.RemoveCamerasRequest{
			Ids: []int64{1000, 2000, 3000},
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := deviceDashboardService.RemoveCameras(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestDeviceDashboardService_ListDeviceDashboards(t *testing.T) {
	deviceDashboard := createDeviceDashboard()
	type args struct {
		ctx context.Context
		in  *pb.ListDeviceDashboardRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test ListDeviceDashboards", args{context.Background(), &pb.ListDeviceDashboardRequest{
			SearchParam: &base.SearchParam{
				Ids: []int64{deviceDashboard.Id},
				QueryFields: []string{
					"edges",
					"edges.id",
					"edges.name",
					"edges.isVisible",
					"edges.deviceType",
					"edges.telemetries",
					"edges.telemetries.id",
					"edges.telemetries.deviceID",
					"edges.telemetries.deviceName",
					"edges.telemetries.telemetryID",
					"edges.telemetries.name",
					"edges.telemetries.value",
					"edges.telemetries.factor",
					"edges.telemetries.scale",
					"edges.telemetries.unit",
					"edges.cameras",
					"edges.cameras.id",
					"edges.cameras.deviceID",
					"edges.cameras.deviceName",
					"totalCount",
				},
			},
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deviceDashboardService.ListDeviceDashboards(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeviceDashboardService.ListDeviceDashboard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.EqualValues(t, 1, len(got.Edges))
			assert.LessOrEqual(t, int64(1), got.TotalCount)
		})
	}
}

func createDeviceDashboard() *pb.CreateDeviceDashboardResponse {
	deviceDashboard, _ := deviceDashboardService.CreateDeviceDashboard(ctx, testDeviceDashboard)
	deviceResp := createDevice()
	device := models.NewDevice()
	device.GetByID(uint(deviceResp.Id))
	ids := make([]int64, 0, len(device.Telemetries))
	for _, v := range device.Telemetries {
		ids = append(ids, int64(v.ID))
	}
	rightTelemetryRequest := &pb.AddTelemetriesRequest{
		DeviceDashboardId: deviceDashboard.Id,
		TelemetryIds:      ids,
	}
	deviceDashboardService.AddTelemetries(ctx, rightTelemetryRequest)
	rightCameraRequest := &pb.AddCamerasRequest{
		DeviceDashboardId: deviceDashboard.Id,
		DeviceIds:         []int64{deviceResp.Id},
	}
	deviceDashboardService.AddCameras(ctx, rightCameraRequest)
	return deviceDashboard
}

func TestDeviceDashboardService_ListAppDeviceDashboards(t *testing.T) {
	createDeviceDashboard()
	type args struct {
		ctx context.Context
		in  *pb.ListAppDeviceDashboardRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test ListAppDeviceDashboards", args{context.Background(), &pb.ListAppDeviceDashboardRequest{
			QueryFields: []string{
				"edges",
				"edges.id",
				"edges.name",
				"edges.isVisible",
				"edges.deviceType",
				"edges.telemetries",
				"edges.telemetries.id",
				"edges.telemetries.deviceID",
				"edges.telemetries.deviceName",
				"edges.telemetries.telemetryID",
				"edges.telemetries.name",
				"edges.telemetries.value",
				"edges.telemetries.factor",
				"edges.telemetries.scale",
				"edges.telemetries.unit",
				"edges.cameras",
				"edges.cameras.id",
				"edges.cameras.deviceID",
				"edges.cameras.deviceName",
				"totalCount",
			},
			DeviceType: ptrs.Int64Ptr(int64(pb.DeviceType_Camera)),
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deviceDashboardService.ListAppDeviceDashboards(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeviceDashboardService.ListDeviceDashboard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.LessOrEqual(t, int64(1), got.TotalCount)
		})
	}
}

func TestDeviceDashboardService_ListCameraTimeLapseVideos(t *testing.T) {
	m := models.NewCameraTimeLapseVideo()
	deviceResp := createDevice()
	m.Save(&models.CameraTimeLapseVideo{
		DeviceID: uint(deviceResp.Id),
		Date:     "2020-12-19",
		VideoURL: "http://oss.domain.local/1.jpg",
	})
	type args struct {
		ctx context.Context
		in  *pb.ListCameraTimeLapseVideoRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test ListCameraTimeLapseVideos", args{context.Background(), &pb.ListCameraTimeLapseVideoRequest{
			QueryFields: []string{
				"edges",
				"edges.id",
				"edges.deviceID",
				"edges.date",
				"edges.videoURL",
				"totalCount",
			},
			DeviceID: deviceResp.Id,
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deviceDashboardService.ListCameraTimeLapseVideos(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeviceDashboardService.ListCameraTimeLapseVideos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, 1, len(got.Edges))
			assert.Equal(t, int64(1), got.TotalCount)

		})
	}
}
