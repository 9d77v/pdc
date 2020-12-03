package models

import (
	"time"

	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/utils"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

//Video 视频
type Video struct {
	gorm.Model
	Title          string `gorm:"size:50;NOT NULL;"`
	Desc           string `gorm:"size:5000;NOT NULL;"`
	PubDate        time.Time
	Cover          string `gorm:"size:500;NOT NULL;"`
	Episodes       []*Episode
	Tags           pq.StringArray `gorm:"type:varchar(10)[]"`
	IsShow         bool
	IsHideOnMobile bool
	Theme          string `gorm:"size:50;"`
}

//GetByID Get video by id
func (v *Video) GetByID(id uint, columns []string) error {
	return db.GetDB().Select(utils.ToDBFields(columns)).
		First(v, "id=?", id).Error
}

//VideoSeriesItem 视频系列视频列表
type VideoSeriesItem struct {
	VideoSeriesID uint   `gorm:"primary_key;auto_increment:false"`
	VideoID       uint   `gorm:"primary_key;auto_increment:false"`
	EpisodeID     uint   `gorm:"<-:false"`
	Title         string `gorm:"<-:false"`
	Alias         string
	Num           int
}

//VideoSeries 视频系列
type VideoSeries struct {
	gorm.Model
	Name  string
	Items []*VideoSeriesItem
}
