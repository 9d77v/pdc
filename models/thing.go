package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

//Thing 物品
type Thing struct {
	gorm.Model
	UID              int64  `grom:"column:uid"`
	Name             string `gorm:"size:50"`
	Desc             string `gorm:"size:5000"`
	Num              float64
	BrandName        string         `gorm:"size:50"`
	Pics             pq.StringArray `gorm:"type:varchar(200)[]"`
	UnitPrice        float64        `gorm:"type:money"`
	Unit             string         `gorm:"size:50"`
	Category         string         `gorm:"size:50"`
	Location         string         `gorm:"size:50"`
	PurchaseDate     time.Time
	Status           int8           `gorm:"size:50"`
	PurchasePlatform string         `gorm:"size:50"`
	RefOrderID       string         `gorm:"size:50"`
	RubbishCategory  pq.StringArray `gorm:"type:varchar(10)[]"`
}
