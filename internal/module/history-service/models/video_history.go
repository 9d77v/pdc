package models

import (
	"github.com/9d77v/go-lib/ptrs"

	"github.com/9d77v/pdc/internal/module/base"
	video "github.com/9d77v/pdc/internal/module/video-service/models"
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
