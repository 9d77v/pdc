package services

import (
	"context"
	"testing"
)

func TestDeviceService_CameraCapture(t *testing.T) {
	type args struct {
		ctx      context.Context
		deviceID int64
		scheme   string
	}
	tests := []struct {
		name    string
		s       DeviceService
		args    args
		want    string
		wantErr bool
	}{
		{"测试图片抓拍", DeviceService{}, args{context.Background(), 1, "http"}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DeviceService{}
			got, err := s.CameraCapture(tt.args.ctx, tt.args.deviceID, tt.args.scheme)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeviceService.CameraCapture() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeviceService.CameraCapture() = %v, want %v", got, tt.want)
			}
		})
	}
}
