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

func (r *mutationResolver) UpdateProfile(ctx context.Context, input model.NewUpdateProfile) (*model.User, error) {
	user := middleware.ForContext(ctx)
	return userService.UpdateProfile(ctx, input, user.ID)
}

func (r *mutationResolver) UpdatePassword(ctx context.Context, oldPassword string, newPassword string) (*model.User, error) {
	user := middleware.ForContext(ctx)
	return userService.UpdatePassword(ctx, oldPassword, newPassword, user.ID)
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
	return videoService.CreateVideo(ctx, input)
}

func (r *mutationResolver) UpdateVideo(ctx context.Context, input model.NewUpdateVideo) (*model.Video, error) {
	return videoService.UpdateVideo(ctx, input)
}

func (r *mutationResolver) CreateEpisode(ctx context.Context, input model.NewEpisode) (*model.Episode, error) {
	return videoService.CreateEpisode(ctx, input)
}

func (r *mutationResolver) UpdateEpisode(ctx context.Context, input model.NewUpdateEpisode) (*model.Episode, error) {
	return videoService.UpdateEpisode(ctx, input)
}

func (r *mutationResolver) UpdateSubtitle(ctx context.Context, input model.NewUpdateSubtitles) (*model.Video, error) {
	return videoService.UpdateSubtitle(ctx, input)
}

func (r *mutationResolver) UpdateMobileVideo(ctx context.Context, input *model.NewUpdateMobileVideos) (*model.Video, error) {
	return videoService.UpdateMobileVideo(ctx, input)
}

func (r *mutationResolver) CreateVideoSeries(ctx context.Context, input model.NewVideoSeries) (*model.VideoSeries, error) {
	return videoService.CreateVideoSeries(ctx, input)
}

func (r *mutationResolver) UpdateVideoSeries(ctx context.Context, input model.NewUpdateVideoSeries) (*model.VideoSeries, error) {
	return videoService.UpdateVideoSeries(ctx, input)
}

func (r *mutationResolver) CreateVideoSeriesItem(ctx context.Context, input model.NewVideoSeriesItem) (*model.VideoSeriesItem, error) {
	return videoService.CreateVideoSeriesItem(ctx, input)
}

func (r *mutationResolver) UpdateVideoSeriesItem(ctx context.Context, input model.NewUpdateVideoSeriesItem) (*model.VideoSeriesItem, error) {
	return videoService.UpdateVideoSeriesItem(ctx, input)
}

func (r *mutationResolver) CreateThing(ctx context.Context, input model.NewThing) (*model.Thing, error) {
	user := middleware.ForContext(ctx)
	return thingService.CreateThing(ctx, input, int64(user.ID))
}

func (r *mutationResolver) UpdateThing(ctx context.Context, input model.NewUpdateThing) (*model.Thing, error) {
	user := middleware.ForContext(ctx)
	return thingService.UpdateThing(ctx, input, int64(user.ID))
}

func (r *mutationResolver) RecordHistory(ctx context.Context, input model.NewHistoryInput) (*model.History, error) {
	user := middleware.ForContext(ctx)
	return historyService.RecordHistory(ctx, input, user.ID)
}

func (r *mutationResolver) CreateDeviceModel(ctx context.Context, input model.NewDeviceModel) (*model.DeviceModel, error) {
	return deviceService.CreateDeviceModel(ctx, input)
}

func (r *mutationResolver) UpdateDeviceModel(ctx context.Context, input model.NewUpdateDeviceModel) (*model.DeviceModel, error) {
	return deviceService.UpdateDeviceModel(ctx, input)
}

func (r *mutationResolver) CreateAttributeModel(ctx context.Context, input model.NewAttributeModel) (*model.AttributeModel, error) {
	return deviceService.CreateAttributeModel(ctx, input)
}

func (r *mutationResolver) UpdateAttributeModel(ctx context.Context, input model.NewUpdateAttributeModel) (*model.AttributeModel, error) {
	return deviceService.UpdateAttributeModel(ctx, input)
}

func (r *mutationResolver) CreateTelemetryModel(ctx context.Context, input model.NewTelemetryModel) (*model.TelemetryModel, error) {
	return deviceService.CreateTelemetryModel(ctx, input)
}

func (r *mutationResolver) UpdateTelemetryModel(ctx context.Context, input model.NewUpdateTelemetryModel) (*model.TelemetryModel, error) {
	return deviceService.UpdateTelemetryModel(ctx, input)
}

func (r *mutationResolver) CreateDevice(ctx context.Context, input model.NewDevice) (*model.Device, error) {
	return deviceService.CreateDevice(ctx, input)
}

func (r *mutationResolver) UpdateDevice(ctx context.Context, input model.NewUpdateDevice) (*model.Device, error) {
	return deviceService.UpdateDevice(ctx, input)
}

func (r *queryResolver) PresignedURL(ctx context.Context, bucketName string, objectName string) (string, error) {
	scheme := middleware.ForSchemeContext(ctx)
	return commonService.PresignedURL(ctx, scheme, bucketName, objectName)
}

func (r *queryResolver) Users(ctx context.Context, keyword *string, page *int64, pageSize *int64, ids []int64, sorts []*model.Sort) (*model.UserConnection, error) {
	con := new(model.UserConnection)
	scheme := middleware.ForSchemeContext(ctx)
	total, data, err := userService.ListUser(ctx, keyword,
		page, pageSize, ids, sorts, scheme)
	con.TotalCount = total
	con.Edges = data
	return con, err
}

func (r *queryResolver) UserInfo(ctx context.Context, uid *int64) (*model.User, error) {
	scheme := middleware.ForSchemeContext(ctx)
	return dtos.ToUserDto(middleware.ForContext(ctx), scheme), nil
}

func (r *queryResolver) Videos(ctx context.Context, keyword *string, page *int64, pageSize *int64, ids []int64, sorts []*model.Sort, isFilterVideoSeries *bool) (*model.VideoConnection, error) {
	con := new(model.VideoConnection)
	scheme := middleware.ForSchemeContext(ctx)
	total, data, err := videoService.ListVideo(ctx, keyword,
		page, pageSize, ids, sorts, scheme, isFilterVideoSeries)
	con.TotalCount = total
	con.Edges = data
	return con, err
}

func (r *queryResolver) VideoSerieses(ctx context.Context, keyword *string, videoID *int64, page *int64, pageSize *int64, ids []int64, sorts []*model.Sort) (*model.VideoSeriesConnection, error) {
	con := new(model.VideoSeriesConnection)
	total, data, err := videoService.ListVideoSeries(ctx, keyword, videoID,
		page, pageSize, ids, sorts)
	con.TotalCount = total
	con.Edges = data
	return con, err
}

func (r *queryResolver) SearchVideo(ctx context.Context, keyword *string, tags []string, page *int64, pageSize *int64) (*model.VideoIndexConnection, error) {
	con := new(model.VideoIndexConnection)
	scheme := middleware.ForSchemeContext(ctx)
	total, data, aggResults, err := videoService.ListVideoIndex(ctx, keyword, tags,
		page, pageSize, scheme)
	con.TotalCount = total
	con.Edges = data
	con.AggResults = aggResults
	return con, err
}

func (r *queryResolver) Things(ctx context.Context, keyword *string, page *int64, pageSize *int64, ids []int64, sorts []*model.Sort) (*model.ThingConnection, error) {
	user := middleware.ForContext(ctx)
	scheme := middleware.ForSchemeContext(ctx)
	con := new(model.ThingConnection)
	total, data, err := thingService.ListThing(ctx, keyword,
		page, pageSize, ids, sorts, int64(user.ID), scheme)
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

func (r *queryResolver) HistoryInfo(ctx context.Context, sourceType int64, sourceID int64) (*model.History, error) {
	user := middleware.ForContext(ctx)
	return historyService.GetHistory(ctx, sourceType, sourceID, user.ID)
}

func (r *queryResolver) Histories(ctx context.Context, sourceType *int64, page *int64, pageSize *int64) (*model.HistoryConnection, error) {
	user := middleware.ForContext(ctx)
	scheme := middleware.ForSchemeContext(ctx)
	con := new(model.HistoryConnection)
	total, data, err := historyService.ListHistory(ctx, sourceType, page, pageSize, int64(user.ID), scheme)
	con.TotalCount = total
	con.Edges = data
	return con, err
}

func (r *queryResolver) DeviceModels(ctx context.Context, keyword *string, page *int64, pageSize *int64, ids []int64, sorts []*model.Sort) (*model.DeviceModelConnection, error) {
	con := new(model.DeviceModelConnection)
	total, data, err := deviceService.ListDeviceModel(ctx, keyword, page, pageSize, ids, sorts)
	con.TotalCount = total
	con.Edges = data
	return con, err
}

func (r *queryResolver) Devices(ctx context.Context, keyword *string, page *int64, pageSize *int64, ids []int64, sorts []*model.Sort) (*model.DeviceConnection, error) {
	con := new(model.DeviceConnection)
	total, data, err := deviceService.ListDevice(ctx, keyword, page, pageSize, ids, sorts)
	con.TotalCount = total
	con.Edges = data
	return con, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
