package models

import (
	"time"

	"github.com/9d77v/pdc/internal/module/base"
)

//History 历史记录
type History struct {
	*base.Model
	UID           uint  `gorm:"unique_index:history_uix"`
	SourceType    uint8 `gorm:"unique_index:history_uix"` //1:video
	SourceID      uint  `gorm:"unique_index:history_uix"`
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

//GetByID ..
func (m *History) GetByID(id uint, columns []string) error {
	return m.Select(columns).IDQuery(id).First(m)
}

//HistoryLog 历史记录
type HistoryLog struct {
	*base.Model
	HistoryID     uint
	SubSourceID   uint
	Platform      string
	CurrentTime   float64
	RemainingTime float64
}
