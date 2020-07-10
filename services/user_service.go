package services

import (
	"context"
	"errors"
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
	return &model.User{ID: int64(m.ID)}, err
}

//UpdateUser ..
func (s UserService) UpdateUser(ctx context.Context, input model.NewUpdateUser) (*model.User, error) {
	User := new(models.User)
	varibales := graphql.GetRequestContext(ctx).Variables
	fields := make([]string, 0)
	for k := range varibales["input"].(map[string]interface{}) {
		fields = append(fields, k)
	}
	if err := models.Gorm.Select(utils.ToDBFields(fields)).First(User, "id=?", input.ID).Error; err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"avatar":     ptrs.String(input.Avatar),
		"role_id":    ptrs.Int64(input.RoleID),
		"gender":     ptrs.Int64(input.Gender),
		"color":      ptrs.String(input.Color),
		"birth_date": time.Unix(ptrs.Int64(input.BirthDate), 0),
		"ip":         ptrs.String(input.IP),
	}
	if input.Password != nil {
		bytes, err := bcrypt.GenerateFromPassword([]byte(ptrs.String(input.Password)), 12)
		if err != nil {
			log.Printf("generate password failed:%v/n", err)
			return nil, err
		}
		updateMap["password"] = string(bytes)
	}
	err := models.Gorm.Model(User).Update(updateMap).Error
	return &model.User{ID: int64(User.ID)}, err
}

//ListUser ..
func (s UserService) ListUser(ctx context.Context, page, pageSize *int64, ids []int64, sorts []*model.Sort) (int64, []*model.User, error) {
	result := make([]*model.User, 0)
	data := make([]*models.User, 0)
	offset, limit := GetPageInfo(page, pageSize)
	filedMap, _ := utils.GetFieldData(ctx, "")
	var err error
	builder := models.Gorm
	if filedMap["edges"] {
		_, edgeFields := utils.GetFieldData(ctx, "edges.")
		builder = builder.Select(utils.ToDBFields(edgeFields, "__typename"))
		if len(ids) > 0 {
			builder = builder.Where("id in (?)", ids)
		}
		subBuilder := builder
		if limit > 0 {
			subBuilder = subBuilder.Offset(offset).Limit(limit)
		}
		for _, v := range sorts {
			if v.IsAsc {
				subBuilder = subBuilder.Order(v.Field + " ASC")
			} else {
				subBuilder = subBuilder.Order(v.Field + " DESC")
			}
		}
		err = subBuilder.Find(&data).Error
		if err != nil {
			return 0, result, err
		}
	}
	var total int64
	if filedMap["totalCount"] {
		if limit == -1 {
			total = int64(len(data))
		} else {
			err = builder.Model(&models.User{}).Count(&total).Error
			if err != nil {
				return 0, result, err
			}
		}
	}
	for _, m := range data {
		r := dtos.ToUserDto(m)
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
		return res, err
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return res, err
	}
	res.AccessToken = utils.JWT([]byte(models.JWTtAccessSecret), int64(user.ID), accessExpire, models.JWTIssuer)
	res.RefreshToken = utils.JWT([]byte(models.JWTRefreshSecret), int64(user.ID), refreshExpire, models.JWTIssuer)
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
func (s UserService) GetByID(uid int64) (*model.User, error) {
	res := new(model.User)
	user := new(models.User)
	if err := user.GetByID(uid); err != nil {
		return res, err
	}
	return res, nil
}
