package services

import (
	"context"
	"testing"

	"github.com/9d77v/go-pkg/ptrs"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/device-service/pb"
	"github.com/stretchr/testify/assert"
)

var (
	deviceService = DeviceService{}
)

func TestDeviceService_CreateDevice(t *testing.T) {
	deviceModel, _ := deviceModelService.CreateDeviceModel(ctx, testDeviceModel)
	testTelemetryModel := &pb.CreateTelemetryModelRequest{
		DeviceModelId: deviceModel.Id,
		Key:           "key",
		Name:          "name",
		Factor:        1,
		Unit:          "A",
		UnitName:      "安培",
		Scale:         2,
	}
	deviceModelService.CreateTelemetryModel(ctx, testTelemetryModel)
	device := &pb.CreateDeviceRequest{
		DeviceModelId: deviceModel.Id,
		Name:          "测试设备",
		Ip:            "192.168.1.111",
		Port:          3030,
		Username:      "user",
		Password:      "password",
	}
	type args struct {
		ctx context.Context
		in  *pb.CreateDeviceRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test CreateDevice", args{ctx, device}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deviceService.CreateDevice(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.NotZero(t, got.Id)
		})
	}
}

func TestDeviceService_UpdateDevice(t *testing.T) {
	device := createDevice()
	type args struct {
		ctx context.Context
		in  *pb.UpdateDeviceRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test device exist", args{ctx, &pb.UpdateDeviceRequest{
			Id:       device.Id,
			Name:     "测试设备2",
			Ip:       "192.168.1.222",
			Port:     3050,
			Username: "user2",
			Password: "password2",
		}}, false},
		{"test device not exist", args{ctx, &pb.UpdateDeviceRequest{
			Id:       device.Id + 1,
			Name:     "测试设备2",
			Ip:       "192.168.1.222",
			Port:     3050,
			Username: "user2",
			Password: "password2",
		}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deviceService.UpdateDevice(tt.args.ctx, tt.args.in)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.NotZero(t, got.Id)
		})
	}
}

func createDevice() *pb.CreateDeviceResponse {
	deviceModel, _ := deviceModelService.CreateDeviceModel(ctx, testDeviceModel)
	testTelemetryModel := &pb.CreateTelemetryModelRequest{
		DeviceModelId: deviceModel.Id,
		Key:           "key",
		Name:          "name",
		Factor:        1,
		Unit:          "A",
		UnitName:      "安培",
		Scale:         2,
	}
	deviceModelService.CreateTelemetryModel(ctx, testTelemetryModel)
	testDevice := &pb.CreateDeviceRequest{
		DeviceModelId: deviceModel.Id,
		Name:          "测试设备",
		Ip:            "192.168.1.111",
		Port:          3030,
		Username:      "user",
		Password:      "password",
	}
	device, _ := deviceService.CreateDevice(ctx, testDevice)
	return device
}

func TestDeviceService_ListDevice(t *testing.T) {
	device := createDevice()
	type args struct {
		ctx context.Context
		in  *pb.ListDeviceRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test ListDevice", args{context.Background(), &pb.ListDeviceRequest{
			SearchParam: &base.SearchParam{
				Ids: []int64{device.Id},
				QueryFields: []string{
					"edges",
					"edges.id",
					"edges.name",
					"edges.ip",
					"edges.port",
					"edges.accessKey",
					"edges.secretKey",
					"edges.username",
					"edges.password",
					"edges.deviceModel.id",
					"edges.deviceModel.name",
					"edges.deviceModel.desc",
					"edges.deviceModel.deviceType",
					"edges.deviceModel.cameraCompany",
					"edges.attributes",
					"edges.attributes.id",
					"edges.attributes.key",
					"edges.attributes.value",
					"edges.attributes.name",
					"edges.attributes.createdAt",
					"edges.attributes.updatedAt",
					"edges.telemetries",
					"edges.telemetries.id",
					"edges.telemetries.key",
					"edges.telemetries.value",
					"edges.telemetries.name",
					"edges.telemetries.factor",
					"edges.telemetries.unit",
					"edges.telemetries.unitName",
					"edges.telemetries.scale",
					"edges.telemetries.createdAt",
					"edges.telemetries.updatedAt",
					"edges.createdAt",
					"edges.updatedAt",
					"totalCount",
				},
			},
			DeviceType: ptrs.Int64Ptr(int64(pb.DeviceType_Camera)),
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deviceService.ListDevice(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeviceService.ListDevice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.EqualValues(t, 1, len(got.Edges))
			assert.LessOrEqual(t, int64(1), got.TotalCount)
		})
	}
}
