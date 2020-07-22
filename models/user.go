package models

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
)

//User 用户
type User struct {
	gorm.Model
	Name      string `gorm:"size:50;unique;NOT NULL;"`
	Password  string `gorm:"size:256;NOT NULL;"`
	Avatar    string `gorm:"size:200;NOT NULL;"`
	RoleID    int    //用户角色，1：owner，2：manager，3：normal user,4: guest
	Gender    int    //性别，0:male,1:female，2：unknown
	Color     string `gorm:"size:50;NOT NULL;"` //favorite color
	BirthDate time.Time
	IP        string `gorm:"size:50;NOT NULL;"` //check if user is online
}

//MarshalBinary ..
func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

//UnmarshalBinary ..
func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

//GetByID ..
func (u *User) GetByID(uid int64) error {
	return Gorm.Where("id=?", uid).First(u).Error
}
