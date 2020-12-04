package models

import (
	"github.com/9d77v/pdc/internal/module/base"
)

//VideoSeriesItem 视频系列视频列表
type VideoSeriesItem struct {
	*base.Model
	VideoSeriesID uint   `gorm:"unique_index:video_series_item_uix"`
	VideoID       uint   `gorm:"unique_index:video_series_item_uix"`
	EpisodeID     uint   `gorm:"<-:false"`
	Title         string `gorm:"<-:false"`
	Alias         string
	Num           int
}

//NewVideoSeriesItem ..
func NewVideoSeriesItem() *VideoSeriesItem {
	vs := &VideoSeriesItem{}
	vs.Model = base.NewModel()
	return vs
}

//GetByVideoID ..
func (m *VideoSeriesItem) GetByVideoID(videoID uint) error {
	return m.Select([]string{"video_series_id"}).
		IDQuery(videoID, "video_id").
		Take(m)
}
