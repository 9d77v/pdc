package models

import (
	"log"

	"github.com/9d77v/pdc/internal/db/db"
)

func init() {
	err := db.Gorm.AutoMigrate(
		&Video{},
		&Episode{},
		&Subtitle{},
		&VideoSeries{},
		&VideoSeriesItem{},
	)
	if err != nil {
		log.Println("auto migrate error:", err)
	}
}
