package models

import (
	"github.com/9d77v/pdc/internal/module/base"
)

//VideoSeries 视频系列
type VideoSeries struct {
	*base.Model
	Name  string
	Items []*VideoSeriesItem
}

//NewVideoSeries ..
func NewVideoSeries() *VideoSeries {
	vs := &VideoSeries{}
	vs.Model = base.NewModel()
	return vs
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
