package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/9d77v/pdc/internal/consts"
	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/internal/graph/generated"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/middleware"
	"github.com/9d77v/pdc/internal/module/device-service/pb"
	"github.com/9d77v/pdc/internal/utils"
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
	return videoService.CreateVideo(ctx, input)
}

func (r *mutationResolver) AddVideoResource(ctx context.Context, input model.NewVideoResource) (*model.Video, error) {
	return videoService.AddVideoResource(ctx, input)
}

func (r *mutationResolver) SaveSubtitles(ctx context.Context, input model.NewSaveSubtitles) (*model.Video, error) {
	return videoService.SaveSubtitles(ctx, input)
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
	return thingService.UpdateThing(ctx, input, user.ID)
}

func (r *mutationResolver) RecordHistory(ctx context.Context, input model.NewHistoryInput) (*model.History, error) {
	user := middleware.ForContext(ctx)
	return historyService.RecordHistory(ctx, input, user.ID)
}

func (r *mutationResolver) CreateDeviceModel(ctx context.Context, input model.NewDeviceModel) (*model.DeviceModel, error) {
	resp, err := deviceModelService.CreateDeviceModel(ctx, getCreateDeviceModelRequest(input))
	return &model.DeviceModel{ID: resp.Id}, err
}

func (r *mutationResolver) UpdateDeviceModel(ctx context.Context, input model.NewUpdateDeviceModel) (*model.DeviceModel, error) {
	resp, err := deviceModelService.UpdateDeviceModel(ctx, getUpdateDeviceModelRequest(input))
	return &model.DeviceModel{ID: resp.Id}, err
}

func (r *mutationResolver) CreateAttributeModel(ctx context.Context, input model.NewAttributeModel) (*model.AttributeModel, error) {
	resp, err := deviceModelService.CreateAttributeModel(ctx, getCreateAttributeModelRequest(input))
	return &model.AttributeModel{ID: resp.Id}, err
}

func (r *mutationResolver) UpdateAttributeModel(ctx context.Context, input model.NewUpdateAttributeModel) (*model.AttributeModel, error) {
	resp, err := deviceModelService.UpdateAttributeModel(ctx, getUpdateAttributeModelRequest(input))
	return &model.AttributeModel{ID: resp.Id}, err
}

func (r *mutationResolver) DeleteAttributeModel(ctx context.Context, id int64) (*model.AttributeModel, error) {
	resp, err := deviceModelService.DeleteAttributeModel(ctx, &pb.DeleteAttributeModelRequest{Ids: []int64{id}})
	return &model.AttributeModel{ID: resp.Ids[0]}, err
}

func (r *mutationResolver) CreateTelemetryModel(ctx context.Context, input model.NewTelemetryModel) (*model.TelemetryModel, error) {
	resp, err := deviceModelService.CreateTelemetryModel(ctx, getCreateTelemetryModelRequest(input))
	return &model.TelemetryModel{ID: resp.Id}, err
}

func (r *mutationResolver) UpdateTelemetryModel(ctx context.Context, input model.NewUpdateTelemetryModel) (*model.TelemetryModel, error) {
	resp, err := deviceModelService.UpdateTelemetryModel(ctx, getUpdateTelemetryModelRequest(input))
	return &model.TelemetryModel{ID: resp.Id}, err
}

func (r *mutationResolver) DeleteTelemetryModel(ctx context.Context, id int64) (*model.TelemetryModel, error) {
	resp, err := deviceModelService.DeleteTelemetryModel(ctx, &pb.DeleteTelemetryModelRequest{Ids: []int64{id}})
	return &model.TelemetryModel{ID: resp.Ids[0]}, err
}

func (r *mutationResolver) CreateDevice(ctx context.Context, input model.NewDevice) (*model.Device, error) {
	resp, err := deviceService.CreateDevice(ctx, getCreateDeviceRequest(input))
	return &model.Device{ID: resp.Id}, err
}

func (r *mutationResolver) UpdateDevice(ctx context.Context, input model.NewUpdateDevice) (*model.Device, error) {
	resp, err := deviceService.UpdateDevice(ctx, getUpdateDeviceRequest(input))
	return &model.Device{ID: resp.Id}, err
}

func (r *mutationResolver) CreateDeviceDashboard(ctx context.Context, input model.NewDeviceDashboard) (*model.DeviceDashboard, error) {
	resp, err := deviceDashboardService.CreateDeviceDashboard(ctx, getCreateDeviceDashboardRequest(input))
	return &model.DeviceDashboard{ID: resp.Id}, err
}

func (r *mutationResolver) UpdateDeviceDashboard(ctx context.Context, input model.NewUpdateDeviceDashboard) (*model.DeviceDashboard, error) {
	resp, err := deviceDashboardService.UpdateDeviceDashboard(ctx, getUpdateDeviceDashboardRequest(input))
	return &model.DeviceDashboard{ID: resp.Id}, err
}

func (r *mutationResolver) DeleteDeviceDashboard(ctx context.Context, id int64) (*model.DeviceDashboard, error) {
	_, err := deviceDashboardService.DeleteDeviceDashboard(ctx, &pb.DeleteDeviceDashboardRequest{Ids: []int64{id}})
	return &model.DeviceDashboard{ID: 0}, err
}

func (r *mutationResolver) AddDeviceDashboardTelemetry(ctx context.Context, input model.NewDeviceDashboardTelemetry) (*model.DeviceDashboard, error) {
	_, err := deviceDashboardService.AddTelemetries(ctx, getAddTelemetriesRequest(input))
	return &model.DeviceDashboard{ID: 0}, err
}

func (r *mutationResolver) RemoveDeviceDashboardTelemetry(ctx context.Context, ids []int64) (*model.DeviceDashboard, error) {
	_, err := deviceDashboardService.RemoveTelemetries(ctx, &pb.RemoveTelemetriesRequest{Ids: ids})
	return &model.DeviceDashboard{ID: 0}, err
}

func (r *mutationResolver) AddDeviceDashboardCamera(ctx context.Context, input model.NewDeviceDashboardCamera) (*model.DeviceDashboard, error) {
	_, err := deviceDashboardService.AddCameras(ctx, getAddCamerasRequest(input))
	return &model.DeviceDashboard{ID: 0}, err
}

func (r *mutationResolver) RemoveDeviceDashboardCamera(ctx context.Context, ids []int64) (*model.DeviceDashboard, error) {
	_, err := deviceDashboardService.RemoveCameras(ctx, &pb.RemoveCamerasRequest{Ids: ids})
	return &model.DeviceDashboard{ID: 0}, err
}

func (r *mutationResolver) CameraCapture(ctx context.Context, deviceID int64) (string, error) {
	scheme := middleware.ForSchemeContext(ctx)
	resp, err := deviceService.CameraCapture(context.Background(),
		&pb.CameraCaptureRequest{DeviceId: uint32(deviceID), Scheme: scheme})
	return resp.ImageUrl, err
}

func (r *queryResolver) PresignedURL(ctx context.Context, bucketName string, objectName string) (string, error) {
	scheme := middleware.ForSchemeContext(ctx)
	return oss.GetPresignedURL(ctx, bucketName, objectName, scheme)
}

func (r *queryResolver) Users(ctx context.Context, searchParam model.SearchParam) (*model.UserConnection, error) {
	con := new(model.UserConnection)
	scheme := middleware.ForSchemeContext(ctx)
	total, data, err := userService.ListUser(ctx, getSearchParam(ctx, searchParam), scheme)
	con.TotalCount = total
	con.Edges = data
	return con, err
}

func (r *queryResolver) UserInfo(ctx context.Context, uid *int64) (*model.User, error) {
	scheme := middleware.ForSchemeContext(ctx)
	user := userService.GetUserInfo(middleware.ForContext(ctx), scheme)
	user.UID = consts.GetEncodeUID(uint(user.ID))
	user.ID = 0
	return user, nil
}

func (r *queryResolver) Videos(ctx context.Context, searchParam model.SearchParam, isFilterVideoSeries *bool, episodeID *int64) (*model.VideoConnection, error) {
	con := new(model.VideoConnection)
	scheme := middleware.ForSchemeContext(ctx)
	total, data, err := videoService.ListVideo(ctx, getSearchParam(ctx, searchParam), scheme, isFilterVideoSeries, episodeID)
	con.TotalCount = total
	con.Edges = data
	return con, err
}

func (r *queryResolver) VideoSerieses(ctx context.Context, searchParam model.SearchParam, episodeID *int64) (*model.VideoSeriesConnection, error) {
	con := new(model.VideoSeriesConnection)
	total, data, err := videoService.ListVideoSeries(ctx, getSearchParam(ctx, searchParam), episodeID)
	con.TotalCount = total
	con.Edges = data
	return con, err
}

func (r *queryResolver) SearchVideo(ctx context.Context, searchParam model.SearchParam) (*model.VideoIndexConnection, error) {
	scheme := middleware.ForSchemeContext(ctx)
	return videoService.SearchVideo(ctx, getSearchParam(ctx, searchParam), scheme)
}

func (r *queryResolver) SimilarVideos(ctx context.Context, searchParam model.SearchParam, episodeID int64) (*model.VideoIndexConnection, error) {
	scheme := middleware.ForSchemeContext(ctx)
	return videoService.SimilarVideos(ctx, getSearchParam(ctx, searchParam), episodeID, scheme)
}

func (r *queryResolver) Things(ctx context.Context, searchParam model.SearchParam) (*model.ThingConnection, error) {
	user := middleware.ForContext(ctx)
	scheme := middleware.ForSchemeContext(ctx)
	con := new(model.ThingConnection)
	total, data, err := thingService.ListThing(ctx, getSearchParam(ctx, searchParam), user.ID, scheme)
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

func (r *queryResolver) Histories(ctx context.Context, sourceType *int64, searchParam model.SearchParam, subSourceID *int64) (*model.HistoryConnection, error) {
	user := middleware.ForContext(ctx)
	scheme := middleware.ForSchemeContext(ctx)
	con := new(model.HistoryConnection)
	total, data, err := historyService.ListHistory(ctx, sourceType, getSearchParam(ctx, searchParam), subSourceID, user.ID, scheme)
	con.TotalCount = total
	con.Edges = data
	return con, err
}

func (r *queryResolver) DeviceModels(ctx context.Context, searchParam model.SearchParam) (*model.DeviceModelConnection, error) {
	resp, err := deviceModelService.ListDeviceModel(context.Background(),
		&pb.ListDeviceModelRequest{SearchParam: getSearchParam(ctx, searchParam)})
	return getDeviceModelConnection(resp), err
}

func (r *queryResolver) Devices(ctx context.Context, searchParam model.SearchParam, deviceType *int64) (*model.DeviceConnection, error) {
	resp, err := deviceService.ListDevice(context.Background(), &pb.ListDeviceRequest{
		SearchParam: getSearchParam(ctx, searchParam),
		DeviceType:  deviceType,
	})
	return getDeviceConnection(resp), err
}

func (r *queryResolver) DeviceDashboards(ctx context.Context, searchParam model.SearchParam) (*model.DeviceDashboardConnection, error) {
	resp, err := deviceDashboardService.ListDeviceDashboards(context.Background(), &pb.ListDeviceDashboardRequest{
		SearchParam: getSearchParam(ctx, searchParam),
	})
	return getDeviceDashboardConnection(resp), err
}

func (r *queryResolver) AppDeviceDashboards(ctx context.Context, deviceType *int64) (*model.DeviceDashboardConnection, error) {
	resp, err := deviceDashboardService.ListAppDeviceDashboards(context.Background(), &pb.ListAppDeviceDashboardRequest{
		DeviceType:  deviceType,
		QueryFields: utils.GetPreloads(ctx),
	})
	return getAppDeviceDashboardConnection(resp), err
}

func (r *queryResolver) CameraTimeLapseVideos(ctx context.Context, deviceID int64) (*model.CameraTimeLapseVideoConnection, error) {
	scheme := middleware.ForSchemeContext(ctx)
	resp, err := deviceDashboardService.ListCameraTimeLapseVideos(context.Background(),
		&pb.ListCameraTimeLapseVideoRequest{
			QueryFields: utils.GetPreloads(ctx),
			DeviceID:    deviceID,
			Scheme:      scheme,
		})
	return getCameraTimeLapseVideoConnection(resp), err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
