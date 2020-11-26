package models

import (
	"log"

	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/utils"
)

var (
	ownerName     = utils.GetEnvStr("ADMIN_NAME", "admin")
	ownerPassword = utils.GetEnvStr("ADMIN_PASSWORD", "123456")
)

func init() {
	err := db.Gorm.AutoMigrate(
		&User{},
	)
	if err != nil {
		log.Println("auto migrate error:", err)
	}
	new(User).generateAdminAccount(ownerName, ownerPassword)
}
