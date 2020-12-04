package services

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/thing-service/models"
	"github.com/9d77v/pdc/internal/utils"
)

//ThingService ..
type ThingService struct {
}

//CreateThing ..
func (s ThingService) CreateThing(ctx context.Context, input model.NewThing, uid int64) (*model.Thing, error) {
	m := &models.Thing{
		UID:                 uid,
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
	err := db.GetDB().Create(m).Error
	if err != nil {
		return &model.Thing{}, err
	}
	return &model.Thing{ID: int64(m.ID)}, err
}

//UpdateThing ..
func (s ThingService) UpdateThing(ctx context.Context, input model.NewUpdateThing, uid uint) (*model.Thing, error) {
	thing := models.NewThing()
	varibales := graphql.GetRequestContext(ctx).Variables
	fields := make([]string, 0)
	for k := range varibales["input"].(map[string]interface{}) {
		fields = append(fields, k)
	}
	if err := thing.GetByID(uint(input.ID), uid, fields); err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"name":                 ptrs.String(input.Name),
		"num":                  ptrs.Float64(input.Num),
		"brand_name":           ptrs.String(input.BrandName),
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
	}
	if len(input.Pics) > 0 {
		updateMap["pics"] = input.Pics
	}
	err := db.GetDB().Model(thing).Updates(updateMap).Error
	return &model.Thing{ID: int64(thing.ID)}, err
}

//ListThing ..
func (s ThingService) ListThing(ctx context.Context, keyword *string,
	page, pageSize *int64, ids []int64, sorts []*model.Sort,
	uid uint, scheme string) (int64, []*model.Thing, error) {
	result := make([]*model.Thing, 0)
	data := make([]*models.Thing, 0)
	offset, limit := utils.GetPageInfo(page, pageSize)
	fieldMap, _ := utils.GetFieldData(ctx, "")
	var err error
	thing := models.NewThing()
	thing.IDQuery(uid, "uid").FuzzyQuery(keyword, "name")
	var total int64
	if fieldMap["totalCount"] {
		if limit == -1 {
			total = int64(len(data))
		} else {
			total, err = thing.Count(thing)
			if err != nil {
				return 0, result, err
			}
		}
	}
	if fieldMap["edges"] {
		_, edgeFields := utils.GetFieldData(ctx, "edges.")
		err = thing.Select(edgeFields).
			IDArrayQuery(thing.ToUintIDs(ids)).
			Pagination(offset, limit).
			Sort(sorts).
			Find(&data)
		if err != nil {
			return 0, result, err
		}
	}
	return total, toThingsDtos(data, scheme), nil
}

//ThingSeries 获取物品数据
func (s ThingService) ThingSeries(ctx context.Context, dimension, index string, start *int64, end *int64, status []int64, uid int64) ([]*model.SerieData, error) {
	data := make([]*model.SerieData, 0)
	builder := db.GetDB().Model(&models.Thing{})
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
	builder = builder.Where("uid=?", uid)
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
	builder := db.GetDB().Model(&models.Thing{})
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
	builder = builder.Where("uid=?", uid)
	builder = builder.Group("x1," + dimension).Order("x1," + dimension)
	data := make([]*models.PieLineSerie, 0)
	err := builder.Scan(&data).Error
	return toPieLineSerieData(data), err
}
