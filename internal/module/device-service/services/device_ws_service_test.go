package services

import (
	"context"
	"reflect"
	"testing"

	"github.com/9d77v/pdc/internal/module/device-service/models"
	devicePB "github.com/9d77v/pdc/internal/module/device-service/pb"
)

func TestDeviceService_DeviceLogin(t *testing.T) {
	deviceResp := createDevice()
	device := models.NewDevice()
	device.GetByID(uint(deviceResp.Id))
	type args struct {
		accessKey string
		secretKey string
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{"test  login success", args{device.AccessKey, device.SecretKey}, device.ID, false},
		{"test  login failed", args{"", ""}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deviceService.DeviceLogin(tt.args.accessKey, tt.args.secretKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeviceService.DeviceLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeviceService.DeviceLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeviceService_GetDeviceInfo(t *testing.T) {
	deviceResp := createDevice()
	device := models.NewDevice()
	device.GetByID(uint(deviceResp.Id))
	type args struct {
		deviceID uint32
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		{"test  GetDeviceInfo success", args{uint32(device.ID)}, uint32(device.ID), false},
		{"test  GetDeviceInfo failed", args{uint32(device.ID + 1)}, uint32(0), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deviceService.GetDeviceInfo(tt.args.deviceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeviceService.GetDeviceInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.DeviceId, tt.want) {
				t.Errorf("DeviceService.GetDeviceInfo() = %v, want %v", got.DeviceId, tt.want)
			}
		})
	}
}

func TestDeviceService_CameraCapture(t *testing.T) {
	deviceResp := createDevice()
	device := models.NewDevice()
	device.GetByID(uint(deviceResp.Id))
	type args struct {
		ctx context.Context
		in  *devicePB.CameraCaptureRequest
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"test  deviceID not exist", args{ctx, &devicePB.CameraCaptureRequest{
			DeviceId: 0,
			Scheme:   "http",
		}}, "", true},
		{"test  device not online", args{ctx, &devicePB.CameraCaptureRequest{
			DeviceId: uint32(device.ID),
			Scheme:   "http",
		}}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := deviceService.CameraCapture(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeviceService.CameraCapture() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.ImageUrl, tt.want) {
				t.Errorf("DeviceService.CameraCapture() = %v, want %v", got, tt.want)
			}
		})
	}
}
