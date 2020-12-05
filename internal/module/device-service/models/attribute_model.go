package models

import (
	"github.com/9d77v/pdc/internal/module/base"
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
	vs := &AttributeModel{}
	vs.DefaultModel = base.NewDefaultModel()
	return vs
}

//GetByID ..
func (m *AttributeModel) GetByID(id uint, columns []string) error {
	return m.Select(columns).IDQuery(id).First(m)
}
