package models

import (
	"github.com/9d77v/pdc/internal/module/base"
)

//TelemetryModel 遥测模型
type TelemetryModel struct {
	base.DefaultModel
	DeviceModelID uint   `gorm:"uniqueIndex:telemetry_model_uix"`
	Key           string `gorm:"uniqueIndex:telemetry_model_uix;size:50"`
	Name          string `gorm:"size:50"`
	Factor        float64
	Unit          string `gorm:"size:10"`
	UnitName      string `gorm:"size:10"`
	Scale         uint8
}

//NewTelemetryModel ..
func NewTelemetryModel() *TelemetryModel {
	vs := &TelemetryModel{}
	vs.DefaultModel = base.NewDefaultModel()
	return vs
}
