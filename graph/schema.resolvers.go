package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"
	"time"

	"git.9d77v.me/9d77v/pdc/graph/generated"
	"git.9d77v.me/9d77v/pdc/graph/model"
	"git.9d77v.me/9d77v/pdc/models"
	"github.com/jinzhu/gorm/dialects/postgres"
)

func (r *mutationResolver) CreateMedia(ctx context.Context, input model.NewMedia) (int64, error) {
	cs := make(postgres.Hstore, len(input.Characters))
	for _, v := range input.Characters {
		cs[v.CharacterName] = &v.OriginalName
	}
	ss := make(postgres.Hstore, len(input.Staffs))
	for _, v := range input.Staffs {
		staffs := strings.Join(v.Persons, ",")
		ss[v.Job] = &staffs
	}
	m := &models.Media{
		Title:      input.Title,
		Desc:       input.Desc,
		PubDate:    time.Unix(input.PubDate, 0),
		Cover:      input.Cover,
		Characters: cs,
		Staffs:     ss,
	}
	err := models.Gorm.Create(m).Error
	return int64(m.ID), err
}

func (r *mutationResolver) CreateEpisode(ctx context.Context, input model.NewEpisode) (int64, error) {
	e := &models.Episode{
		Order:    input.Order,
		MediaID:  input.MediaID,
		Title:    input.Title,
		Desc:     input.Desc,
		Cover:    input.Cover,
		URL:      input.URL,
		Subtitle: input.Subtitle,
	}
	err := models.Gorm.Create(e).Error
	return int64(e.ID), err
}

func (r *queryResolver) ListMedia(ctx context.Context) ([]*model.Media, error) {
	result := make([]*model.Media, 0)
	data := make([]*models.Media, 0)
	err := models.Gorm.Preload("Episodes").Find(&data).Error
	if err != nil {
		return result, err
	}
	for _, m := range data {
		es := make([]*model.Episode, 0)
		for _, e := range m.Episodes {
			es = append(es, &model.Episode{
				ID:        int64(e.ID),
				Order:     e.Order,
				Title:     e.Title,
				Desc:      e.Desc,
				Cover:     e.Cover,
				URL:       e.URL,
				Subtitle:  e.Subtitle,
				CreatedAt: e.CreatedAt.Unix(),
				UpdatedAt: e.UpdatedAt.Unix(),
			})
		}
		r := &model.Media{
			ID:        int64(m.ID),
			Title:     m.Title,
			Desc:      m.Desc,
			PubDate:   m.PubDate.Unix(),
			Episodes:  es,
			Cover:     m.Cover,
			CreatedAt: m.CreatedAt.Unix(),
			UpdatedAt: m.UpdatedAt.Unix(),
		}
		result = append(result, r)
	}

	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
