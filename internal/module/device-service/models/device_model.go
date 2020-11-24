package models

import "gorm.io/gorm"

//DeviceModel 设备模型
type DeviceModel struct {
	gorm.Model
	Name            string            `gorm:"size:50"`   //设备名称
	Desc            string            `gorm:"size:5000"` //设备描述
	DeviceType      uint8             //设备类型，0:默认设备,1:摄像头
	CameraCompany   uint8             //摄像头厂家，0：海康威视，1：大华
	AttributeModels []*AttributeModel //属性模型
	TelemetryModels []*TelemetryModel //遥测模型
}

//AttributeModel 属性模型
type AttributeModel struct {
	gorm.Model
	DeviceModelID uint   `gorm:"unique_index:attribute_model_uix"`         //设备模型id
	Key           string `gorm:"unique_index:attribute_model_uix;size:50"` //设备模型下唯一
	Name          string `gorm:"size:50"`                                  //名称
}

//TelemetryModel 遥测模型
type TelemetryModel struct {
	gorm.Model
	DeviceModelID uint    `gorm:"unique_index:telemetry_model_uix"`         //设备模型id
	Key           string  `gorm:"unique_index:telemetry_model_uix;size:50"` //设备模型下唯一
	Name          string  `gorm:"size:50"`                                  //名称
	Factor        float64 //系数
	Unit          string  `gorm:"size:10"` //单位
	UnitName      string  `gorm:"size:10"` //单位名称
	Scale         uint8   //小数位数
}
