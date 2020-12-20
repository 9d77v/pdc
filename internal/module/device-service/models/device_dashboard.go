package models

import (
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/device-service/pb"
	"github.com/golang/protobuf/ptypes"
	"gorm.io/gorm"
)

//DeviceDashboard 设备仪表盘
type DeviceDashboard struct {
	base.DefaultModel
	Name        string `gorm:"size:50"`
	IsVisible   bool
	DeviceType  uint8 //设备类型，0:默认设备,1:摄像头
	Telemetries []*DeviceDashboardTelemetry
	Cameras     []*DeviceDashboardCamera
}

//NewDeviceDashboard ..
func NewDeviceDashboard() *DeviceDashboard {
	m := &DeviceDashboard{}
	m.SetDB(db.GetDB())
	return m
}

//NewDeviceDashboardFromPB ..
func NewDeviceDashboardFromPB(in *pb.CreateDeviceDashboardRequest) *DeviceDashboard {
	m := &DeviceDashboard{
		Name:       in.Name,
		IsVisible:  in.IsVisible,
		DeviceType: uint8(in.DeviceType),
	}
	m.SetDB(db.GetDB())
	return m
}

//GetByID ..
func (m *DeviceDashboard) GetByID(id uint) error {
	return m.IDQuery(id).Preload("Cameras", func(db *gorm.DB) *gorm.DB {
		return db.Model(&DeviceDashboardCamera{})
	}).Preload("Telemetries", func(db *gorm.DB) *gorm.DB {
		return db.Model(&DeviceDashboardTelemetry{})
	}).First(m)
}

//ToDeviceDashboardPBs ..
func (m *DeviceDashboard) ToDeviceDashboardPBs(data []*DeviceDashboard) []*pb.DeviceDashboard {
	result := make([]*pb.DeviceDashboard, 0, len(data))
	for _, v := range data {
		r := m.toDeviceDashboardPB(v)
		result = append(result, r)
	}
	return result
}
func (m *DeviceDashboard) toDeviceDashboardPB(deviceDashboard *DeviceDashboard) *pb.DeviceDashboard {
	ts := make([]*pb.DeviceDashboardTelemetry, 0, len(deviceDashboard.Telemetries))
	for _, v := range deviceDashboard.Telemetries {
		createdAt, _ := ptypes.TimestampProto(v.CreatedAt)
		updatedAt, _ := ptypes.TimestampProto(v.UpdatedAt)
		ts = append(ts, &pb.DeviceDashboardTelemetry{
			Id:                int64(v.ID),
			DeviceDashboardId: int64(v.DeviceDashboardID),
			DeviceId:          int64(v.Telemetry.DeviceID),
			DeviceName:        v.Telemetry.Device.Name,
			TelemetryId:       int64(v.TelemetryID),
			Key:               v.Telemetry.TelemetryModel.Key,
			Name:              v.Telemetry.TelemetryModel.Name,
			Unit:              v.Telemetry.TelemetryModel.Unit,
			UnitName:          v.Telemetry.TelemetryModel.UnitName,
			Factor:            v.Telemetry.TelemetryModel.Factor,
			Scale:             int64(v.Telemetry.TelemetryModel.Scale),
			CreatedAt:         createdAt,
			UpdatedAt:         updatedAt,
		})
	}
	cs := make([]*pb.DeviceDashboardCamera, 0, len(deviceDashboard.Cameras))
	for _, v := range deviceDashboard.Cameras {
		createdAt, _ := ptypes.TimestampProto(v.CreatedAt)
		updatedAt, _ := ptypes.TimestampProto(v.UpdatedAt)
		cs = append(cs, &pb.DeviceDashboardCamera{
			Id:                int64(v.ID),
			DeviceDashboardId: int64(v.DeviceDashboardID),
			DeviceId:          int64(v.DeviceID),
			DeviceName:        v.Device.Name,
			CreatedAt:         createdAt,
			UpdatedAt:         updatedAt,
		})
	}
	createdAt, _ := ptypes.TimestampProto(deviceDashboard.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(deviceDashboard.UpdatedAt)
	return &pb.DeviceDashboard{
		Id:          int64(deviceDashboard.ID),
		Name:        deviceDashboard.Name,
		IsVisible:   deviceDashboard.IsVisible,
		Telemetries: ts,
		Cameras:     cs,
		DeviceType:  pb.DeviceType(deviceDashboard.DeviceType),
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}
