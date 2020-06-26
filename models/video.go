package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

//Episode 分集
type Episode struct {
	gorm.Model
	VideoID   int64
	Num       float64
	Title     string `gorm:"size:50;NOT NULL;"`
	Desc      string `gorm:"size:5000;NOT NULL;"`
	Cover     string `gorm:"size:500;NOT NULL;"`
	URL       string `gorm:"size:500;NOT NULL;"`
	Subtitles postgres.Hstore
}

//Video 视频
type Video struct {
	gorm.Model
	Title      string `gorm:"size:50;NOT NULL;"`
	Desc       string `gorm:"size:5000;NOT NULL;"`
	PubDate    time.Time
	Cover      string `gorm:"size:500;NOT NULL;"`
	Episodes   []*Episode
	Characters postgres.Hstore
	Staffs     postgres.Hstore
	Tags       pq.StringArray `gorm:"type:varchar(10)[]"`
	IsShow     bool
}
