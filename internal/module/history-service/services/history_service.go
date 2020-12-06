package services

import (
	"context"
	"errors"
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
	return toHistoryDto(history), nil
}

//ListHistory ..
func (s HistoryService) ListHistory(ctx context.Context,
	sourceType *int64, searchParam model.SearchParam, uid uint, scheme string) (int64, []*model.History, error) {
	history := models.NewHistory()
	history.Where("uid=? and source_type=?", uid, ptrs.Int64(sourceType))
	replaceFunc := func(edgeField base.GraphQLField) error {
		tableHistory := history.TableName()
		switch ptrs.Int64(sourceType) {
		case 1:
			history.
				Select([]string{"uid", "source_type", "source_id", "sub_source_id", "current_time",
					"remaining_time", "platform",
					"cast(EXTRACT(epoch FROM CAST( " +
						tableHistory + ".updated_at AS TIMESTAMP)) as bigint) updated_at",
					"b.title", "b.cover", "c.num", "c.title sub_title"}).
				LeftJoin("pdc_video b ON " + tableHistory + ".source_id=b.id").
				LeftJoin("pdc_episode c on " + tableHistory + ".sub_source_id=c.id")
		default:
			return errors.New("sourceType not exist")
		}
		history.Order("updated_at desc")
		return nil
	}
	data := make([]*model.History, 0)
	total, err := s.GetConnection(ctx, history, searchParam, &data, replaceFunc)
	for _, v := range data {
		v.Cover = oss.GetOSSPrefixByScheme(scheme) + v.Cover
	}
	return total, data, err
}
