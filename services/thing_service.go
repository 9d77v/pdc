package services

import (
	"context"
	"fmt"
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
		UID:                 1,
		Name:                input.Name,
		Num:                 input.Num,
		BrandName:           ptrs.String(input.BrandName),
		Pics:                input.Pics,
		UnitPrice:           input.UnitPrice,
		Unit:                ptrs.String(input.Unit),
		Specifications:      ptrs.String(input.Specifications),
		Category:            int8(input.Category),
		ConsumerExpenditure: input.ConsumerExpenditure,
		Location:            ptrs.String(input.Location),
		PurchaseDate:        time.Unix(input.PurchaseDate, 0),
		Status:              int8(input.Status),
		PurchasePlatform:    ptrs.String(input.PurchasePlatform),
		RefOrderID:          ptrs.String(input.RefOrderID),
		RubbishCategory:     input.RubbishCategory,
	}
	err := models.Gorm.Create(m).Error
	if err != nil {
		return &model.Thing{}, err
	}
	return &model.Thing{ID: int64(m.ID)}, err
}

//UpdateThing ..
func (s ThingService) UpdateThing(ctx context.Context, input model.NewUpdateThing) (*model.Thing, error) {
	thing := new(models.Thing)
	varibales := graphql.GetRequestContext(ctx).Variables
	fields := make([]string, 0)
	for k := range varibales["input"].(map[string]interface{}) {
		fields = append(fields, k)
	}
	if err := models.Gorm.Select(utils.ToDBFields(fields)).First(thing, "id=?", input.ID).Error; err != nil {
		return nil, err
	}
	err := models.Gorm.Model(thing).Update(map[string]interface{}{
		"name":                 ptrs.String(input.Name),
		"num":                  ptrs.Float64(input.Num),
		"brand_name":           ptrs.String(input.BrandName),
		"pics":                 input.Pics,
		"unit_price":           ptrs.Float64(input.UnitPrice),
		"unit":                 ptrs.String(input.Unit),
		"specifications":       ptrs.String(input.Specifications),
		"category":             ptrs.Int64(input.Category),
		"consumer_expenditure": ptrs.String(input.ConsumerExpenditure),
		"location":             ptrs.String(input.Location),
		"status":               ptrs.Int64(input.Status),
		"purchase_date":        time.Unix(ptrs.Int64(input.PurchaseDate), 0),
		"purchase_platform":    ptrs.String(input.PurchasePlatform),
		"ref_order_id":         ptrs.String(input.RefOrderID),
		"rubbish_category":     input.RubbishCategory,
	}).Error
	return &model.Thing{ID: int64(thing.ID)}, err
}

//ListThing ..
func (s ThingService) ListThing(ctx context.Context, page, pageSize *int64, ids []int64, sorts []*model.Sort) (int64, []*model.Thing, error) {
	result := make([]*model.Thing, 0)
	data := make([]*models.Thing, 0)
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
			err = builder.Model(&models.Thing{}).Count(&total).Error
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

//ThingSeries 获取物品数据
func (s ThingService) ThingSeries(ctx context.Context, dimension, index string, start *int64, end *int64, status []int64, uid int64) ([]*model.SerieData, error) {
	data := make([]*model.SerieData, 0)
	builder := models.Gorm.Model(&models.Thing{})
	if index == "num" {
		builder = builder.Select(fmt.Sprintf("%s as name, sum(num) as value ", dimension))
	} else if index == "price" {
		builder = builder.Select(fmt.Sprintf("%s as name, sum(num*unit_price)::money::numeric::float8 as value ", dimension))
	}
	startTime := ptrs.Int64(start)
	endTime := ptrs.Int64(end)
	if endTime == 0 {
		endTime = time.Now().Unix()
	}
	builder = builder.Where("purchase_date>=? and purchase_date<=?", time.Unix(startTime, 0), time.Unix(endTime, 0))
	if len(status) > 0 {
		builder = builder.Where("status in (?)", status)
	}
	builder = builder.Where("uid=?", 1)
	builder = builder.Group(dimension).Order("name")
	err := builder.Scan(&data).Error
	return data, err
}

//ThingAnalyze 物品分析
func (s ThingService) ThingAnalyze(ctx context.Context, dimension, index string, start *int64, group string, uid int64) (*model.PieLineSerieData, error) {
	startInt := ptrs.Int64(start)
	startTime := time.Unix(startInt, 0)
	var endTime time.Time
	var format string
	switch group {
	case "month":
		year, month, _ := startTime.Date()
		startTime = time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
		endTime = startTime.AddDate(0, 1, 0)
		format = "YYYY-MM-DD"
	case "year":
		year, _, _ := startTime.Date()
		startTime = time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
		endTime = startTime.AddDate(1, 0, 0)
		format = "YYYY-MM"
	default:
		startTime = time.Date(1991, 4, 18, 0, 0, 0, 0, time.Local)
		endTime = time.Now()
		format = "YYYY"
	}
	builder := models.Gorm.Model(&models.Thing{})
	if index == "num" {
		builder = builder.Select(
			fmt.Sprintf("to_char(purchase_date,'%s') as x1, %s as x2, sum(num) as y ",
				format, dimension))
	} else if index == "price" {
		builder = builder.Select(
			fmt.Sprintf("to_char(purchase_date,'%s') as x1,%s as x2, sum(num*unit_price)::money::numeric::float8 as y ",
				format, dimension))
	}
	builder = builder.Where("purchase_date>=? and purchase_date<?", startTime, endTime)
	builder = builder.Where("uid=?", 1)
	builder = builder.Group("x1," + dimension).Order("x1," + dimension)
	data := make([]*models.PieLineSerie, 0)
	err := builder.Scan(&data).Error
	return dtos.ToPieLineSerieData(data), err
}
