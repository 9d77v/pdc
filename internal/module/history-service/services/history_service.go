package services

import (
	"context"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/9d77v/go-pkg/cache/redis"
	"github.com/9d77v/go-pkg/ptrs"
	"github.com/9d77v/pdc/internal/consts"
	"github.com/9d77v/pdc/internal/db/clickhouse"

	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/history-service/chmodels"
	"github.com/9d77v/pdc/internal/module/history-service/models"
)

//HistoryService ..
type HistoryService struct {
	base.Service
}

//RecordHistory ..
func (s HistoryService) RecordHistory(ctx context.Context,
	input model.NewHistoryInput, uid uint) (*model.History, error) {
	now := time.Now()
	h := &models.History{
		UID:           uid,
		SourceType:    uint8(input.SourceType),
		SourceID:      uint(input.SourceID),
		SubSourceID:   uint(input.SubSourceID),
		Platform:      input.Platform,
		CurrentTime:   input.CurrentTime,
		RemainingTime: input.RemainingTime,
		UpdatedAt:     now,
	}
	h.SetDB(db.GetDB())
	err := h.Save(h)
	if err != nil {
		return &model.History{}, err
	}
	clientTs := math.Trunc(input.ClientTs)
	hl := &chmodels.HistoryLog{
		UID:           uint64(uid),
		SourceType:    uint8(input.SourceType),
		SourceID:      uint64(input.SourceID),
		SubSourceID:   uint64(input.SubSourceID),
		ServerTs:      now,
		ClientTs:      time.Unix(int64(clientTs), int64((input.ClientTs-clientTs)*1e6)),
		Platform:      input.Platform,
		CurrentTime:   input.CurrentTime,
		RemainingTime: input.RemainingTime,
		Duration:      input.Duration,
	}
	hl.SetDB(clickhouse.GetDB())
	err = hl.Save(hl)
	if err != nil {
		log.Println("save history log failed,err:", err)
	}
	go s.saveToCache(hl)
	return &model.History{
		SubSourceID: input.SubSourceID,
	}, err
}

func (s HistoryService) saveToCache(hl *chmodels.HistoryLog) {
	today := hl.ClientTs.Format("2006-01-02")
	dataMap := make(map[string]interface{})
	dataMap[consts.PrefixVideoDataUser] = hl.UID
	dataMap[consts.PrefixVideoDataAnime] = hl.SourceID
	dataMap[consts.PrefixVideoDataEpisode] = hl.SubSourceID
	ctx := context.Background()
	pipe := redis.GetClient().Pipeline()
	for k, v := range dataMap {
		keys := []string{fmt.Sprintf("%s:%s", k, today), fmt.Sprintf("%s:%s:%d", k, today, hl.UID)}
		for _, key := range keys {
			pipe.HSet(ctx, key, v, hl.ServerTs.Unix())
			pipe.Expire(ctx, key, 72*time.Hour)
		}
	}
	durationKeys := []string{fmt.Sprintf("%s:%s", consts.PrefixVideoDataDuration, today),
		fmt.Sprintf("%s:%s:%d", consts.PrefixVideoDataDuration, today, hl.UID)}
	for _, durationKey := range durationKeys {
		pipe.IncrByFloat(ctx, durationKey, hl.Duration)
		pipe.Expire(ctx, durationKey, 72*time.Hour)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		log.Println("redis set cache key failed:", err)
	}
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
	sourceType *int64, searchParam *base.SearchParam, subSourceID *int64, uid uint, scheme string) (int64, []*model.History, error) {
	var history base.Repository = models.NewHistory()
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
		history.Order("updated_at desc")
		return historyer.JoinSource(history, models.NewHistory().TableName(), sourceID)
	}
	data := make([]*model.History, 0)
	total, err := s.GetConnection(ctx, history, searchParam, &data, replaceFunc)
	for _, v := range data {
		v.Cover = oss.GetOSSPrefixByScheme(scheme) + v.Cover
	}
	return total, data, err
}

//HistoryStatistic ..
func (s HistoryService) HistoryStatistic(ctx context.Context, sourceType *int64) (*model.HistoryStatistic, error) {
	result := new(model.HistoryStatistic)
	historyer := models.CreateHistory(sourceType)
	result.Data = historyer.GetStatistic(ctx, 0)
	return result, nil
}

//AppHistoryStatistic ..
func (s HistoryService) AppHistoryStatistic(ctx context.Context, sourceType *int64,
	uid uint) (*model.HistoryStatistic, error) {
	result := new(model.HistoryStatistic)
	historyer := models.CreateHistory(sourceType)
	result.Data = historyer.GetStatistic(ctx, uid)
	return result, nil
}
