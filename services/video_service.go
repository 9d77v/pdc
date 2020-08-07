package services

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/dtos"
	"github.com/9d77v/pdc/graph/model"
	"github.com/9d77v/pdc/models"
	"github.com/9d77v/pdc/utils"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

//VideoService ..
type VideoService struct {
}

//CreateVideo ..
func (s VideoService) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	cs := make(postgres.Hstore, len(input.Characters))
	for _, v := range input.Characters {
		cs[v.CharacterName] = &v.OriginalName
	}
	ss := make(postgres.Hstore, len(input.Staffs))
	for _, v := range input.Staffs {
		staffs := strings.Join(v.Persons, ",")
		ss[v.Job] = &staffs
	}
	m := &models.Video{
		Title:      input.Title,
		Desc:       ptrs.String(input.Desc),
		PubDate:    time.Unix(ptrs.Int64(input.PubDate), 0),
		Cover:      ptrs.String(input.Cover),
		Characters: cs,
		Staffs:     ss,
		Tags:       input.Tags,
		IsShow:     input.IsShow,
	}
	tx := models.Gorm.Begin()
	err := tx.Create(m).Error
	if err != nil {
		tx.Rollback()
		return &model.Video{}, err
	}
	for i, url := range input.VideoURLs {
		e := &models.Episode{
			Num:     float64(i + 1),
			VideoID: int64(m.ID),
			URL:     url,
		}
		cs := make(postgres.Hstore, len(input.VideoURLs))

		if input.Subtitles != nil && len(input.Subtitles.Urls) > 0 {
			cs[input.Subtitles.Name] = &input.Subtitles.Urls[i]
			e.Subtitles = cs
		}
		err := models.Gorm.Create(e).Error
		if err != nil {
			tx.Rollback()
			return &model.Video{}, err
		}
	}
	tx.Commit()
	return &model.Video{ID: int64(m.ID)}, err
}

//UpdateVideo ..
func (s VideoService) UpdateVideo(ctx context.Context, input model.NewUpdateVideo) (*model.Video, error) {
	video := new(models.Video)
	fields := make([]string, 0)
	varibales := graphql.GetRequestContext(ctx).Variables
	for k := range varibales["input"].(map[string]interface{}) {
		fields = append(fields, k)
	}
	if err := models.Gorm.Select(utils.ToDBFields(fields)).
		First(video, "id=?", input.ID).Error; err != nil {
		return nil, err
	}
	cs := make(postgres.Hstore, len(input.Characters))
	for _, v := range input.Characters {
		cs[v.CharacterName] = &v.OriginalName
	}
	ss := make(postgres.Hstore, len(input.Staffs))
	for _, v := range input.Staffs {
		staffs := strings.Join(v.Persons, ",")
		ss[v.Job] = &staffs
	}
	updateMap := map[string]interface{}{
		"title":      ptrs.String(input.Title),
		"pub_date":   time.Unix(ptrs.Int64(input.PubDate), 0),
		"desc":       ptrs.String(input.Desc),
		"characters": cs,
		"staffs":     ss,
		"tags":       input.Tags,
		"is_show":    ptrs.Bool(input.IsShow),
	}
	if input.Cover != nil {
		updateMap["cover"] = ptrs.String(input.Cover)
	}
	err := models.Gorm.Model(video).Update(updateMap).Error
	return &model.Video{ID: int64(video.ID)}, err
}

//CreateEpisode ..
func (s VideoService) CreateEpisode(ctx context.Context, input model.NewEpisode) (*model.Episode, error) {
	cs := make(postgres.Hstore, len(input.Subtitles))
	for _, v := range input.Subtitles {
		cs[v.Name] = &v.URL
	}
	e := &models.Episode{
		Num:       input.Num,
		VideoID:   input.VideoID,
		Title:     ptrs.String(input.Title),
		Desc:      ptrs.String(input.Desc),
		Cover:     ptrs.String(input.Cover),
		URL:       input.URL,
		Subtitles: cs,
	}
	err := models.Gorm.Create(e).Error
	return &model.Episode{ID: int64(e.ID)}, err
}

//UpdateEpisode ..
func (s VideoService) UpdateEpisode(ctx context.Context, input model.NewUpdateEpisode) (*model.Episode, error) {
	episode := new(models.Episode)
	varibales := graphql.GetRequestContext(ctx).Variables
	fields := make([]string, 0)
	for k := range varibales["input"].(map[string]interface{}) {
		fields = append(fields, k)
	}
	if err := models.Gorm.Select(utils.ToDBFields(fields)).First(episode, "id=?", input.ID).Error; err != nil {
		return nil, err
	}
	cs := make(postgres.Hstore, len(input.Subtitles))
	for _, v := range input.Subtitles {
		cs[v.Name] = &v.URL
	}
	updateMap := map[string]interface{}{
		"num":       ptrs.Float64(input.Num),
		"title":     ptrs.String(input.Title),
		"desc":      ptrs.String(input.Desc),
		"subtitles": cs,
	}
	if input.Cover != nil {
		updateMap["cover"] = ptrs.String(input.Cover)
	}
	if input.URL != "" {
		updateMap["url"] = input.URL
	}
	err := models.Gorm.Model(episode).Update(updateMap).Error
	return &model.Episode{ID: int64(episode.ID)}, err
}

//UpdateSubtitle ..
func (s VideoService) UpdateSubtitle(ctx context.Context, input model.NewUpdateSubtitles) (*model.Video, error) {
	data := make([]*models.Episode, 0)
	if err := models.Gorm.Select("id,subtitles").Where("video_id=?", input.ID).Order("num asc").Find(&data).Error; err != nil {
		return nil, err
	}
	tx := models.Gorm.Begin()
	if len(input.Subtitles.Urls) == 0 {
		for _, d := range data {
			if d.Subtitles == nil {
				continue
			}
			cs := make(postgres.Hstore, len(d.Subtitles))
			for k, v := range d.Subtitles {
				if k != input.Subtitles.Name {
					cs[k] = v
				}
			}
			err := models.Gorm.Model(d).Update(map[string]interface{}{
				"subtitles": cs,
			}).Error
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	} else {
		if len(input.Subtitles.Urls) != len(data) {
			return nil, errors.New("视频与字幕数量不一致")
		}
		for i, d := range data {
			if d.Subtitles == nil {
				cs := make(postgres.Hstore, len(d.Subtitles))
				cs[input.Subtitles.Name] = &input.Subtitles.Urls[i]
				d.Subtitles = cs
			} else {
				d.Subtitles[input.Subtitles.Name] = &input.Subtitles.Urls[i]
			}
			err := models.Gorm.Model(d).Update(map[string]interface{}{
				"subtitles": d.Subtitles,
			}).Error
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}
	tx.Commit()
	return &model.Video{ID: int64(input.ID)}, nil
}

//UpdateMobileVideo ..
func (s VideoService) UpdateMobileVideo(ctx context.Context, input *model.NewUpdateMobileVideos) (*model.Video, error) {
	data := make([]*models.Episode, 0)
	if err := models.Gorm.Select("id,subtitles").Where("video_id=?", input.ID).Order("num asc").Find(&data).Error; err != nil {
		return nil, err
	}
	tx := models.Gorm.Begin()
	if len(input.VideoURLs) != len(data) {
		return nil, errors.New("移动端视频与已有视频数量不一致")
	}

	for i, d := range data {
		err := models.Gorm.Model(d).Update(map[string]interface{}{
			"mobile_url": input.VideoURLs[i],
		}).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	return &model.Video{ID: int64(input.ID)}, nil
}

//ListVideo ..
func (s VideoService) ListVideo(ctx context.Context, keyword *string,
	page, pageSize *int64, ids []int64, sorts []*model.Sort,
	scheme string, isCombo *bool) (int64, []*model.Video, error) {
	offset, limit := GetPageInfo(page, pageSize)
	result := make([]*model.Video, 0)
	data := make([]*models.Video, 0)
	fieldMap, _ := utils.GetFieldData(ctx, "")
	var err error
	builder := models.Gorm
	if keyword != nil && ptrs.String(keyword) != "" {
		builder = builder.Where("title like ?", "%"+ptrs.String(keyword)+"%")
	}
	if ptrs.Bool(isCombo) {
		builder = builder.Where("NOT EXISTS (select video_id from " + models.DBPrefix + "_video_series_item where video_id=id)")
	}
	if fieldMap["edges"] {
		edgeFieldMap, edgeFields := utils.GetFieldData(ctx, "edges.")
		builder = builder.Select(utils.ToDBFields(edgeFields, "episodes", "__typename"))
		if len(ids) > 0 {
			builder = builder.Where("id in (?)", ids)
		}
		subBuilder := builder
		if limit > 0 {
			subBuilder = subBuilder.Offset(offset).Limit(limit)
		}
		for _, v := range sorts {
			if v.IsAsc {
				subBuilder = subBuilder.Order(v.Field + " ASC")
			} else {
				subBuilder = subBuilder.Order(v.Field + " DESC")
			}
		}
		if edgeFieldMap["episodes"] {
			err = subBuilder.Preload("Episodes", func(db *gorm.DB) *gorm.DB {
				return models.Gorm.Model(&models.Episode{}).Order("num ASC").Order("id ASC")
			}).Find(&data).Error
		} else {
			err = subBuilder.Find(&data).Error
		}
		if err != nil {
			return 0, result, err
		}
	}
	var total int64
	if fieldMap["totalCount"] {
		if limit == -1 {
			total = int64(len(data))
		} else {
			err = builder.Model(&models.Video{}).Count(&total).Error
			if err != nil {
				return 0, result, err
			}
		}
	}
	for _, m := range data {
		r := dtos.ToVideoDto(m, scheme)
		result = append(result, r)
	}
	return total, result, nil
}

//CreateVideoSeries ..
func (s VideoService) CreateVideoSeries(ctx context.Context, input model.NewVideoSeries) (*model.VideoSeries, error) {
	m := &models.VideoSeries{
		Name: input.Name,
	}
	err := models.Gorm.Create(m).Error
	if err != nil {
		return &model.VideoSeries{}, err
	}
	return &model.VideoSeries{ID: int64(m.ID)}, err
}

//UpdateVideoSeries ..
func (s VideoService) UpdateVideoSeries(ctx context.Context, input model.NewUpdateVideoSeries) (*model.VideoSeries, error) {
	videoSeries := new(models.VideoSeries)
	fields := make([]string, 0)
	varibales := graphql.GetRequestContext(ctx).Variables
	for k := range varibales["input"].(map[string]interface{}) {
		fields = append(fields, k)
	}
	if err := models.Gorm.Select(utils.ToDBFields(fields)).
		First(videoSeries, "id=?", input.ID).Error; err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"name": input.Name,
	}
	err := models.Gorm.Model(videoSeries).Update(updateMap).Error
	return &model.VideoSeries{ID: int64(videoSeries.ID)}, err
}

//CreateVideoSeriesItem ..
func (s VideoService) CreateVideoSeriesItem(ctx context.Context, input model.NewVideoSeriesItem) (*model.VideoSeriesItem, error) {
	e := &models.VideoSeriesItem{
		VideoSeriesID: uint(input.VideoSeriesID),
		VideoID:       uint(input.VideoID),
		Alias:         input.Alias,
	}
	maxItem := &models.VideoSeriesItem{}
	err := models.Gorm.Select("max(num) num").Where("video_series_id=?", input.VideoSeriesID).Take(maxItem).Error
	if err != nil {
		return nil, err
	}
	e.Num = maxItem.Num + 1
	err = models.Gorm.Create(e).Error
	return &model.VideoSeriesItem{VideoID: input.VideoID, VideoSeriesID: input.VideoSeriesID}, err
}

//UpdateVideoSeriesItem ..
func (s VideoService) UpdateVideoSeriesItem(ctx context.Context, input model.NewUpdateVideoSeriesItem) (*model.VideoSeriesItem, error) {
	item := new(models.VideoSeriesItem)
	varibales := graphql.GetRequestContext(ctx).Variables
	fields := make([]string, 0)
	for k := range varibales["input"].(map[string]interface{}) {
		fields = append(fields, k)
	}
	if err := models.Gorm.Select(utils.ToDBFields(fields)).First(item, "video_id=? and video_series_id=?",
		input.VideoID, input.VideoSeriesID).Error; err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"alias": input.Alias,
	}
	err := models.Gorm.Model(item).Update(updateMap).Error
	return &model.VideoSeriesItem{VideoID: input.VideoID, VideoSeriesID: input.VideoSeriesID}, err
}

//ListVideoSeries ..
func (s VideoService) ListVideoSeries(ctx context.Context, keyword *string, videoID *int64,
	page, pageSize *int64, ids []int64, sorts []*model.Sort) (int64, []*model.VideoSeries, error) {
	offset, limit := GetPageInfo(page, pageSize)
	result := make([]*model.VideoSeries, 0)
	data := make([]*models.VideoSeries, 0)
	fieldMap, _ := utils.GetFieldData(ctx, "")
	if videoID != nil && ptrs.Int64(videoID) > 0 {
		//获取视频所属系列
		item := new(models.VideoSeriesItem)
		err := models.Gorm.Select("video_series_id").Where("video_id=?", ptrs.Int64(videoID)).Take(item).Error
		if err != nil {
			return 0, result, nil
		}
		ids = []int64{int64(item.VideoSeriesID)}
	}
	var err error
	builder := models.Gorm
	if keyword != nil && ptrs.String(keyword) != "" {
		builder = builder.Where("name like ?", "%"+ptrs.String(keyword)+"%")
	}
	if fieldMap["edges"] {
		edgeFieldMap, edgeFields := utils.GetFieldData(ctx, "edges.")
		builder = builder.Select(utils.ToDBFields(edgeFields, "items", "__typename"))
		if len(ids) > 0 {
			builder = builder.Where("id in (?)", ids)
		}
		subBuilder := builder
		if limit > 0 {
			subBuilder = subBuilder.Offset(offset).Limit(limit)
		}
		for _, v := range sorts {
			if v.IsAsc {
				subBuilder = subBuilder.Order(v.Field + " ASC")
			} else {
				subBuilder = subBuilder.Order(v.Field + " DESC")
			}
		}
		err = subBuilder.Find(&data).Error
		if err != nil {
			return 0, result, err
		}
		if edgeFieldMap["items"] && len(data) > 0 {
			itemFIeldMap, itemFields := utils.GetFieldData(ctx, "edges.items.")
			ids := make([]uint, 0)
			for _, v := range data {
				ids = append(ids, v.ID)
			}
			items := make([]*models.VideoSeriesItem, 0)
			itemBuilder := models.Gorm
			videoTableName := models.DBPrefix + "_video"
			videoSeriesItemTableName := models.DBPrefix + "_video_series_item"
			if itemFIeldMap["title"] {
				itemBuilder = itemBuilder.Select(append(utils.ToDBFields(itemFields, "title", "__typename"),
					videoTableName+".\"title\"")).
					Joins(fmt.Sprintf("left join %s on %s.id=%s.video_id",
						videoTableName, videoTableName, videoSeriesItemTableName))
			} else {
				itemBuilder = itemBuilder.Select(utils.ToDBFields(itemFields, "title", "__typename"))
			}
			subErr := itemBuilder.Where("video_series_id in (?)", ids).
				Order("video_series_id asc").Order("num asc").Find(&items).Error
			if subErr != nil {
				return 0, result, subErr
			}
			itemMap := make(map[uint][]*models.VideoSeriesItem)
			for _, v := range items {
				if itemMap[v.VideoSeriesID] == nil {
					itemMap[v.VideoSeriesID] = make([]*models.VideoSeriesItem, 0)
				}
				itemMap[v.VideoSeriesID] = append(itemMap[v.VideoSeriesID], v)
			}
			for _, v := range data {
				v.Items = itemMap[v.ID]
			}
		}
	}
	var total int64
	if fieldMap["totalCount"] {
		if limit == -1 {
			total = int64(len(data))
		} else {
			err = builder.Model(&models.VideoSeries{}).Count(&total).Error
			if err != nil {
				return 0, result, err
			}
		}
	}
	for _, m := range data {
		r := dtos.ToVideoSeriesDto(m)
		result = append(result, r)
	}
	return total, result, nil
}
