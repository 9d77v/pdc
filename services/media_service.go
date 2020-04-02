package services

import (
	"strings"
	"time"

	"git.9d77v.me/9d77v/pdc/graph/model"
	"git.9d77v.me/9d77v/pdc/models"
	"github.com/jinzhu/gorm/dialects/postgres"
)

//MediaService ..
type MediaService struct {
}

//CreateMedia ..
func (s MediaService) CreateMedia(input model.NewMedia) (int64, error) {
	cs := make(postgres.Hstore, len(input.Characters))
	for _, v := range input.Characters {
		cs[v.CharacterName] = &v.OriginalName
	}
	ss := make(postgres.Hstore, len(input.Staffs))
	for _, v := range input.Staffs {
		staffs := strings.Join(v.Persons, ",")
		ss[v.Job] = &staffs
	}
	desc := ""
	if input.Desc != nil {
		desc = *input.Desc
	}
	var pubdate int64
	if input.PubDate != nil {
		pubdate = *input.PubDate
	}
	cover := ""
	if input.Cover != nil {
		cover = *input.Cover
	}
	m := &models.Media{
		Title:      input.Title,
		Desc:       desc,
		PubDate:    time.Unix(pubdate, 0),
		Cover:      cover,
		Characters: cs,
		Staffs:     ss,
	}
	err := models.Gorm.Create(m).Error
	return int64(m.ID), err
}

//CreateEpisode ..
func (s MediaService) CreateEpisode(input model.NewEpisode) (int64, error) {
	desc := ""
	if input.Desc != nil {
		desc = *input.Desc
	}
	cs := make(postgres.Hstore, len(input.Subtitles))
	for _, v := range input.Subtitles {
		cs[v.Name] = &v.URL
	}
	cover := ""
	if input.Cover != nil {
		cover = *input.Cover
	}
	e := &models.Episode{
		Order:     input.Order,
		MediaID:   input.MediaID,
		Title:     input.Title,
		Desc:      desc,
		Cover:     cover,
		URL:       input.URL,
		Subtitles: cs,
	}
	err := models.Gorm.Create(e).Error
	return int64(e.ID), err
}

//ListMedia ..
func (s MediaService) ListMedia() ([]*model.Media, error) {
	result := make([]*model.Media, 0)
	data := make([]*models.Media, 0)
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
				Order:     e.Order,
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
		r := &model.Media{
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

//PresignedURL ..
func (s MediaService) PresignedURL(bucketName, objectName string) (string, error) {
	u, err := models.MinioClient.PresignedPutObject(bucketName, objectName, 12*time.Hour)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
