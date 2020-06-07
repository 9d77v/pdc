package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//User 用户
type User struct {
	gorm.Model
	Name      string `gorm:"size:50"`
	Password  string `gorm:"size:256"`
	Avatar    string `gorm:"size:200"`
	RoleID    int    //用户角色，1：owner，2：manager，3：normal user
	Gender    int    //性别，0：unknown,1:male,2:female
	Color     string `gorm:"size:50"` //favorite color
	BirthDate time.Time
	IP        string `gorm:"size:50"` //check if user is online
}
