package services

import (
	"context"
	"time"

	"git.9d77v.me/9d77v/pdc/dtos"
	"git.9d77v.me/9d77v/pdc/graph/model"
	"git.9d77v.me/9d77v/pdc/models"
	"git.9d77v.me/9d77v/pdc/utils"
	"github.com/99designs/gqlgen/graphql"
	"github.com/9d77v/go-lib/ptrs"
)

//ThingService ..
type ThingService struct {
}

//CreateThing ..
func (s ThingService) CreateThing(input model.NewThing) (*model.Thing, error) {
	m := &models.Thing{
		UID:              1,
		Name:             input.Name,
		Desc:             ptrs.String(input.Desc),
		Num:              input.Num,
		BrandName:        ptrs.String(input.BrandName),
		Pics:             input.Pics,
		UnitPrice:        input.UnitPrice,
		Unit:             ptrs.String(input.Unit),
		Category:         input.Category,
		Location:         input.Location,
		PurchaseDate:     time.Unix(input.PurchaseDate, 0),
		Status:           int8(input.Status),
		PurchasePlatform: ptrs.String(input.PurchasePlatform),
		RefOrderID:       ptrs.String(input.RefOrderID),
		RubbishCategory:  input.RubbishCategory,
	}
	err := models.Gorm.Create(m).Error
	if err != nil {
		return &model.Thing{}, err
	}
	return &model.Thing{ID: int64(m.ID)}, err
}

//UpdateThing ..
func (s ThingService) UpdateThing(ctx context.Context, input *model.NewUpdateThing) (*model.Thing, error) {
	Thing := new(models.Thing)
	varibales := graphql.GetRequestContext(ctx).Variables
	fields := make([]string, 0)
	for k := range varibales["input"].(map[string]interface{}) {
		fields = append(fields, k)
	}
	if err := models.Gorm.Select(utils.ToDBFields(fields)).First(Thing, "id=?", input.ID).Error; err != nil {
		return nil, err
	}
	err := models.Gorm.Model(Thing).Update(map[string]interface{}{
		"name":              ptrs.String(input.Name),
		"desc":              ptrs.String(input.Desc),
		"num":               ptrs.Float64(input.Num),
		"brand_name":        ptrs.String(input.BrandName),
		"pics":              input.Pics,
		"unit_price":        ptrs.Float64(input.UnitPrice),
		"unit":              ptrs.String(input.Unit),
		"category":          ptrs.String(input.Category),
		"location":          ptrs.String(input.Location),
		"status":            ptrs.Int64(input.Status),
		"purchase_date":     time.Unix(ptrs.Int64(input.PurchaseDate), 0),
		"purchase_platform": ptrs.String(input.PurchasePlatform),
		"ref_order_id":      ptrs.String(input.RefOrderID),
		"rubbish_category":  input.RubbishCategory,
	}).Error
	return &model.Thing{ID: int64(Thing.ID)}, err
}

//ListThing ..
func (s ThingService) ListThing(ctx context.Context, page, pageSize *int64, ids []int64, sorts []*model.Sort) (int64, []*model.Thing, error) {
	offset := ptrs.Int64(page)
	limit := ptrs.Int64(pageSize)
	if offset < 1 {
		offset = 1
	}
	if limit == 0 {
		limit = 10
	}
	if limit < 0 {
		limit = -1
	}
	if limit > 50 {
		limit = 50
	}
	offset = (offset - 1) * limit
	result := make([]*model.Thing, 0)
	data := make([]*models.Thing, 0)

	filedMap, _ := utils.GetFieldData(ctx, "")
	var err error
	if filedMap["edges"] {
		_, edgeFields := utils.GetFieldData(ctx, "edges.")
		builder := models.Gorm.Select(utils.ToDBFields(edgeFields, "__typename"))
		if len(ids) > 0 {
			builder = builder.Where("id in (?)", ids)
		}
		if limit > 0 {
			builder = builder.Offset(offset).Limit(limit)
		}
		for _, v := range sorts {
			if v.IsAsc {
				builder = builder.Order(v.Field + " ASC")
			} else {
				builder = builder.Order(v.Field + " DESC")
			}
		}
		err = builder.Find(&data).Error
		if err != nil {
			return 0, result, err
		}
	}
	var total int64
	if filedMap["totalCount"] {
		if limit == -1 {
			total = int64(len(data))
		} else {
			err = models.Gorm.Model(&models.Thing{}).Count(&total).Error
			if err != nil {
				return 0, result, err
			}
		}
	}
	for _, m := range data {
		r := dtos.ToThingDto(m)
		result = append(result, r)
	}
	return total, result, nil
}
