package services

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/db/clickhouse"
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/internal/db/redis"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/history-service/chmodels"
	"github.com/9d77v/pdc/internal/module/history-service/models"
	redisGo "github.com/go-redis/redis/v8"
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
	dataMap[redis.PrefixVideoDataUser] = hl.UID
	dataMap[redis.PrefixVideoDataAnime] = hl.SourceID
	dataMap[redis.PrefixVideoDataEpisode] = hl.SubSourceID
	ctx := context.Background()
	var err error
	for k, v := range dataMap {
		key := fmt.Sprintf("%s:%s", k, today)
		err = redis.GetClient().HSet(ctx, key, v, hl.ServerTs.Unix()).Err()
		if err != nil {
			log.Println("redis set failed:", err)
		}
		err = redis.GetClient().Expire(ctx, key, 72*time.Hour).Err()
		if err != nil {
			log.Println("redis expire key failed:", err)
		}
	}
	durationKey := fmt.Sprintf("%s:%s", redis.PrefixVideoDataDuration, today)
	err = redis.GetClient().IncrByFloat(ctx, durationKey, hl.Duration).Err()
	if err != nil {
		log.Println("redis set failed:", err)
	}
	err = redis.GetClient().Expire(ctx, durationKey, 72*time.Hour).Err()
	if err != nil {
		log.Println("redis expire key failed:", err)
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
	now := time.Now()
	today := now.Format("2006-01-02")
	yesterday := now.Add(-24 * time.Hour).Format("2006-01-02")
	theDayBeforeYesterday := now.Add(-48 * time.Hour).Format("2006-01-02")
	days := []string{today, yesterday, theDayBeforeYesterday}
	keyPrefixs := []string{redis.PrefixVideoDataUser, redis.PrefixVideoDataAnime,
		redis.PrefixVideoDataEpisode}
	pipe := redis.GetClient().Pipeline()
	m, n := len(keyPrefixs), len(days)
	intCmds := make([]*redisGo.IntCmd, 0, m*n)
	for _, keyPrefix := range keyPrefixs {
		for _, day := range days {
			cmd := pipe.HLen(ctx, fmt.Sprintf("%s:%s", keyPrefix, day))
			intCmds = append(intCmds, cmd)
		}
	}
	stringCmds := make([]*redisGo.StringCmd, 0, n)
	for _, day := range days {
		cmd := pipe.Get(ctx, fmt.Sprintf("%s:%s", redis.PrefixVideoDataDuration, day))
		stringCmds = append(stringCmds, cmd)
	}
	pipe.Exec(ctx)
	data := make([][]float64, m+1)
	for i := 0; i < m; i++ {
		data[i] = make([]float64, 0, n)
		for j := 0; j < n; j++ {
			data[i] = append(data[i], float64(intCmds[i*n+j].Val()))
		}
	}
	data[m] = make([]float64, 0, n)
	for i := 0; i < n; i++ {
		v, _ := strconv.ParseFloat(stringCmds[i].Val(), 64)
		data[m] = append(data[m], v)
	}
	result.Data = data
	return result, nil
}
