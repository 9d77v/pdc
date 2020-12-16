package models

import (
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
)

//VideoSeriesRepository ..
type VideoSeriesRepository interface {
	base.Repository
	AddItemsToList(data []*VideoSeries, items []*VideoSeriesItem)
}

//VideoSeries 视频系列
type VideoSeries struct {
	base.DefaultModel
	Name  string
	Items []*VideoSeriesItem
}

//NewVideoSeries ..
func NewVideoSeries() *VideoSeries {
	vs := &VideoSeries{}
	vs.DefaultModel = base.DefaultModel{Model: &base.Model{}}
	vs.SetDB(db.GetDB())
	return vs
}

//TableName ..
func (m *VideoSeries) TableName() string {
	return db.TablePrefix() + "video_series"
}

//AddItemsToList ..
func (m *VideoSeries) AddItemsToList(data []*VideoSeries, items []*VideoSeriesItem) {
	itemMap := make(map[uint][]*VideoSeriesItem)
	for _, v := range items {
		if itemMap[v.VideoSeriesID] == nil {
			itemMap[v.VideoSeriesID] = make([]*VideoSeriesItem, 0)
		}
		itemMap[v.VideoSeriesID] = append(itemMap[v.VideoSeriesID], v)
	}
	for _, v := range data {
		v.Items = itemMap[v.ID]
	}
}
