package models

import (
	"time"

	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/lib/pq"
)

//Thing 物品
type Thing struct {
	base.DefaultModel
	UID                 int64  `grom:"column:uid"`
	Name                string `gorm:"size:100"`
	Num                 float64
	BrandName           string         `gorm:"size:50"`
	Pics                pq.StringArray `gorm:"type:varchar(200)[]"`
	UnitPrice           float64        `gorm:"type:money"`
	Unit                string         `gorm:"size:50"`  //数量单位
	Specifications      string         `gorm:"size:100"` //规格
	Category            int8           //0:未知，1：摄像头
	ConsumerExpenditure string         `gorm:"size:2"` //居民消费支出分类,01:食品烟酒，02：衣着，03：居住，04:生活用品及服务，05：交通通信，06：教育文化娱乐，07：医疗保健，08：其他用品及服务
	Location            string         `gorm:"size:50"`
	PurchaseDate        time.Time
	Status              int8          `gorm:"size:50"` //0：待采购，1：使用中，2：已收纳，3.故障，4.维修中，5.待清理，6.已清理
	PurchasePlatform    string        `gorm:"size:50"`
	RefOrderID          string        `gorm:"size:50"`
	RubbishCategory     pq.Int64Array `gorm:"type:smallint[]"` //0:其他垃圾，1：可回收垃圾，2：厨余垃圾，3：有害垃圾
}

//NewThing ..
func NewThing() *Thing {
	vs := &Thing{}
	vs.SetDB(db.GetDB())
	return vs
}

//PieLineSerie ..
type PieLineSerie struct {
	X1 string  `json:"x1"`
	X2 string  `json:"x2"`
	Y  float64 `json:"y"`
}
