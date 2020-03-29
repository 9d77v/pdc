package models

import (
	"github.com/jinzhu/gorm"
)

//Episode 分集
type Episode struct {
	gorm.Model
	MediaID  int64
	Order    int64
	Title    string `gorm:"size:50"`
	Desc     string `gorm:"size:5000"`
	Cover    string `gorm:"size:200"`
	URL      string `gorm:"size:200"`
	Subtitle string `gorm:"size:200"`
}
