package models

import (
	"time"

	"github.com/9d77v/pdc/internal/module/base"
	"github.com/lib/pq"
)

//VideoRepository ..
type VideoRepository interface {
	GetByID(id uint, columns []string) error
}

//Video 视频
type Video struct {
	base.DefaultModel
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

//NewVideo ..
func NewVideo() *Video {
	vs := &Video{}
	vs.DefaultModel = base.NewDefaultModel()
	return vs
}

//GetByID Get video by id
func (m *Video) GetByID(id uint, columns []string) error {
	return m.Select(columns).IDQuery(id).First(m)
}
