package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"git.9d77v.me/9d77v/pdc/graph/generated"
	"git.9d77v.me/9d77v/pdc/graph/model"
)

func (r *mutationResolver) CreateMedia(ctx context.Context, input model.NewMedia) (int64, error) {
	return mediaService.CreateMedia(input)
}

func (r *mutationResolver) CreateEpisode(ctx context.Context, input model.NewEpisode) (int64, error) {
	return mediaService.CreateEpisode(input)
}

func (r *queryResolver) ListMedia(ctx context.Context) ([]*model.Media, error) {
	return mediaService.ListMedia()
}

func (r *queryResolver) PresignedURL(ctx context.Context, bucketName string, objectName string) (string, error) {
	return mediaService.PresignedURL(bucketName, objectName)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
