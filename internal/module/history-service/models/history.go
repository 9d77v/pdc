package models

import (
	"time"

	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
)

//history category
const (
	HistoryNull = iota
	HistoryVideo
)

//History 历史记录
type History struct {
	*base.Model
	UID           uint  `gorm:"primary_key;auto_increment:false"`
	SourceType    uint8 `gorm:"primary_key;auto_increment:false"` //1:video
	SourceID      uint  `gorm:"primary_key;auto_increment:false"`
	SubSourceID   uint
	Platform      string
	CurrentTime   float64
	RemainingTime float64
	UpdatedAt     time.Time
}

//NewHistory ..
func NewHistory() *History {
	vs := &History{}
	vs.Model = base.NewModel()
	return vs
}

//TableName ..
func (m *History) TableName() string {
	return db.TablePrefix() + "history"
}

//HistoryLog 历史记录
type HistoryLog struct {
	*base.Model
	ID            uint  `gorm:"primarykey"`
	UID           uint  `gorm:"primary_key;auto_increment:false"`
	SourceType    uint8 `gorm:"primary_key;auto_increment:false"` //1:video
	SourceID      uint  `gorm:"primary_key;auto_increment:false"`
	SubSourceID   uint
	Platform      string
	CurrentTime   float64
	RemainingTime float64
	CreatedAt     time.Time
}
