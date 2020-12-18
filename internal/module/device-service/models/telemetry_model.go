package models

import (
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/device-service/pb"
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
	m := &TelemetryModel{}
	m.SetDB(db.GetDB())
	return m
}

//NewTelemetryModelFromPB ..
func NewTelemetryModelFromPB(in *pb.CreateTelemetryModelRequest) *TelemetryModel {
	m := &TelemetryModel{
		DeviceModelID: uint(in.DeviceModelId),
		Key:           in.Key,
		Name:          in.Name,
		Factor:        in.Factor,
		Unit:          in.Unit,
		UnitName:      in.UnitName,
		Scale:         uint8(in.Scale),
	}
	m.SetDB(db.GetDB())
	return m
}
