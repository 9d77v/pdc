package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

//Video 视频文件
type Video struct {
	gorm.Model
	Title      string `gorm:"size:50"`
	Desc       string `gorm:"size:5000"`
	PubDate    time.Time
	Cover      string `gorm:"size:500"`
	Episodes   []*Episode
	Characters postgres.Hstore
	Staffs     postgres.Hstore
	Tags       pq.StringArray `gorm:"type:varchar(10)[]"`
	IsShow     bool
}
