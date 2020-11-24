package models

import (
	"log"

	"github.com/9d77v/pdc/internal/db"
)

func init() {
	err := db.Gorm.AutoMigrate(
		&History{},
	)
	if err != nil {
		log.Println("auto migrate error:", err)
	}
}
