package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
	"golang.org/x/crypto/bcrypt"
)

//User 用户
type User struct {
	base.DefaultModel
	Name      string `gorm:"size:50;unique;NOT NULL;"`
	Password  string `gorm:"size:256;NOT NULL;"`
	Avatar    string `gorm:"size:200;NOT NULL;"`
	RoleID    int    //用户角色，1：owner，2：manager，3：normal user,4: guest
	Gender    int    //性别，0:male,1:female，2：unknown
	Color     string `gorm:"size:50;NOT NULL;"` //favorite color
	BirthDate time.Time
	IP        string `gorm:"size:50;NOT NULL;"` //check if user is online
}

//NewUser ..
func NewUser() *User {
	vs := &User{}
	vs.DefaultModel = base.NewDefaultModel()
	return vs
}

//TableName ..
func (m *User) TableName() string {
	return db.TablePrefix() + "user"
}

//MarshalBinary ..
func (m *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

//UnmarshalBinary ..
func (m *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

//GenerateAdminAccount ..
func (m *User) GenerateAdminAccount(ownerName, ownerPassword string) {
	var total int64
	err := db.GetDB().Model(&User{}).Count(&total).Error
	if err != nil {
		log.Panicf("Get User total failed:%v/n", err)
	}
	if total == 0 {
		bytes, err := bcrypt.GenerateFromPassword([]byte(ownerPassword), 12)
		if err != nil {
			log.Panicf("generate password failed:%v/n", err)
		}
		user := &User{
			Name:     ownerName,
			Password: string(bytes),
			RoleID:   1,
		}
		err = db.GetDB().Create(user).Error
		if err != nil {
			log.Panicf("create owner failed:%v/n", err)
		}
	}
}
