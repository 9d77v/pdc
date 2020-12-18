package models

import (
	"testing"

	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/device-service/pb"
)

func TestDeviceModel_IsHikivisionCamera(t *testing.T) {
	type fields struct {
		DefaultModel    base.DefaultModel
		Name            string
		Desc            string
		DeviceType      uint8
		CameraCompany   uint8
		AttributeModels []*AttributeModel
		TelemetryModels []*TelemetryModel
	}
	type args struct {
		deviceType    pb.DeviceType
		cameraCompany pb.CameraCompany
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"test is HikivisionCamera", fields{}, args{pb.DeviceType_Default, pb.CameraCompany_DaHua}, false},
		{"test is not HikivisionCamera", fields{}, args{pb.DeviceType_Camera, pb.CameraCompany_Hikvision}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &DeviceModel{
				DefaultModel:    tt.fields.DefaultModel,
				Name:            tt.fields.Name,
				Desc:            tt.fields.Desc,
				DeviceType:      tt.fields.DeviceType,
				CameraCompany:   tt.fields.CameraCompany,
				AttributeModels: tt.fields.AttributeModels,
				TelemetryModels: tt.fields.TelemetryModels,
			}
			if got := m.IsHikivisionCamera(tt.args.deviceType, tt.args.cameraCompany); got != tt.want {
				t.Errorf("DeviceModel.IsHikivisionCamera() = %v, want %v", got, tt.want)
			}
		})
	}
}
