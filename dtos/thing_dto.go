package dtos

import (
	"git.9d77v.me/9d77v/pdc/graph/model"
	"git.9d77v.me/9d77v/pdc/models"
)

//ToThingDto ...
func ToThingDto(m *models.Thing) *model.Thing {
	return &model.Thing{
		ID:               int64(m.ID),
		UID:              m.UID,
		Name:             m.Name,
		Num:              m.Num,
		BrandName:        &m.BrandName,
		Pics:             m.Pics,
		UnitPrice:        m.UnitPrice,
		Unit:             &m.Unit,
		Specifications:   &m.Specifications,
		Category:         m.Category,
		Location:         m.Location,
		PurchaseDate:     m.PurchaseDate.Unix(),
		Status:           int64(m.Status),
		PurchasePlatform: &m.PurchasePlatform,
		RefOrderID:       m.RefOrderID,
		RubbishCategory:  m.RubbishCategory,
		CreatedAt:        m.CreatedAt.Unix(),
		UpdatedAt:        m.UpdatedAt.Unix(),
	}
}
