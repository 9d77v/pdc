package models

import (
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
)

//Attribute 属性
type Attribute struct {
	base.DefaultModel
	DeviceID         uint `gorm:"uniqueIndex:attribute_uix"`
	AttributeModelID uint `gorm:"uniqueIndex:attribute_uix"`
	AttributeModel   AttributeModel
	Value            string `gorm:"size:50;NOT NULL;"`
}

//NewAttribute ..
func NewAttribute() *Attribute {
	m := &Attribute{}
	m.SetDB(db.GetDB())
	return m
}

//TableName ..
func (m *Attribute) TableName() string {
	return db.TablePrefix() + "attribute"
}
