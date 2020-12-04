package models

import (
	"github.com/9d77v/pdc/internal/module/base"
)

//TelemetryModel 遥测模型
type TelemetryModel struct {
	*base.Model
	DeviceModelID uint   `gorm:"unique_index:telemetry_model_uix"`
	Key           string `gorm:"unique_index:telemetry_model_uix;size:50"`
	Name          string `gorm:"size:50"`
	Factor        float64
	Unit          string `gorm:"size:10"`
	UnitName      string `gorm:"size:10"`
	Scale         uint8
}

//NewTelemetryModel ..
func NewTelemetryModel() *TelemetryModel {
	vs := &TelemetryModel{}
	vs.Model = base.NewModel()
	return vs
}

//GetByID ..
func (m *TelemetryModel) GetByID(id uint, columns []string) error {
	return m.Select(columns).IDQuery(id).First(m)
}
