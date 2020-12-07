package services

import (
	"context"
	"log"
	"time"

	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/history-service/models"
)

//HistoryService ..
type HistoryService struct {
	base.Service
}

//RecordHistory ..
func (s HistoryService) RecordHistory(ctx context.Context,
	input model.NewHistoryInput, uid uint) (*model.History, error) {
	tx := db.GetDB().Begin()
	h := &models.History{
		UID:           uid,
		SourceType:    uint8(input.SourceType),
		SourceID:      uint(input.SourceID),
		SubSourceID:   uint(input.SubSourceID),
		Platform:      input.Platform,
		CurrentTime:   input.CurrentTime,
		RemainingTime: input.RemainingTime,
		UpdatedAt:     time.Now(),
	}
	err := tx.Save(h).Error
	if err != nil {
		tx.Rollback()
		return &model.History{}, err
	}
	hl := &models.HistoryLog{
		UID:           uid,
		SourceType:    uint8(input.SourceType),
		SourceID:      uint(input.SourceID),
		SubSourceID:   uint(input.SubSourceID),
		Platform:      input.Platform,
		CurrentTime:   input.CurrentTime,
		RemainingTime: input.RemainingTime,
	}
	err = tx.Save(hl).Error
	if err != nil {
		tx.Rollback()
		return &model.History{}, err
	}
	return &model.History{
		SubSourceID: input.SubSourceID,
	}, tx.Commit().Error
}

//GetHistory ..
func (s HistoryService) GetHistory(ctx context.Context,
	sourceType int64, sourceID int64, uid uint) (*model.History, error) {
	history := new(models.History)
	err := db.GetDB().Where("uid=? and source_type=? and source_id=?",
		uid, sourceType, sourceID).First(history).Error
	if err != nil {
		log.Println("get history error", err)
		return nil, nil
	}
	return s.getHistory(history), nil
}

//ListHistory ..
func (s HistoryService) ListHistory(ctx context.Context,
	sourceType *int64, searchParam model.SearchParam, subSourceID *int64, uid uint, scheme string) (int64, []*model.History, error) {
	history := models.NewHistory()
	history.Where("uid=? and source_type=?", uid, ptrs.Int64(sourceType))
	var sourceID uint
	historyer := models.CreateHistory(sourceType)
	if ptrs.Int64(subSourceID) > 0 {
		sourceID = historyer.GetSourceID(subSourceID)
		if sourceID == 0 {
			return 0, []*model.History{}, nil
		}
	}
	replaceFunc := func(edgeField base.GraphQLField) error {
		tableHistory := history.TableName()
		history.Order("updated_at desc")
		return historyer.JoinSource(history, tableHistory, sourceID)
	}
	data := make([]*model.History, 0)
	total, err := s.GetConnection(ctx, history, searchParam, &data, replaceFunc)
	for _, v := range data {
		v.Cover = oss.GetOSSPrefixByScheme(scheme) + v.Cover
	}
	return total, data, err
}
