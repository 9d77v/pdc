package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
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
	MobileURL string `gorm:"size:500;"`
	Subtitles []*Subtitle
}

//Subtitle 字幕
type Subtitle struct {
	gorm.Model
	EpisodeID uint
	Name      string `gorm:"size:50;NOT NULL;"`
	URL       string `gorm:"size:500;NOT NULL;"`
}

//Video 视频
type Video struct {
	gorm.Model
	Title    string `gorm:"size:50;NOT NULL;"`
	Desc     string `gorm:"size:5000;NOT NULL;"`
	PubDate  time.Time
	Cover    string `gorm:"size:500;NOT NULL;"`
	Episodes []*Episode
	Tags     pq.StringArray `gorm:"type:varchar(10)[]"`
	IsShow   bool
}

//VideoSeriesItem 视频系列视频列表
type VideoSeriesItem struct {
	VideoSeriesID uint   `gorm:"primary_key;auto_increment:false"`
	VideoID       uint   `gorm:"primary_key;auto_increment:false"`
	Title         string `gorm:"-"`
	Alias         string
	Num           int
}

//VideoSeries 视频系列
type VideoSeries struct {
	gorm.Model
	Name  string
	Items []*VideoSeriesItem
}
