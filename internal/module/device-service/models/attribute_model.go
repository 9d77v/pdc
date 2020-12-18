package models

import (
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/device-service/pb"
)

//AttributeModel 属性模型
type AttributeModel struct {
	base.DefaultModel
	DeviceModelID uint   `gorm:"uniqueIndex:attribute_model_uix"`
	Key           string `gorm:"uniqueIndex:attribute_model_uix;size:50"`
	Name          string `gorm:"size:50"`
}

//NewAttributeModel ..
func NewAttributeModel() *AttributeModel {
	m := &AttributeModel{}
	m.SetDB(db.GetDB())
	return m
}

//NewAttributeModelFromPB ..
func NewAttributeModelFromPB(in *pb.CreateAttributeModelRequest) *AttributeModel {
	m := &AttributeModel{
		DeviceModelID: uint(in.DeviceModelId),
		Key:           in.Key,
		Name:          in.Name,
	}
	m.SetDB(db.GetDB())
	return m
}
