package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/9d77v/go-lib/ptrs"
	"gorm.io/gorm"

	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/db/mq"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/video-service/models"

	"github.com/9d77v/pdc/internal/utils"
)

//VideoService ..
type VideoService struct {
	base.Service
}

//CreateVideo ..
func (s VideoService) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	m := &models.Video{
		Title:          input.Title,
		Desc:           ptrs.String(input.Desc),
		PubDate:        time.Unix(ptrs.Int64(input.PubDate), 0),
		Cover:          ptrs.String(input.Cover),
		Tags:           input.Tags,
		IsShow:         input.IsShow,
		IsHideOnMobile: input.IsHideOnMobile,
		Theme:          input.Theme,
	}
	err := db.GetDB().Create(m).Error
	if err != nil {
		return &model.Video{}, err
	}
	return &model.Video{ID: int64(m.ID)}, err
}

//AddVideoResource ..
func (s VideoService) AddVideoResource(ctx context.Context, input model.NewVideoResource) (*model.Video, error) {
	video := new(models.Video)
	if err := db.GetDB().Select("id").
		First(video, "id=?", input.ID).Error; err != nil {
		return nil, err
	}
	episodes := make([]*models.Episode, 0, len(input.VideoURLs))
	for i, url := range input.VideoURLs {
		episodes = append(episodes, &models.Episode{
			Num:     float64(i + 1),
			VideoID: uint(input.ID),
			URL:     url,
		})
	}
	err := db.GetDB().Create(&episodes).Error
	if err != nil {
		return &model.Video{}, err
	}
	sendMsgToUpdateES(input.ID)
	return &model.Video{ID: int64(input.ID)}, nil
}

//SaveSubtitles ..
func (s VideoService) SaveSubtitles(ctx context.Context, input model.NewSaveSubtitles) (*model.Video, error) {
	data := make([]*models.Episode, 0)
	if err := db.GetDB().Select("id").Preload("Subtitles", "name=?", input.Subtitles.Name).
		Where("video_id=?", input.ID).Order("num asc").Find(&data).Error; err != nil {
		return nil, err
	}
	if len(input.Subtitles.Urls) == 0 {
		ids := make([]uint, 0, len(data))
		for _, v := range data {
			ids = append(ids, v.ID)
		}
		err := db.GetDB().Where("episode_id in(?) and name=?", ids, input.Subtitles.Name).
			Delete(&models.Subtitle{}).Error
		if err != nil {
			return nil, err
		}
	} else {
		if len(input.Subtitles.Urls) != len(data) {
			return nil, errors.New("视频与字幕数量不一致")
		}
		subtitles := make([]*models.Subtitle, 0, 0)
		for i, d := range data {
			if len(d.Subtitles) == 0 {
				subtitles = append(subtitles, &models.Subtitle{
					EpisodeID: d.ID,
					Name:      input.Subtitles.Name,
					URL:       input.Subtitles.Urls[i],
				})
			} else {
				if d.Subtitles[0].URL != input.Subtitles.Urls[i] {
					d.Subtitles[0].URL = input.Subtitles.Urls[i]
					subtitles = append(subtitles, d.Subtitles[0])
				}
			}
		}
		err := db.GetDB().Save(&subtitles).Error
		if err != nil {
			return nil, err
		}
	}
	return &model.Video{ID: int64(input.ID)}, nil
}

//UpdateVideo ..
func (s VideoService) UpdateVideo(ctx context.Context,
	input model.NewUpdateVideo) (*model.Video, error) {
	video := models.NewVideo()
	fields := s.GetInputFields(ctx)
	if err := s.GetByID(video, uint(input.ID), fields); err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"title":             ptrs.String(input.Title),
		"pub_date":          time.Unix(ptrs.Int64(input.PubDate), 0),
		"desc":              ptrs.String(input.Desc),
		"tags":              &input.Tags,
		"theme":             input.Theme,
		"is_show":           ptrs.Bool(input.IsShow),
		"is_hide_on_mobile": ptrs.Bool(input.IsHideOnMobile),
	}
	if input.Cover != nil {
		updateMap["cover"] = ptrs.String(input.Cover)
	}
	err := db.GetDB().Model(video).Updates(updateMap).Error
	sendMsgToUpdateES(input.ID)
	return &model.Video{ID: int64(video.ID)}, err
}

//CreateEpisode ..
func (s VideoService) CreateEpisode(ctx context.Context,
	input model.NewEpisode) (*model.Episode, error) {
	e := &models.Episode{
		Num:     input.Num,
		VideoID: uint(input.VideoID),
		Title:   ptrs.String(input.Title),
		Desc:    ptrs.String(input.Desc),
		Cover:   ptrs.String(input.Cover),
		URL:     input.URL,
	}
	tx := db.GetDB().Begin()
	err := tx.Create(e).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if len(input.Subtitles) > 0 {
		subtitles := make([]*models.Subtitle, 0, len(input.Subtitles))
		for _, v := range input.Subtitles {
			subtitles = append(subtitles, &models.Subtitle{
				EpisodeID: e.ID,
				Name:      v.Name,
				URL:       v.URL,
			})
		}
		err = tx.Create(&subtitles).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	tx.Commit()
	sendMsgToUpdateES(input.VideoID)
	return &model.Episode{ID: int64(e.ID)}, err
}

//UpdateEpisode ..
func (s VideoService) UpdateEpisode(ctx context.Context,
	input model.NewUpdateEpisode) (*model.Episode, error) {
	episode := new(models.Episode)
	fields := s.GetInputFields(ctx)
	if err := s.GetByID(episode, uint(input.ID), fields); err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"num":   ptrs.Float64(input.Num),
		"title": ptrs.String(input.Title),
		"desc":  ptrs.String(input.Desc),
	}
	tx := db.GetDB().Begin()
	err := tx.Where("episode_id=?", episode.ID).Delete(&models.Subtitle{}).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if len(input.Subtitles) > 0 {
		subtitles := make([]*models.Subtitle, 0, len(input.Subtitles))
		for _, v := range input.Subtitles {
			subtitles = append(subtitles, &models.Subtitle{
				EpisodeID: episode.ID,
				Name:      v.Name,
				URL:       v.URL,
			})
		}
		err := tx.Create(&subtitles).Error
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}
	if input.Cover != nil && *input.Cover != "" {
		updateMap["cover"] = ptrs.String(input.Cover)
	}
	if input.URL != "" {
		updateMap["url"] = input.URL
	}
	err = tx.Model(episode).Updates(updateMap).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &model.Episode{ID: int64(episode.ID)}, err
}

//ListVideo ..
func (s VideoService) ListVideo(ctx context.Context, searchParam model.SearchParam,
	scheme string, isFilterVideoSeries *bool, episodeID *int64) (int64, []*model.Video, error) {
	video := models.NewVideo()
	video.FuzzyQuery(searchParam.Keyword, "title")
	if ptrs.Bool(isFilterVideoSeries) {
		video.FilterVideoSeries()
	}
	if ptrs.Int64(episodeID) > 0 {
		videoID := models.NewEpisode().GetVideoIDByID(uint(ptrs.Int64(episodeID)))
		if videoID == 0 {
			return 0, []*model.Video{}, nil
		}
		searchParam.Ids = []int64{int64(videoID)}
	}
	replaceFunc := func(edgeField base.GraphQLField) error {
		if edgeField.FieldMap["episodes"] {
			video.Preload("Episodes", func(db *gorm.DB) *gorm.DB {
				return db.Model(&models.Episode{}).Order("num ASC").Order("id ASC")
			}).Preload("Episodes.Subtitles")
		}
		return nil
	}
	data := make([]*models.Video, 0)
	total, err := s.GetConnection(ctx, video, searchParam, &data, replaceFunc, "episodes")
	return total, s.getVideos(data, scheme), err
}

//CreateVideoSeries ..
func (s VideoService) CreateVideoSeries(ctx context.Context,
	input model.NewVideoSeries) (*model.VideoSeries, error) {
	m := &models.VideoSeries{
		Name: input.Name,
	}
	err := db.GetDB().Create(m).Error
	if err != nil {
		return &model.VideoSeries{}, err
	}
	return &model.VideoSeries{ID: int64(m.ID)}, err
}

//UpdateVideoSeries ..
func (s VideoService) UpdateVideoSeries(ctx context.Context,
	input model.NewUpdateVideoSeries) (*model.VideoSeries, error) {
	fields := s.GetInputFields(ctx)
	videoSeries := models.NewVideoSeries()
	if err := s.GetByID(videoSeries, uint(input.ID), fields); err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"name": input.Name,
	}
	err := db.GetDB().Model(videoSeries).Updates(updateMap).Error
	return &model.VideoSeries{ID: int64(videoSeries.ID)}, err
}

//CreateVideoSeriesItem ..
func (s VideoService) CreateVideoSeriesItem(ctx context.Context,
	input model.NewVideoSeriesItem) (*model.VideoSeriesItem, error) {
	e := &models.VideoSeriesItem{
		VideoSeriesID: uint(input.VideoSeriesID),
		VideoID:       uint(input.VideoID),
		Alias:         input.Alias,
	}
	videoSeriesItem := models.NewVideoSeriesItem()
	maxNum, err := videoSeriesItem.GetTheMaxNumOfOneVideoSeries(uint(input.VideoSeriesID))
	if err != nil {
		return nil, err
	}
	e.Num = maxNum + 1
	err = db.GetDB().Create(e).Error
	sendMsgToUpdateES(input.VideoID)
	return &model.VideoSeriesItem{VideoID: input.VideoID, VideoSeriesID: input.VideoSeriesID}, err
}

//UpdateVideoSeriesItem ..
func (s VideoService) UpdateVideoSeriesItem(ctx context.Context,
	input model.NewUpdateVideoSeriesItem) (*model.VideoSeriesItem, error) {
	videoSeriesItem := models.NewVideoSeriesItem()
	fields := s.GetInputFields(ctx)
	err := videoSeriesItem.GetByVideoIDVideoSeriesID(fields, uint(input.VideoID), uint(input.VideoSeriesID))
	if err != nil {
		return nil, err
	}
	updateMap := map[string]interface{}{
		"alias": input.Alias,
	}
	err = db.GetDB().Model(videoSeriesItem).Updates(updateMap).Error
	sendMsgToUpdateES(input.VideoID)
	return &model.VideoSeriesItem{VideoID: input.VideoID, VideoSeriesID: input.VideoSeriesID}, err
}

//ListVideoSeries ..
func (s VideoService) ListVideoSeries(ctx context.Context, searchParam model.SearchParam, episodeID *int64) (int64, []*model.VideoSeries, error) {
	result := make([]*model.VideoSeries, 0)
	data := make([]*models.VideoSeries, 0)
	field := base.NewGraphQLField(ctx, "")
	var err error
	videoSeries := models.NewVideoSeries()
	videoSeries.FuzzyQuery(searchParam.Keyword, "name")
	omitFields := []string{"title"}
	videoSeriesItem := models.NewVideoSeriesItem()
	if episodeID != nil && ptrs.Int64(episodeID) > 0 {
		videoID := models.NewEpisode().GetVideoIDByID(uint(ptrs.Int64(episodeID)))
		if videoID == 0 {
			return 0, result, errors.New("episode not found")
		}
		videoSeriesID := videoSeriesItem.GetVideoSeriesIDByVideoID(videoID)
		if videoSeriesID == 0 {
			return 0, result, nil
		}
		searchParam.Ids = []int64{int64(videoSeriesItem.VideoSeriesID)}
		omitFields = append(omitFields, "videoID", "episodeID")
	}
	var total int64
	offset, limit := s.GetPageInfo(searchParam)
	if field.FieldMap["totalCount"] {
		if limit == -1 {
			total = int64(len(data))
		} else {
			total, err = videoSeries.Count(videoSeries)
			if err != nil {
				return 0, result, err
			}
		}
	}
	if field.FieldMap["edges"] {
		edgeField := base.NewGraphQLField(ctx, "edges.")
		err = videoSeries.
			Select(edgeField.Fields, "items").
			IDArrayQuery(videoSeries.ToUintIDs(searchParam.Ids)).
			Pagination(offset, limit).
			Sort(searchParam.Sorts).
			Find(&data)
		if err != nil {
			return 0, result, err
		}
		if edgeField.FieldMap["items"] && len(data) > 0 {
			itemField := base.NewGraphQLField(ctx, "edges.items.")
			ids := make([]uint, 0)
			for _, v := range data {
				ids = append(ids, v.ID)
			}
			items := make([]*models.VideoSeriesItem, 0)
			tableVideo := new(models.Video).TableName()
			tableEpisode := new(models.Episode).TableName()
			tableVideoSeriesItem := new(models.VideoSeriesItem).TableName()
			if itemField.FieldMap["title"] {
				itemField.Fields = append(itemField.Fields, tableVideo+".\"title\"")
				videoSeriesItem.LeftJoin(fmt.Sprintf("%s on %s.id=%s.video_id",
					tableVideo, tableVideo, tableVideoSeriesItem))
			}
			if itemField.FieldMap["episodeID"] {
				itemField.Fields = append(itemField.Fields,
					tableEpisode+".\"episode_id\"",
					tableVideoSeriesItem+".\"video_id\"")
				videoSeriesItem.LeftJoin("(select p.video_id,q.id episode_id from (SELECT video_id, min(num) num from " + tableEpisode + " group by (video_id)) p left join " + tableEpisode + "  q on p.video_id=q.video_id and p.num=q.num) " + tableEpisode +
					" on " + tableEpisode + ".video_id=" + tableVideoSeriesItem + ".video_id")
			}
			subErr := videoSeriesItem.
				Select(itemField.Fields, omitFields...).
				IDArrayQuery(ids, "video_series_id").
				Order("video_series_id asc,num asc").
				Find(&items)
			if subErr != nil {
				return 0, result, subErr
			}
			videoSeries.AddItemsToList(data, items)
		}
	}
	return total, s.getVideoSerieses(data), nil
}

func sendMsgToUpdateES(videoID int64) {
	guid, err := mq.GetClient().PublishAsync(mq.SubjectVideo, []byte(strconv.Itoa(int(videoID))),
		utils.AckHandler)
	if err != nil {
		log.Println("mq publish failed,guid:", guid, " error:", err)
	}
}
