package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/dtos"
	"github.com/9d77v/pdc/graph/model"
	"github.com/9d77v/pdc/models"
	"github.com/9d77v/pdc/utils"
	"golang.org/x/crypto/bcrypt"
)

//UserService ..
type UserService struct {
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
	err = models.Gorm.Create(m).Error
	if err != nil {
		return &model.User{}, err
	}
	return &model.User{UID: models.GetEncodeUID(m.ID)}, err
}

//UpdateUser ..
func (s UserService) UpdateUser(ctx context.Context, input model.NewUpdateUser) (*model.User, error) {
	user := new(models.User)
	varibales := graphql.GetRequestContext(ctx).Variables
	fields := make([]string, 0)
	for k := range varibales["input"].(map[string]interface{}) {
		fields = append(fields, k)
	}
	if err := models.Gorm.Select(utils.ToDBFields(fields)).First(user, "id=?", input.ID).Error; err != nil {
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
	err := models.Gorm.Model(user).Updates(updateMap).Error
	key := fmt.Sprintf("%s:%d", models.PrefixUser, input.ID)
	redisErr := models.RedisClient.Del(ctx, key).Err()
	if redisErr != nil {
		log.Println("redis del failed:", redisErr)
	}
	return &model.User{UID: models.GetEncodeUID(user.ID)}, err
}

//ListUser ..
func (s UserService) ListUser(ctx context.Context, keyword *string,
	page, pageSize *int64, ids []int64, sorts []*model.Sort,
	scheme string) (int64, []*model.User, error) {
	result := make([]*model.User, 0)
	data := make([]*models.User, 0)
	offset, limit := GetPageInfo(page, pageSize)
	fieldMap, _ := utils.GetFieldData(ctx, "")
	var err error
	builder := models.Gorm
	if keyword != nil && ptrs.String(keyword) != "" {
		builder = builder.Where("name like ?", "%"+ptrs.String(keyword)+"%")
	}
	var total int64
	if fieldMap["totalCount"] {
		if limit == -1 {
			total = int64(len(data))
		} else {
			err = builder.Model(&models.User{}).Count(&total).Error
			if err != nil {
				return 0, result, err
			}
		}
	}
	if fieldMap["edges"] {
		_, edgeFields := utils.GetFieldData(ctx, "edges.")
		builder = builder.Select(utils.ToDBFields(edgeFields, "__typename"))
		if len(ids) > 0 {
			builder = builder.Where("id in (?)", ids)
		}
		if limit > 0 {
			builder = builder.Offset(offset).Limit(limit)
		}
		for _, v := range sorts {
			sort := " DESC"
			if v.IsAsc {
				sort = " ASC"
			}
			builder = builder.Order(v.Field + sort)
		}
		err = builder.Find(&data).Error
		if err != nil {
			return 0, result, err
		}
	}
	for _, m := range data {
		r := dtos.ToUserDto(m, scheme)
		result = append(result, r)
	}
	return total, result, nil
}

func (s UserService) checkUserName(ctx context.Context, name string) (bool, error) {
	var total int64
	if err := models.Gorm.Model(&models.User{}).Select("name").Where("name=?", name).Count(&total).Error; err != nil {
		return false, err
	}
	return total == 0, nil
}

//Login ..
func (s UserService) Login(ctx context.Context, username string, password string) (*model.LoginResponse, error) {
	res := new(model.LoginResponse)
	user := new(models.User)
	if err := models.Gorm.Select("id,name,password").Where("name=?", username).First(user).Error; err != nil {
		return res, errors.New("用户名或密码错误")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return res, errors.New("用户名或密码错误")
	}
	uid := models.GetEncodeUID(user.ID)
	res.AccessToken = utils.JWT([]byte(models.JWTtAccessSecret), uid, accessExpire, models.JWTIssuer)
	res.RefreshToken = utils.JWT([]byte(models.JWTRefreshSecret), uid, refreshExpire, models.JWTIssuer)
	return res, nil
}

//RefreshToken ..
func (s UserService) RefreshToken(ctx context.Context, refreshToken string) (*model.LoginResponse, error) {
	res := new(model.LoginResponse)
	token, err := utils.ParseJWT([]byte(models.JWTRefreshSecret), refreshToken)
	if err != nil {
		return res, err
	}
	data, _ := token.Claims.(*utils.MyCustomClaims)
	res.AccessToken = utils.JWT([]byte(models.JWTtAccessSecret), data.UID, accessExpire, models.JWTIssuer)
	res.RefreshToken = utils.JWT([]byte(models.JWTRefreshSecret), data.UID, refreshExpire, models.JWTIssuer)
	return res, nil
}

//GetByID ..
func (s UserService) GetByID(ctx context.Context, uid int64) (*models.User, error) {
	user := new(models.User)
	key := fmt.Sprintf("%s:%d", models.PrefixUser, uid)
	err := models.RedisClient.Get(ctx, key).Scan(user)
	if err != nil {
		if err := user.GetByID(uid); err != nil {
			return user, err
		}
		err = models.RedisClient.Set(ctx, key, user, time.Hour).Err()
		if err != nil {
			log.Printf("set redis key %s error：%v", key, err)
		}
	}
	return user, nil
}

//UpdateProfile ..
func (s UserService) UpdateProfile(ctx context.Context, input model.NewUpdateProfile, uid uint) (*model.User, error) {
	user := new(models.User)
	varibales := graphql.GetRequestContext(ctx).Variables
	fields := make([]string, 0)
	for k := range varibales["input"].(map[string]interface{}) {
		fields = append(fields, k)
	}
	fields = append(fields, "id")
	if err := models.Gorm.Select(utils.ToDBFields(fields)).First(user, "id=?", uid).Error; err != nil {
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
	err := models.Gorm.Model(user).Updates(updateMap).Error
	key := fmt.Sprintf("%s:%d", models.PrefixUser, uid)
	redisErr := models.RedisClient.Del(ctx, key).Err()
	if redisErr != nil {
		log.Println("redis del failed:", redisErr)
	}
	return &model.User{UID: models.GetEncodeUID(uid)}, err
}

//UpdatePassword ..
func (s UserService) UpdatePassword(ctx context.Context, oldPassword string, newPassword string, uid uint) (*model.User, error) {
	if len(newPassword) < 10 || len(newPassword) > 32 {
		return nil, errors.New("新旧密码长度需要在10-32之间")
	}
	if oldPassword == newPassword {
		return nil, errors.New("新旧密码长度不能相同")
	}
	user := new(models.User)
	if err := models.Gorm.Select("id,password").First(user, "id=?", uid).Error; err != nil {
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
	err = models.Gorm.Model(user).Updates(updateMap).Error
	if err != nil {
		log.Println("update password failed")
		return nil, err
	}
	return &model.User{UID: models.GetEncodeUID(uid)}, err
}
