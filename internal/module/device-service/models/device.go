package models

import (
	"github.com/9d77v/pdc/internal/module/base"
)

//Device 设备
type Device struct {
	base.DefaultModel
	DeviceModelID uint
	DeviceModel   DeviceModel
	Name          string `gorm:"size:50;NOT NULL;"`
	IP            string `gorm:"size:50;NOT NULL;"`
	AccessKey     string `gorm:"size:12;NOT NULL;"` //设备key hashid(id,16)
	SecretKey     string `gorm:"size:32;NOT NULL;"` //设备访问密钥,randString()
	Port          uint16
	Username      string       `gorm:"size:32;"`
	Password      string       `gorm:"size:32;"`
	Attributes    []*Attribute //属性，由采集程序启动时注册
	Telemetries   []*Telemetry //遥测，由采集程序按一定频率上传
}

//NewDevice ..
func NewDevice() *Device {
	vs := &Device{}
	vs.DefaultModel = base.NewDefaultModel()
	return vs
}

//GetByID ..
func (m *Device) GetByID(id uint, columns []string) error {
	return m.Select(columns).IDQuery(id).First(m)
}

//Attribute 属性
type Attribute struct {
	base.DefaultModel
	DeviceID         uint `gorm:"uniqueIndex:attribute_uix"`
	AttributeModelID uint `gorm:"uniqueIndex:attribute_uix"`
	AttributeModel   AttributeModel
	Value            string `gorm:"size:50;NOT NULL;"`
}

//Telemetry 遥测
type Telemetry struct {
	*base.DefaultModel
	DeviceID         uint `gorm:"uniqueIndex:telemetry_uix"`
	Device           Device
	TelemetryModelID uint `gorm:"uniqueIndex:telemetry_uix"`
	TelemetryModel   TelemetryModel
}
