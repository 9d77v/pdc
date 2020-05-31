package services

import (
	"context"
	"strings"
	"time"

	"git.9d77v.me/9d77v/pdc/dtos"
	"git.9d77v.me/9d77v/pdc/graph/model"
	"git.9d77v.me/9d77v/pdc/models"
	"git.9d77v.me/9d77v/pdc/utils"
	"github.com/9d77v/go-lib/ptrs"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

//VideoService ..
type VideoService struct {
}

//CreateVideo ..
func (s VideoService) CreateVideo(input model.NewVideo) (*model.Video, error) {
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
func (s VideoService) UpdateVideo(input *model.NewUpdateVideo) (*model.Video, error) {
	video := new(models.Video)
	if err := models.Gorm.First(video, "id=?", input.ID).Error; err != nil {
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
	err := models.Gorm.Model(video).Update(map[string]interface{}{
		"title":      ptrs.String(input.Title),
		"cover":      ptrs.String(input.Cover),
		"pub_date":   time.Unix(ptrs.Int64(input.PubDate), 0),
		"desc":       ptrs.String(input.Desc),
		"characters": cs,
		"staffs":     ss,
		"tags":       input.Tags,
		"is_show":    ptrs.Bool(input.IsShow),
	}).Error
	return &model.Video{ID: int64(video.ID)}, err
}

//CreateEpisode ..
func (s VideoService) CreateEpisode(input model.NewEpisode) (*model.Episode, error) {
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
func (s VideoService) UpdateEpisode(input *model.NewUpdateEpisode) (*model.Episode, error) {
	episode := new(models.Episode)
	if err := models.Gorm.First(episode, "id=?", input.ID).Error; err != nil {
		return nil, err
	}
	cs := make(postgres.Hstore, len(input.Subtitles))
	for _, v := range input.Subtitles {
		cs[v.Name] = &v.URL
	}
	err := models.Gorm.Model(episode).Update(map[string]interface{}{
		"num":       input.Num,
		"title":     ptrs.String(input.Title),
		"cover":     ptrs.String(input.Cover),
		"desc":      ptrs.String(input.Desc),
		"url":       input.URL,
		"subtitles": cs,
	}).Error
	return &model.Episode{ID: int64(episode.ID)}, err
}

//ListVideo ..
func (s VideoService) ListVideo(ctx context.Context, offset, limit int64) (int64, []*model.Video, error) {
	result := make([]*model.Video, 0)
	data := make([]*models.Video, 0)
	filedMap, _ := utils.GetFieldData(ctx, "")
	var err error
	if filedMap["edges"] {
		edgeFieldMap, edgeFields := utils.GetFieldData(ctx, "edges.")
		builder := models.Gorm.Select(utils.ToDBFields(edgeFields, []string{"episodes", "__typename"}))
		if edgeFieldMap["episodes"] {
			err = builder.Preload("Episodes", func(db *gorm.DB) *gorm.DB {
				return models.Gorm.Model(&models.Episode{}).Order("num ASC").Order("id ASC")
			}).Offset(offset).Limit(limit).Order("id DESC").Find(&data).Error
		} else {
			err = builder.Offset(offset).Limit(limit).Order("id DESC").Find(&data).Error
		}
		if err != nil {
			return 0, result, err
		}
	}
	var total int64
	if filedMap["totalCount"] {
		err = models.Gorm.Model(&models.Video{}).Count(&total).Error
		if err != nil {
			return 0, result, err
		}
	}
	for _, m := range data {
		r := dtos.ToVideoDto(m)
		result = append(result, r)
	}
	return total, result, nil
}
