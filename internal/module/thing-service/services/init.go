package services

import (
	"github.com/9d77v/pdc/internal/db"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/thing-service/models"
)

func toThingDto(m *models.Thing, scheme string) *model.Thing {
	newPics := make([]string, 0)
	for _, v := range m.Pics {
		newPics = append(newPics, db.GetOSSPrefix(scheme)+v)
	}
	return &model.Thing{
		ID:                  int64(m.ID),
		UID:                 m.UID,
		Name:                m.Name,
		Num:                 m.Num,
		BrandName:           &m.BrandName,
		Pics:                newPics,
		UnitPrice:           m.UnitPrice,
		Unit:                &m.Unit,
		Specifications:      &m.Specifications,
		Category:            int64(m.Category),
		ConsumerExpenditure: m.ConsumerExpenditure,
		Location:            m.Location,
		PurchaseDate:        m.PurchaseDate.Unix(),
		Status:              int64(m.Status),
		PurchasePlatform:    &m.PurchasePlatform,
		RefOrderID:          m.RefOrderID,
		RubbishCategory:     m.RubbishCategory,
		CreatedAt:           m.CreatedAt.Unix(),
		UpdatedAt:           m.UpdatedAt.Unix(),
	}
}

func toPieLineSerieData(dbData []*models.PieLineSerie) *model.PieLineSerieData {
	data := &model.PieLineSerieData{
		X1: make([]string, 0),
		X2: make([]string, 0),
		Y:  make([]float64, 0),
	}
	for _, v := range dbData {
		data.X1 = append(data.X1, v.X1)
		data.X2 = append(data.X2, v.X2)
		data.Y = append(data.Y, v.Y)
	}
	return data
}
