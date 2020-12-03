package models

import (
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/lib/pq"
)

//VideoIndex 视频索引
type VideoIndex struct {
	ID             uint           `json:"id"`
	Title          string         `json:"title"`
	Desc           string         `json:"desc"`
	PubDate        int64          `json:"pub_date"`
	Cover          string         `json:"cover"`
	TotalNum       int32          `json:"total_num"`
	Tags           pq.StringArray `gorm:"type:varchar(10)[]" json:"tags"`
	EpisodeID      uint           `json:"episode_id"`
	IsShow         bool           `json:"is_show"`
	IsHideOnMobile bool           `json:"is_hide_on_mobile"`
	SeriesID       uint           `json:"series_id"`
	SeriesName     string         `json:"series_name"`
	SeriesAlias    string         `json:"series_alias"`
	SeriesNum      uint           `json:"series_num"`
}

//GetByID ..
func (v *VideoIndex) GetByID(id string) error {
	return db.GetDB().Select(`a.id,a.title,a.desc,cast(EXTRACT(epoch FROM CAST( a.pub_date AS TIMESTAMP)) as bigint) pub_date,a.cover,e.episode_id episode_id ,b.total_num,a.tags,a.is_show,a.is_hide_on_mobile,c.video_series_id series_id,
	c.alias series_alias,c.num series_num,d.name series_name`).
		Table(db.TablePrefix+"_video a").
		Joins("left join (select video_id,count(video_id) total_num from "+db.TablePrefix+"_episode where video_id=? group by video_id) b on a.id=b.video_id", id).
		Joins("left join "+db.TablePrefix+"_video_series_item c on a.id=c.video_id").
		Joins("left join "+db.TablePrefix+"_video_series d on d.id=c.video_series_id").
		Joins("left join (select video_id,id episode_id from "+db.TablePrefix+"_episode where video_id=? order by num asc limit 1) e on a.id=e.video_id", id).
		Where("a.id=?", id).Take(v).Error
}

//Find ..
func (v *VideoIndex) Find() ([]*VideoIndex, error) {
	data := make([]*VideoIndex, 0)
	err := db.GetDB().Select(`a.id,a.title,a.desc,cast(EXTRACT(epoch FROM CAST( a.pub_date AS TIMESTAMP)) as bigint) pub_date,a.cover,e.episode_id episode_id ,b.total_num,a.tags,a.is_show,a.is_hide_on_mobile,c.video_series_id series_id,
	c.alias series_alias,c.num series_num,d.name series_name`).
		Table(db.TablePrefix+"_video a").
		Joins("left join (select video_id,count(video_id) total_num from "+db.TablePrefix+"_episode  group by video_id) b on a.id=b.video_id").
		Joins("left join "+db.TablePrefix+"_video_series_item c on a.id=c.video_id").
		Joins("left join "+db.TablePrefix+"_video_series d on d.id=c.video_series_id").
		Joins("left join (select p.video_id,q.id episode_id from (SELECT video_id, min(num) num from "+db.TablePrefix+"_episode group by (video_id)) p left join "+db.TablePrefix+"_episode q on p.video_id=q.video_id and p.num=q.num) e on a.id=e.video_id").
		Where("a.is_show=?", true).
		Find(&data).Error
	return data, err
}
