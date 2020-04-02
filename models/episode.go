package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

//Episode 分集
type Episode struct {
	gorm.Model
	MediaID   int64
	Order     string `gorm:"size:10"`
	Title     string `gorm:"size:50"`
	Desc      string `gorm:"size:5000"`
	Cover     string `gorm:"size:200"`
	URL       string `gorm:"size:200"`
	Subtitles postgres.Hstore
}
