package services

import (
	"strings"
	"time"

	"github.com/9d77v/go-lib/ptrs"

	"git.9d77v.me/9d77v/pdc/graph/model"
	"git.9d77v.me/9d77v/pdc/models"
	"github.com/jinzhu/gorm/dialects/postgres"
)

//VideoService ..
type VideoService struct {
}

//CreateVideo ..
func (s VideoService) CreateVideo(input model.NewVideo) (int64, error) {
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
	}
	err := models.Gorm.Create(m).Error
	return int64(m.ID), err
}

//CreateEpisode ..
func (s VideoService) CreateEpisode(input model.NewEpisode) (int64, error) {
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
	return int64(e.ID), err

}

//ListVideo ..
func (s VideoService) ListVideo() ([]*model.Video, error) {
	result := make([]*model.Video, 0)
	data := make([]*models.Video, 0)
	err := models.Gorm.Preload("Episodes").Find(&data).Error
	if err != nil {
		return result, err
	}
	for _, m := range data {
		es := make([]*model.Episode, 0, len(m.Episodes))
		for _, e := range m.Episodes {
			sArr := make([]*model.Subtitle, 0, len(e.Subtitles))
			for k, v := range e.Subtitles {
				sArr = append(sArr, &model.Subtitle{
					Name: k,
					URL:  *v,
				})
			}
			es = append(es, &model.Episode{
				ID:        int64(e.ID),
				Num:       e.Num,
				Title:     e.Title,
				Desc:      e.Desc,
				Cover:     e.Cover,
				URL:       e.URL,
				Subtitles: sArr,
				CreatedAt: e.CreatedAt.Unix(),
				UpdatedAt: e.UpdatedAt.Unix(),
			})
		}
		cArr := make([]*model.Character, 0, len(m.Characters))
		for k, v := range m.Characters {
			cArr = append(cArr, &model.Character{
				CharacterName: k,
				OriginalName:  *v,
			})
		}
		sArr := make([]*model.Staff, 0, len(m.Staffs))
		for k, v := range m.Staffs {
			sArr = append(sArr, &model.Staff{
				Job:     k,
				Persons: strings.Split(*v, ","),
			})
		}
		r := &model.Video{
			ID:         int64(m.ID),
			Title:      m.Title,
			Desc:       m.Desc,
			PubDate:    m.PubDate.Unix(),
			Episodes:   es,
			Cover:      m.Cover,
			Characters: cArr,
			Staffs:     sArr,
			CreatedAt:  m.CreatedAt.Unix(),
			UpdatedAt:  m.UpdatedAt.Unix(),
		}
		result = append(result, r)
	}

	return result, nil
}
