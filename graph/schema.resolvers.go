package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"git.9d77v.me/9d77v/pdc/graph/generated"
	"git.9d77v.me/9d77v/pdc/graph/model"
	"github.com/9d77v/go-lib/ptrs"
)

func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	if len(input.VideoURLs) > 0 && len(input.Subtitles) > 0 && len(input.VideoURLs) != len(input.Subtitles) {
		return nil, errors.New("视频与字幕数量不一致")
	}
	return videoService.CreateVideo(input)
}

func (r *mutationResolver) CreateEpisode(ctx context.Context, input model.NewEpisode) (*model.Episode, error) {
	return videoService.CreateEpisode(input)
}

func (r *mutationResolver) UpdateVideo(ctx context.Context, input *model.NewUpdateVideo) (*model.Video, error) {
	return videoService.UpdateVideo(input)
}

func (r *mutationResolver) UpdateEpisode(ctx context.Context, input *model.NewUpdateEpisode) (*model.Episode, error) {
	return videoService.UpdateEpisode(input)
}

func (r *queryResolver) Videos(ctx context.Context, page *int64, pageSize *int64) (*model.VideoConnection, error) {
	o := ptrs.Int64(page)
	l := ptrs.Int64(pageSize)
	if o < 1 {
		o = 1
	}
	if l < 1 {
		l = 1
	}
	if l > 100 {
		l = 100
	}
	o = (o - 1) * l
	con := new(model.VideoConnection)
	total, data, err := videoService.ListVideo(o, l)
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
