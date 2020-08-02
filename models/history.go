package models

import "time"

//History 历史记录
type History struct {
	UID           uint  `gorm:"primary_key;auto_increment:false"`
	SourceType    uint8 `gorm:"primary_key;auto_increment:false"` //1:video
	SourceID      uint  `gorm:"primary_key;auto_increment:false"`
	SubSourceID   uint
	CurrentTime   float64
	RemainingTime float64
	UpdatedAt     time.Time
}
