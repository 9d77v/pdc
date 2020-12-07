package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/consts"
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/db/redis"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/user-service/models"
	"github.com/9d77v/pdc/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

//UserService ..
type UserService struct {
	base.Service
}

const (
	accessExpire  = time.Hour
	refreshExpire = 14 * 24 * time.Hour
)

//CreateUser ..
func (s UserService) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	flag, err := s.checkUserName(ctx, input.Name)
	if err != nil {
		return nil, err
	}
	if !flag {
		return nil, errors.New("user is exist")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), 12)
	if err != nil {
		log.Printf("generate password failed:%v/n", err)
		return nil, err
	}
	m := &models.User{
		Name:      input.Name,
		Avatar:    ptrs.String(input.Avatar),
		Password:  string(bytes),
		RoleID:    int(input.RoleID),
		Gender:    int(input.Gender),
		BirthDate: time.Unix(ptrs.Int64(input.BirthDate), 0),
		IP:        ptrs.String(input.IP),
	}
	err = db.GetDB().Create(m).Error
	if err != nil {
		return &model.User{}, err
	}
	return &model.User{UID: consts.GetEncodeUID(m.ID)}, err
}

//UpdateUser ..
func (s UserService) UpdateUser(ctx context.Context, input model.NewUpdateUser) (*model.User, error) {
	user := models.NewUser()
	fields := s.GetInputFields(ctx)
	if err := s.GetByID(user, uint(input.ID), fields); err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"role_id":    ptrs.Int64(input.RoleID),
		"gender":     ptrs.Int64(input.Gender),
		"color":      ptrs.String(input.Color),
		"birth_date": time.Unix(ptrs.Int64(input.BirthDate), 0),
		"ip":         ptrs.String(input.IP),
	}
	if input.Avatar != nil {
		updateMap["avatar"] = ptrs.String(input.Avatar)
	}
	if input.Password != nil {
		bytes, err := bcrypt.GenerateFromPassword([]byte(ptrs.String(input.Password)), 12)
		if err != nil {
			log.Printf("generate password failed:%v/n", err)
			return nil, err
		}
		updateMap["password"] = string(bytes)
	}
	err := db.GetDB().Model(user).Updates(updateMap).Error
	key := fmt.Sprintf("%s:%d", redis.PrefixUser, input.ID)
	redisErr := redis.GetClient().Del(ctx, key).Err()
	if redisErr != nil {
		log.Println("redis del failed:", redisErr)
	}
	return &model.User{UID: consts.GetEncodeUID(user.ID)}, err
}

//ListUser ..
func (s UserService) ListUser(ctx context.Context, searchParam model.SearchParam, scheme string) (int64, []*model.User, error) {
	user := models.NewUser()
	user.FuzzyQuery(searchParam.Keyword, "name")
	data := make([]*models.User, 0)
	total, err := s.GetConnection(ctx, user, searchParam, &data, nil)
	return total, s.getUsers(data, scheme), err
}

func (s UserService) checkUserName(ctx context.Context, name string) (bool, error) {
	var total int64
	if err := db.GetDB().Model(&models.User{}).Select("name").Where("name=?", name).Count(&total).Error; err != nil {
		return false, err
	}
	return total == 0, nil
}

//Login ..
func (s UserService) Login(ctx context.Context, username string, password string) (*model.LoginResponse, error) {
	res := new(model.LoginResponse)
	user := new(models.User)
	if err := db.GetDB().Select("id,name,password").Where("name=?", username).First(user).Error; err != nil {
		return res, errors.New("用户名或密码错误")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return res, errors.New("用户名或密码错误")
	}
	uid := consts.GetEncodeUID(user.ID)
	res.AccessToken = utils.JWT([]byte(consts.JWTtAccessSecret), uid, accessExpire, consts.JWTIssuer)
	res.RefreshToken = utils.JWT([]byte(consts.JWTRefreshSecret), uid, refreshExpire, consts.JWTIssuer)
	return res, nil
}

//RefreshToken ..
func (s UserService) RefreshToken(ctx context.Context, refreshToken string) (*model.LoginResponse, error) {
	res := new(model.LoginResponse)
	token, err := utils.ParseJWT([]byte(consts.JWTRefreshSecret), refreshToken)
	if err != nil {
		return res, err
	}
	data, _ := token.Claims.(*utils.MyCustomClaims)
	res.AccessToken = utils.JWT([]byte(consts.JWTtAccessSecret), data.UID, accessExpire, consts.JWTIssuer)
	res.RefreshToken = utils.JWT([]byte(consts.JWTRefreshSecret), data.UID, refreshExpire, consts.JWTIssuer)
	return res, nil
}

//GetUserByID ..
func (s UserService) GetUserByID(ctx context.Context, uid int64) (*models.User, error) {
	user := models.NewUser()
	key := fmt.Sprintf("%s:%d", redis.PrefixUser, uid)
	err := redis.GetClient().Get(ctx, key).Scan(user)
	if err != nil {
		if err := user.IDQuery(uint(uid)).First(user); err != nil {
			return user, err
		}
		err = redis.GetClient().Set(ctx, key, user, time.Hour).Err()
		if err != nil {
			log.Printf("set redis key %s error：%v", key, err)
		}
	}
	return user, nil
}

//UpdateProfile ..
func (s UserService) UpdateProfile(ctx context.Context, input model.NewUpdateProfile, uid uint) (*model.User, error) {
	user := models.NewUser()
	fields := s.GetInputFields(ctx)
	fields = append(fields, "id")
	if err := s.GetByID(user, uid, fields); err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"gender":     ptrs.Int64(input.Gender),
		"color":      ptrs.String(input.Color),
		"birth_date": time.Unix(ptrs.Int64(input.BirthDate), 0),
		"ip":         ptrs.String(input.IP),
	}
	if input.Avatar != nil {
		updateMap["avatar"] = ptrs.String(input.Avatar)
	}
	err := db.GetDB().Model(user).Updates(updateMap).Error
	key := fmt.Sprintf("%s:%d", redis.PrefixUser, uid)
	redisErr := redis.GetClient().Del(ctx, key).Err()
	if redisErr != nil {
		log.Println("redis del failed:", redisErr)
	}
	return &model.User{UID: consts.GetEncodeUID(uid)}, err
}

//UpdatePassword ..
func (s UserService) UpdatePassword(ctx context.Context,
	oldPassword string, newPassword string, uid uint) (*model.User, error) {
	if len(newPassword) < 10 || len(newPassword) > 32 {
		return nil, errors.New("新旧密码长度需要在10-32之间")
	}
	if oldPassword == newPassword {
		return nil, errors.New("新旧密码长度不能相同")
	}
	user := new(models.User)
	if err := db.GetDB().Select("id,password").First(user, "id=?", uid).Error; err != nil {
		return nil, err
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword))
	if err != nil {
		return nil, errors.New("旧密码错误")
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(newPassword), 12)
	if err != nil {
		log.Printf("generate password failed:%v/n", err)
		return nil, err
	}
	updateMap := map[string]interface{}{
		"password": string(bytes),
	}
	err = db.GetDB().Model(user).Updates(updateMap).Error
	if err != nil {
		log.Println("update password failed")
		return nil, err
	}
	return &model.User{UID: consts.GetEncodeUID(uid)}, err
}

//GetUserInfo ..
func (s UserService) GetUserInfo(user *models.User, scheme string) *model.User {
	return s.getUser(user, scheme)
}
