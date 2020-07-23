package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/9d77v/pdc/dtos"
	"github.com/9d77v/pdc/graph/generated"
	"github.com/9d77v/pdc/graph/model"
	"github.com/9d77v/pdc/middleware"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return userService.CreateUser(ctx, input)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.NewUpdateUser) (*model.User, error) {
	return userService.UpdateUser(ctx, input)
}

func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*model.LoginResponse, error) {
	return userService.Login(ctx, username, password)
}

func (r *mutationResolver) RefreshToken(ctx context.Context, refreshToken string) (*model.LoginResponse, error) {
	return userService.RefreshToken(ctx, refreshToken)
}

func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	if len(input.VideoURLs) > 0 && input.Subtitles != nil && len(input.Subtitles.Urls) > 0 && len(input.VideoURLs) != len(input.Subtitles.Urls) {
		return nil, errors.New("视频与字幕数量不一致")
	}
	return videoService.CreateVideo(input)
}

func (r *mutationResolver) UpdateVideo(ctx context.Context, input model.NewUpdateVideo) (*model.Video, error) {
	return videoService.UpdateVideo(ctx, input)
}

func (r *mutationResolver) CreateEpisode(ctx context.Context, input model.NewEpisode) (*model.Episode, error) {
	return videoService.CreateEpisode(input)
}

func (r *mutationResolver) UpdateEpisode(ctx context.Context, input model.NewUpdateEpisode) (*model.Episode, error) {
	return videoService.UpdateEpisode(ctx, input)
}

func (r *mutationResolver) UpdateSubtitle(ctx context.Context, input model.NewUpdateSubtitles) (*model.Video, error) {
	return videoService.UpdateSubtitle(ctx, input)
}

func (r *mutationResolver) CreateThing(ctx context.Context, input model.NewThing) (*model.Thing, error) {
	user := middleware.ForContext(ctx)
	return thingService.CreateThing(input, int64(user.ID))
}

func (r *mutationResolver) UpdateThing(ctx context.Context, input model.NewUpdateThing) (*model.Thing, error) {
	user := middleware.ForContext(ctx)
	return thingService.UpdateThing(ctx, input, int64(user.ID))
}

func (r *queryResolver) PresignedURL(ctx context.Context, bucketName string, objectName string) (string, error) {
	return commonService.PresignedURL(bucketName, objectName)
}

func (r *queryResolver) Users(ctx context.Context, keyword *string, page *int64, pageSize *int64, ids []int64, sorts []*model.Sort) (*model.UserConnection, error) {
	con := new(model.UserConnection)
	total, data, err := userService.ListUser(ctx, keyword, page, pageSize, ids, sorts)
	con.TotalCount = total
	con.Edges = data
	return con, err
}

func (r *queryResolver) UserInfo(ctx context.Context, uid *int64) (*model.User, error) {
	return dtos.ToUserDto(middleware.ForContext(ctx)), nil
}

func (r *queryResolver) Videos(ctx context.Context, keyword *string, page *int64, pageSize *int64, ids []int64, sorts []*model.Sort) (*model.VideoConnection, error) {
	con := new(model.VideoConnection)
	total, data, err := videoService.ListVideo(ctx, keyword, page, pageSize, ids, sorts)
	con.TotalCount = total
	con.Edges = data
	return con, err
}

func (r *queryResolver) Things(ctx context.Context, keyword *string, page *int64, pageSize *int64, ids []int64, sorts []*model.Sort) (*model.ThingConnection, error) {
	user := middleware.ForContext(ctx)
	con := new(model.ThingConnection)
	total, data, err := thingService.ListThing(ctx, keyword, page, pageSize, ids, sorts, int64(user.ID))
	con.TotalCount = total
	con.Edges = data
	return con, err
}

func (r *queryResolver) ThingSeries(ctx context.Context, dimension string, index string, start *int64, end *int64, status []int64) ([]*model.SerieData, error) {
	user := middleware.ForContext(ctx)
	return thingService.ThingSeries(ctx, dimension, index, start, end, status, int64(user.ID))
}

func (r *queryResolver) ThingAnalyze(ctx context.Context, dimension string, index string, start *int64, group string) (*model.PieLineSerieData, error) {
	user := middleware.ForContext(ctx)
	return thingService.ThingAnalyze(ctx, dimension, index, start, group, int64(user.ID))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
