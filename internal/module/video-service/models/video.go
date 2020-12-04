package models

import (
	"time"

	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/utils"
	"github.com/lib/pq"
)

//Video 视频
type Video struct {
	*base.Model
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
	vs.Model = base.NewModel()
	return vs
}

//GetByID Get video by id
func (v *Video) GetByID(id uint, columns []string) error {
	return db.GetDB().Select(utils.ToDBFields(columns)).
		First(v, "id=?", id).Error
}
