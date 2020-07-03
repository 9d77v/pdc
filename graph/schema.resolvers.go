package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"git.9d77v.me/9d77v/pdc/graph/generated"
	"git.9d77v.me/9d77v/pdc/graph/model"
)

func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	if len(input.VideoURLs) > 0 && input.Subtitles != nil && len(input.Subtitles.Urls) > 0 && len(input.VideoURLs) != len(input.Subtitles.Urls) {
		return nil, errors.New("视频与字幕数量不一致")
	}
	return videoService.CreateVideo(input)
}

func (r *mutationResolver) CreateEpisode(ctx context.Context, input model.NewEpisode) (*model.Episode, error) {
	return videoService.CreateEpisode(input)
}

func (r *mutationResolver) CreateThing(ctx context.Context, input model.NewThing) (*model.Thing, error) {
	return thingService.CreateThing(input)
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return userService.CreateUser(ctx, input)
}

func (r *mutationResolver) UpdateVideo(ctx context.Context, input model.NewUpdateVideo) (*model.Video, error) {
	return videoService.UpdateVideo(ctx, input)
}

func (r *mutationResolver) UpdateEpisode(ctx context.Context, input model.NewUpdateEpisode) (*model.Episode, error) {
	return videoService.UpdateEpisode(ctx, input)
}

func (r *mutationResolver) UpdateThing(ctx context.Context, input model.NewUpdateThing) (*model.Thing, error) {
	return thingService.UpdateThing(ctx, input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.NewUpdateUser) (*model.User, error) {
	return userService.UpdateUser(ctx, input)
}

func (r *queryResolver) Videos(ctx context.Context, page *int64, pageSize *int64, ids []int64, sorts []*model.Sort) (*model.VideoConnection, error) {
	con := new(model.VideoConnection)
	total, data, err := videoService.ListVideo(ctx, page, pageSize, ids, sorts)
	con.TotalCount = total
	con.Edges = data
	return con, err
}

func (r *queryResolver) Things(ctx context.Context, page *int64, pageSize *int64, ids []int64, sorts []*model.Sort) (*model.ThingConnection, error) {
	con := new(model.ThingConnection)
	total, data, err := thingService.ListThing(ctx, page, pageSize, ids, sorts)
	con.TotalCount = total
	con.Edges = data
	return con, err
}

func (r *queryResolver) ThingSeries(ctx context.Context, dimension string, index string, start *int64, end *int64, status []int64) ([]*model.SerieData, error) {
	return thingService.ThingSeries(ctx, dimension, index, start, end, status, 1)
}

func (r *queryResolver) ThingAnalyze(ctx context.Context, dimension string, index string, start *int64, group string) (*model.PieLineSerieData, error) {
	return thingService.ThingAnalyze(ctx, dimension, index, start, group, 1)
}

func (r *queryResolver) Users(ctx context.Context, page *int64, pageSize *int64, ids []int64, sorts []*model.Sort) (*model.UserConnection, error) {
	con := new(model.UserConnection)
	total, data, err := userService.ListUser(ctx, page, pageSize, ids, sorts)
	con.TotalCount = total
	con.Edges = data
	return con, err
}

func (r *queryResolver) PresignedURL(ctx context.Context, bucketName string, objectName string) (string, error) {
	return commonService.PresignedURL(bucketName, objectName)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
