package models

import (
	"gorm.io/gorm"
)

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

//DeviceDashboard 设备仪表盘
type DeviceDashboard struct {
	gorm.Model
	Name        string `gorm:"size:50"`
	IsVisible   bool
	DeviceType  uint8 //设备类型，0:默认设备,1:摄像头
	Telemetries []*DeviceDashboardTelemetry
	Cameras     []*DeviceDashboardCamera
}

//DeviceDashboardTelemetry 仪表盘遥测
type DeviceDashboardTelemetry struct {
	gorm.Model
	DeviceDashboardID uint
	TelemetryID       uint
	Telemetry         Telemetry
}

//DeviceDashboardCamera 仪表盘摄像头
type DeviceDashboardCamera struct {
	gorm.Model
	DeviceDashboardID uint
	DeviceID          uint
	Device            Device
}
