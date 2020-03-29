package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

//Media 媒体文件
type Media struct {
	gorm.Model
	Title      string `gorm:"size:50"`
	Desc       string `gorm:"size:5000"`
	PubDate    time.Time
	Cover      string `gorm:"size:200"`
	Episodes   []*Episode
	Characters postgres.Hstore
	Staffs     postgres.Hstore
}
