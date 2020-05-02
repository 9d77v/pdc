package dtos

import (
	"strings"

	"git.9d77v.me/9d77v/pdc/graph/model"
	"git.9d77v.me/9d77v/pdc/models"
)

//ToVideoDto ...
func ToVideoDto(m *models.Video) *model.Video {
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
	return &model.Video{
		ID:         int64(m.ID),
		Title:      m.Title,
		Desc:       m.Desc,
		PubDate:    m.PubDate.Unix(),
		Episodes:   es,
		Cover:      m.Cover,
		Characters: cArr,
		Staffs:     sArr,
		Tags:       m.Tags,
		IsShow:     m.IsShow,
		CreatedAt:  m.CreatedAt.Unix(),
		UpdatedAt:  m.UpdatedAt.Unix(),
	}
}
