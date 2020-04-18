package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

//Episode 分集
type Episode struct {
	gorm.Model
	VideoID   int64
	Num       float64
	Title     string `gorm:"size:50"`
	Desc      string `gorm:"size:5000"`
	Cover     string `gorm:"size:500"`
	URL       string `gorm:"size:500"`
	Subtitles postgres.Hstore
}
