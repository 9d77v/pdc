package models

import (
	"gorm.io/gorm"
)

//Device 设备
type Device struct {
	gorm.Model
	DeviceModelID uint //设备模型id
	DeviceModel   DeviceModel
	Name          string       `gorm:"size:50;NOT NULL;"` //设备名称
	IP            string       `gorm:"size:50;NOT NULL;"` //设备ip
	AccessKey     string       `gorm:"size:12;NOT NULL;"` //设备key hashid(id,16)
	SecretKey     string       `gorm:"size:32;NOT NULL;"` //设备访问密钥,randString()
	Port          uint16       //设备端口
	Username      string       `gorm:"size:32;"`
	Password      string       `gorm:"size:32;"`
	LivePath      string       `gorm:"size:100"` //直播路径
	Attributes    []*Attribute //属性，由采集程序启动时注册
	Telemetries   []*Telemetry //遥测，由采集程序按一定频率上传
}

//Attribute 属性
type Attribute struct {
	gorm.Model
	DeviceID         uint `gorm:"unique_index:attribute_uix"` //设备模型id
	AttributeModelID uint `gorm:"unique_index:attribute_uix"`
	AttributeModel   AttributeModel
	Value            string `gorm:"size:50;NOT NULL;"`
}

//Telemetry 遥测
type Telemetry struct {
	gorm.Model
	DeviceID         uint `gorm:"unique_index:telemetry_uix"` //设备id
	Device           Device
	TelemetryModelID uint `gorm:"unique_index:telemetry_uix"`
	TelemetryModel   TelemetryModel
}
