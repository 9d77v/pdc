package models

import (
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/device-service/pb"
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
