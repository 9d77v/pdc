package services

import (
	"context"
	"testing"

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
