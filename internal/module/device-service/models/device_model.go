package models

import (
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/device-service/pb"
	"github.com/golang/protobuf/ptypes"
	"gorm.io/gorm"
)

//DeviceModel 设备模型
type DeviceModel struct {
	base.DefaultModel
	Name            string `gorm:"size:50"`
	Desc            string `gorm:"size:5000"`
	DeviceType      uint8  //设备类型，0:默认设备,1:摄像头
	CameraCompany   uint8  //摄像头厂家，0：海康威视，1：大华
	AttributeModels []*AttributeModel
	TelemetryModels []*TelemetryModel
}

//NewDeviceModel ..
func NewDeviceModel() *DeviceModel {
	m := &DeviceModel{}
	m.SetDB(db.GetDB())
	return m
}

//NewDeviceModelFromPB ..
func NewDeviceModelFromPB(in *pb.CreateDeviceModelRequest) *DeviceModel {
	m := &DeviceModel{
		Name:          in.Name,
		Desc:          in.Desc,
		DeviceType:    uint8(in.DeviceType),
		CameraCompany: uint8(in.CameraCompany),
	}
	m.SetDB(db.GetDB())
	return m
}

//TableName ..
func (m *DeviceModel) TableName() string {
	return db.TablePrefix() + "device_model"
}

//GetByID ..
func (m *DeviceModel) GetByID(id uint) error {
	return m.IDQuery(id).Preload("AttributeModels", func(db *gorm.DB) *gorm.DB {
		return db.Model(&AttributeModel{})
	}).Preload("TelemetryModels", func(db *gorm.DB) *gorm.DB {
		return db.Model(&TelemetryModel{})
	}).First(m)
}

//IsHikivisionCamera ..
func (m *DeviceModel) IsHikivisionCamera(deviceType pb.DeviceType, cameraCompany pb.CameraCompany) bool {
	return deviceType == pb.DeviceType_Camera && cameraCompany == pb.CameraCompany_Hikvision
}

//ToDeviceModelPBs ..
func (m *DeviceModel) ToDeviceModelPBs(data []*DeviceModel) []*pb.DeviceModel {
	result := make([]*pb.DeviceModel, 0, len(data))
	for _, v := range data {
		r := m.toDeviceModelPB(v)
		result = append(result, r)
	}
	return result
}

func (m *DeviceModel) toDeviceModelPB(deviceModel *DeviceModel) *pb.DeviceModel {
	as := make([]*pb.AttributeModel, 0, len(deviceModel.AttributeModels))
	for _, v := range deviceModel.AttributeModels {
		createdAt, _ := ptypes.TimestampProto(v.CreatedAt)
		updatedAt, _ := ptypes.TimestampProto(v.UpdatedAt)
		as = append(as, &pb.AttributeModel{
			Id:            int64(v.ID),
			DeviceModelId: int64(v.DeviceModelID),
			Key:           v.Key,
			Name:          v.Name,
			CreatedAt:     createdAt,
			UpdatedAt:     updatedAt,
		})
	}
	ts := make([]*pb.TelemetryModel, 0, len(deviceModel.TelemetryModels))
	for _, v := range deviceModel.TelemetryModels {
		createdAt, _ := ptypes.TimestampProto(v.CreatedAt)
		updatedAt, _ := ptypes.TimestampProto(v.UpdatedAt)
		ts = append(ts, &pb.TelemetryModel{
			Id:            int64(v.ID),
			DeviceModelId: int64(v.DeviceModelID),
			Key:           v.Key,
			Name:          v.Name,
			Factor:        v.Factor,
			Unit:          v.Unit,
			UnitName:      v.UnitName,
			Scale:         int64(v.Scale),
			CreatedAt:     createdAt,
			UpdatedAt:     updatedAt,
		})
	}
	createdAt, _ := ptypes.TimestampProto(deviceModel.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(deviceModel.UpdatedAt)
	return &pb.DeviceModel{
		Id:              int64(deviceModel.ID),
		Name:            deviceModel.Name,
		Desc:            deviceModel.Desc,
		DeviceType:      pb.DeviceType(deviceModel.DeviceType),
		AttributeModels: as,
		TelemetryModels: ts,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
}
