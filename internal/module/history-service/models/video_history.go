package models

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/9d77v/go-lib/ptrs"

	"github.com/9d77v/pdc/internal/db/redis"
	"github.com/9d77v/pdc/internal/module/base"
	video "github.com/9d77v/pdc/internal/module/video-service/models"
	redisGo "github.com/go-redis/redis/v8"
)

//VideoHistory ...
type VideoHistory struct {
}

//GetSourceID ..
func (h VideoHistory) GetSourceID(subSourceID *int64) uint {
	return video.NewEpisode().GetVideoIDByID(uint(ptrs.Int64(subSourceID)))
}

//JoinSource ..
func (h VideoHistory) JoinSource(r base.Repository, tableName string, sourceID uint) error {
	if sourceID > 0 {
		r.Where("source_id=?", sourceID)
	}
	r.Select([]string{"uid", "source_type", "source_id", "sub_source_id", "current_time",
		"remaining_time", "platform",
		"cast(EXTRACT(epoch FROM CAST( " +
			tableName + ".updated_at AS TIMESTAMP)) as bigint) updated_at",
		"b.title", "b.cover", "c.num", "c.title sub_title"}).
		LeftJoin("pdc_video b ON " + tableName + ".source_id=b.id").
		LeftJoin("pdc_episode c on " + tableName + ".sub_source_id=c.id")
	return nil
}

//GetStatistic 获取统计数据
func (h VideoHistory) GetStatistic(ctx context.Context, uid uint) [][]float64 {
	now := time.Now()
	today := now.Format("2006-01-02")
	yesterday := now.Add(-24 * time.Hour).Format("2006-01-02")
	theDayBeforeYesterday := now.Add(-48 * time.Hour).Format("2006-01-02")
	days := []string{today, yesterday, theDayBeforeYesterday}
	keyPrefixs := []string{redis.PrefixVideoDataAnime, redis.PrefixVideoDataEpisode}
	if uid == 0 {
		keyPrefixs = append([]string{redis.PrefixVideoDataUser}, keyPrefixs...)
	}
	pipe := redis.GetClient().Pipeline()
	m, n := len(keyPrefixs), len(days)
	intCmds := make([]*redisGo.IntCmd, 0, m*n)
	for _, keyPrefix := range keyPrefixs {
		for _, day := range days {
			key := fmt.Sprintf("%s:%s", keyPrefix, day)
			if uid != 0 {
				key += fmt.Sprintf(":%d", uid)
			}
			cmd := pipe.HLen(ctx, key)
			intCmds = append(intCmds, cmd)
		}
	}
	stringCmds := make([]*redisGo.StringCmd, 0, n)
	for _, day := range days {
		key := fmt.Sprintf("%s:%s", redis.PrefixVideoDataDuration, day)
		if uid != 0 {
			key += fmt.Sprintf(":%d", uid)
		}
		cmd := pipe.Get(ctx, key)
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
	return data
}
