package services

import (
	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/video-service/models"
)

func toVideoDto(m *models.Video, scheme string) *model.Video {
	es := make([]*model.Episode, 0, len(m.Episodes))
	for _, e := range m.Episodes {
		sArr := make([]*model.Subtitle, 0, len(e.Subtitles))
		for _, v := range e.Subtitles {
			sArr = append(sArr, &model.Subtitle{
				Name: v.Name,
				URL:  oss.GetOSSPrefix(scheme) + v.URL,
			})
		}
		cover := ""
		if e.Cover != "" {
			cover = oss.GetOSSPrefix(scheme) + e.Cover
		}
		url := ""
		if e.URL != "" {
			url = oss.GetOSSPrefix(scheme) + e.URL
		}
		es = append(es, &model.Episode{
			ID:        int64(e.ID),
			Num:       e.Num,
			Title:     e.Title,
			Desc:      e.Desc,
			Cover:     cover,
			URL:       url,
			Subtitles: sArr,
			CreatedAt: e.CreatedAt.Unix(),
			UpdatedAt: e.UpdatedAt.Unix(),
		})
	}
	cover := ""
	if m.Cover != "" {
		cover = oss.GetOSSPrefix(scheme) + m.Cover
	}
	return &model.Video{
		ID:             int64(m.ID),
		Title:          m.Title,
		Desc:           m.Desc,
		PubDate:        m.PubDate.Unix(),
		Episodes:       es,
		Cover:          cover,
		Tags:           m.Tags,
		IsShow:         m.IsShow,
		IsHideOnMobile: m.IsHideOnMobile,
		Theme:          m.Theme,
		CreatedAt:      m.CreatedAt.Unix(),
		UpdatedAt:      m.UpdatedAt.Unix(),
	}
}

func toVideoSeriesDtos(data []*models.VideoSeries) []*model.VideoSeries {
	result := make([]*model.VideoSeries, 0, len(data))
	for _, m := range data {
		r := toVideoSeriesDto(m)
		result = append(result, r)
	}
	return result
}

func toVideoSeriesDto(m *models.VideoSeries) *model.VideoSeries {
	items := make([]*model.VideoSeriesItem, 0, len(m.Items))
	for _, item := range m.Items {
		items = append(items, &model.VideoSeriesItem{
			VideoID:       int64(item.VideoID),
			EpisodeID:     int64(item.EpisodeID),
			VideoSeriesID: int64(item.VideoSeriesID),
			Alias:         item.Alias,
			Title:         item.Title,
			Num:           int64(item.Num),
		})
	}
	return &model.VideoSeries{
		ID:        int64(m.ID),
		Name:      m.Name,
		Items:     items,
		CreatedAt: m.CreatedAt.Unix(),
		UpdatedAt: m.UpdatedAt.Unix(),
	}
}

func toVideoIndexDto(vi *models.VideoIndex) *model.VideoIndex {
	return &model.VideoIndex{
		ID:        int64(vi.ID),
		Title:     vi.Title,
		Desc:      vi.Desc,
		Cover:     vi.Cover,
		TotalNum:  int64(vi.TotalNum),
		EpisodeID: int64(vi.EpisodeID),
	}
}
