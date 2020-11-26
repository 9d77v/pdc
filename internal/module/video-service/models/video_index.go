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
	IsShow         bool           `json:"is_show"`
	IsHideOnMobile bool           `json:"is_hide_on_mobile"`
	SeriesID       uint           `json:"series_id"`
	SeriesName     string         `json:"series_name"`
	SeriesAlias    string         `json:"series_alias"`
	SeriesNum      uint           `json:"series_num"`
}

//GetByID ..
func (v *VideoIndex) GetByID(id string) error {
	return db.GetDB().Select(`a.id,a.title,a.desc,cast(EXTRACT(epoch FROM CAST( a.pub_date AS TIMESTAMP)) as bigint) pub_date,a.cover,b.total_num,a.tags,a.is_show,a.is_hide_on_mobile,c.video_series_id series_id,
	c.alias series_alias,c.num series_num,d.name series_name`).
		Table(db.TablePrefix+"_video a").
		Joins("left join (select video_id,count(video_id) total_num from "+db.TablePrefix+"_episode where video_id=? group by video_id) b on a.id=b.video_id", id).
		Joins("left join "+db.TablePrefix+"_video_series_item c on a.id=c.video_id").
		Joins("left join "+db.TablePrefix+"_video_series d on d.id=c.video_series_id").
		Where("a.id=?", id).Take(v).Error
}

//Find ..
func (v *VideoIndex) Find() ([]*VideoIndex, error) {
	data := make([]*VideoIndex, 0)
	err := db.GetDB().Select(`a.id,a.title,a.desc,cast(EXTRACT(epoch FROM CAST( a.pub_date AS TIMESTAMP)) as bigint) pub_date,a.cover,b.total_num,a.tags,a.is_show,a.is_hide_on_mobile,c.video_series_id series_id,
	c.alias series_alias,c.num series_num,d.name series_name`).
		Table(db.TablePrefix+"_video a").
		Joins("left join (select video_id,count(video_id) total_num from "+db.TablePrefix+"_episode  group by video_id) b on a.id=b.video_id").
		Joins("left join "+db.TablePrefix+"_video_series_item c on a.id=c.video_id").
		Joins("left join "+db.TablePrefix+"_video_series d on d.id=c.video_series_id").
		Where("a.is_show=?", true).
		Find(&data).Error
	return data, err
}
