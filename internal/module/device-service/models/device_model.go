package models

import (
	"github.com/9d77v/pdc/internal/module/base"
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
	vs := &DeviceModel{}
	vs.DefaultModel = base.NewDefaultModel()
	return vs
}

//GetByID ..
func (m *DeviceModel) GetByID(id uint, columns []string) error {
	return m.Select(columns).IDQuery(id).First(m)
}
