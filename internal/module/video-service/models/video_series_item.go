package models

import (
	"time"

	"github.com/9d77v/pdc/internal/module/base"
)

//VideoSeriesItemRepository ..
type VideoSeriesItemRepository interface {
	GetByVideoID(videoID uint) error
	GetByVideoIDVideoSeriesID(fields []string, videoID uint, videoSeriesID uint) error
	GetTheMaxNumOfOneVideoSeries(videoSeriesID uint) (int, error)
}

//VideoSeriesItem 视频系列视频列表
type VideoSeriesItem struct {
	*base.Model
	VideoSeriesID uint   `gorm:"primary_key;auto_increment:false"`
	VideoID       uint   `gorm:"primary_key;auto_increment:false"`
	EpisodeID     uint   `gorm:"<-:false"`
	Title         string `gorm:"<-:false"`
	Alias         string
	Num           int
	CreatedAt     time.Time
	UpdatedAt     time.Time
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

//GetByVideoIDVideoSeriesID ..
func (m *VideoSeriesItem) GetByVideoIDVideoSeriesID(fields []string, videoID uint, videoSeriesID uint) error {
	return m.Select(fields).
		IDQuery(videoID, "video_id").
		IDQuery(videoSeriesID, "video_series_id").
		First(m)
}

//GetTheMaxNumOfOneVideoSeries ..
func (m *VideoSeriesItem) GetTheMaxNumOfOneVideoSeries(videoSeriesID uint) (int, error) {
	err := m.
		Select([]string{"max(num) num"}).
		IDQuery(videoSeriesID, "video_series_id").
		Take(m)
	return m.Num, err
}
