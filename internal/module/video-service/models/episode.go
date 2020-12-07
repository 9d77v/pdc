package models

import (
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
	"gorm.io/gorm"
)

//Episode 分集
type Episode struct {
	base.DefaultModel
	VideoID   uint
	Num       float64
	Title     string `gorm:"size:50;NOT NULL;"`
	Desc      string `gorm:"size:5000;NOT NULL;"`
	Cover     string `gorm:"size:500;NOT NULL;"`
	URL       string `gorm:"size:500;NOT NULL;"`
	Subtitles []*Subtitle
}

//Subtitle 字幕
type Subtitle struct {
	gorm.Model
	EpisodeID uint
	Name      string `gorm:"size:50;NOT NULL;"`
	URL       string `gorm:"size:500;NOT NULL;"`
}

//NewEpisode ..
func NewEpisode() *Episode {
	vs := &Episode{}
	vs.DefaultModel = base.NewDefaultModel()
	return vs
}

//TableName ..
func (m *Episode) TableName() string {
	return db.TablePrefix() + "episode"
}

//GetVideoIDByID ..
func (m *Episode) GetVideoIDByID(id uint) uint {
	err := m.Select([]string{"video_id"}).IDQuery(id).First(m)
	if err != nil {
		return 0
	}
	return m.VideoID
}
