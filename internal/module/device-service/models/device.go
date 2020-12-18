package models

import (
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/device-service/pb"
	"gorm.io/gorm"
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
	m := &Device{}
	m.SetDB(db.GetDB())
	return m
}

//NewDeviceFromPB ..
func NewDeviceFromPB(in *pb.CreateDeviceRequest) *Device {
	m := &Device{
		DeviceModelID: uint(in.DeviceModelId),
		Name:          in.Name,
		IP:            in.Ip,
		Port:          uint16(in.Port),
		Username:      in.Username,
		Password:      in.Password,
	}
	m.SetDB(db.GetDB())
	return m
}

//TableName ..
func (m *Device) TableName() string {
	return db.TablePrefix() + "device"
}

//SaveAttributesFromAttributeModels ..
func (m *Device) SaveAttributesFromAttributeModels(
	attributeModels []*AttributeModel) (err error) {
	if len(attributeModels) > 0 {
		attributes := make([]*Attribute, 0, len(attributeModels))
		for _, v := range attributeModels {
			attributes = append(attributes, &Attribute{
				DeviceID:         m.ID,
				AttributeModelID: v.ID,
			})
		}
		return m.Create(&attributes)
	}
	return nil
}

//SaveTelemetriesFromTelemetryModels ..
func (m *Device) SaveTelemetriesFromTelemetryModels(
	telemetryModels []*TelemetryModel) (err error) {
	if len(telemetryModels) > 0 {
		telemetries := make([]*Telemetry, 0, len(telemetryModels))
		for _, v := range telemetryModels {
			telemetries = append(telemetries, &Telemetry{
				DeviceID:         m.ID,
				TelemetryModelID: v.ID,
			})
		}
		return m.Create(&telemetries)
	}
	return nil
}

//SelectDeviceType ..
func (m *Device) SelectDeviceType(deviceType int64) {
	tableDeviceModel := new(DeviceModel).TableName()
	m.Where(tableDeviceModel+".device_type = ?", deviceType).
		LeftJoin(tableDeviceModel + " ON " + tableDeviceModel + ".id = " +
			m.TableName() + ".device_model_id")
}

//GetByID ..
func (m *Device) GetByID(id uint) error {
	return m.IDQuery(id).Preload("Attributes", func(db *gorm.DB) *gorm.DB {
		return db.Model(&Attribute{})
	}).Preload("Telemetries", func(db *gorm.DB) *gorm.DB {
		return db.Model(&Telemetry{})
	}).First(m)
}
