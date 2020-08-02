package services

import (
	"context"
	"log"
	"time"

	"github.com/9d77v/pdc/dtos"
	"github.com/9d77v/pdc/graph/model"
	"github.com/9d77v/pdc/models"
)

//HistoryService ..
type HistoryService struct{}

//RecordHistory ..
func (s HistoryService) RecordHistory(ctx context.Context, input model.NewHistoryInput, uid uint) (*model.History, error) {
	m := &models.History{
		UID:           uid,
		SourceType:    uint8(input.SourceType),
		SourceID:      uint(input.SourceID),
		SubSourceID:   uint(input.SubSourceID),
		CurrentTime:   input.CurrentTime,
		RemainingTime: input.RemainingTime,
		UpdatedAt:     time.Now(),
	}
	err := models.Gorm.Save(m).Error
	if err != nil {
		return &model.History{}, err
	}
	return &model.History{}, err
}

//GetHistory ..
func (s HistoryService) GetHistory(ctx context.Context, sourceType int64, sourceID int64, uid uint) (*model.History, error) {
	history := new(models.History)
	err := models.Gorm.Where("uid=? and source_type=? and source_id=?", uid, sourceType, sourceID).Take(history).Error
	if err != nil {
		log.Println("get history error", err)
		return nil, nil
	}
	return dtos.ToHistoryDto(history), nil
}
