package models

import (
	"gorm.io/gorm"
)

//DeviceModel 设备模型
type DeviceModel struct {
	gorm.Model
	Name            string            `gorm:"size:50"`   //设备名称
	Desc            string            `gorm:"size:5000"` //设备描述
	DeviceType      uint8             //设备类型，0:默认设备
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

//Device 设备
type Device struct {
	gorm.Model
	DeviceModelID uint //设备模型id
	DeviceModel   DeviceModel
	Name          string       `gorm:"size:50"` //设备名称
	Attributes    []*Attribute //属性，由采集程序启动时注册
	Telemetries   []*Telemetry //遥测，由采集程序按一定频率上传
}

//Attribute 属性
type Attribute struct {
	gorm.Model
	DeviceID         uint `gorm:"unique_index:attribute_uix"` //设备模型id
	AttributeModelID uint
	AttributeModel   AttributeModel
	Key              string `gorm:"unique_index:attribute_uix;size:50"` //属性key
	Name             string `gorm:"-"`
	Value            string `gorm:"size:50"` //属性值
}

//Telemetry 遥测
type Telemetry struct {
	gorm.Model
	DeviceID         uint `gorm:"unique_index:telemetry_uix"` //设备id
	TelemetryModelID uint
	TelemetryModel   TelemetryModel
	Key              string `gorm:"unique_index:telemetry_uix;size:50"` //遥测key
	Name             string `gorm:"-"`
}
