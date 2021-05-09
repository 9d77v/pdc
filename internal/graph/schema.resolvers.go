package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/consts"
	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/internal/graph/dtos/book_dto"
	"github.com/9d77v/pdc/internal/graph/dtos/common_dto"
	"github.com/9d77v/pdc/internal/graph/dtos/device_dto"
	"github.com/9d77v/pdc/internal/graph/dtos/note_dto"
	"github.com/9d77v/pdc/internal/graph/generated"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/middleware"
	bookPB "github.com/9d77v/pdc/internal/module/book-service/pb"
	devicePB "github.com/9d77v/pdc/internal/module/device-service/pb"
	notePB "github.com/9d77v/pdc/internal/module/note-service/pb"
	"github.com/9d77v/pdc/internal/utils"
	"github.com/golang/protobuf/ptypes"
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
	resp, err := deviceModelService.CreateDeviceModel(ctx, &devicePB.CreateDeviceModelRequest{
		Name:          input.Name,
		Desc:          ptrs.String(input.Desc),
		DeviceType:    devicePB.DeviceType(input.DeviceType),
		CameraCompany: devicePB.CameraCompany(input.CameraCompany),
	})
	return &model.DeviceModel{ID: resp.Id}, err
}

func (r *mutationResolver) UpdateDeviceModel(ctx context.Context, input model.NewUpdateDeviceModel) (*model.DeviceModel, error) {
	resp, err := deviceModelService.UpdateDeviceModel(ctx, &devicePB.UpdateDeviceModelRequest{
		Id:   input.ID,
		Name: input.Name,
		Desc: ptrs.String(input.Desc),
	})
	return &model.DeviceModel{ID: resp.Id}, err
}

func (r *mutationResolver) CreateAttributeModel(ctx context.Context, input model.NewAttributeModel) (*model.AttributeModel, error) {
	resp, err := deviceModelService.CreateAttributeModel(ctx, &devicePB.CreateAttributeModelRequest{
		DeviceModelId: input.DeviceModelID,
		Key:           input.Key,
		Name:          input.Name,
	})
	return &model.AttributeModel{ID: resp.Id}, err
}

func (r *mutationResolver) UpdateAttributeModel(ctx context.Context, input model.NewUpdateAttributeModel) (*model.AttributeModel, error) {
	resp, err := deviceModelService.UpdateAttributeModel(ctx, &devicePB.UpdateAttributeModelRequest{
		Id:   input.ID,
		Name: input.Name,
	})
	return &model.AttributeModel{ID: resp.Id}, err
}

func (r *mutationResolver) DeleteAttributeModel(ctx context.Context, id int64) (*model.AttributeModel, error) {
	resp, err := deviceModelService.DeleteAttributeModel(ctx, &devicePB.DeleteAttributeModelRequest{Ids: []int64{id}})
	return &model.AttributeModel{ID: resp.Ids[0]}, err
}

func (r *mutationResolver) CreateTelemetryModel(ctx context.Context, input model.NewTelemetryModel) (*model.TelemetryModel, error) {
	resp, err := deviceModelService.CreateTelemetryModel(ctx, &devicePB.CreateTelemetryModelRequest{
		DeviceModelId: input.DeviceModelID,
		Key:           input.Key,
		Name:          input.Name,
		Factor:        input.Factor,
		Unit:          ptrs.String(input.Unit),
		UnitName:      ptrs.String(input.UnitName),
		Scale:         input.Scale,
	})
	return &model.TelemetryModel{ID: resp.Id}, err
}

func (r *mutationResolver) UpdateTelemetryModel(ctx context.Context, input model.NewUpdateTelemetryModel) (*model.TelemetryModel, error) {
	resp, err := deviceModelService.UpdateTelemetryModel(ctx, &devicePB.UpdateTelemetryModelRequest{
		Id:       input.ID,
		Name:     input.Name,
		Factor:   input.Factor,
		Unit:     input.Unit,
		UnitName: input.UnitName,
		Scale:    input.Scale,
	})
	return &model.TelemetryModel{ID: resp.Id}, err
}

func (r *mutationResolver) DeleteTelemetryModel(ctx context.Context, id int64) (*model.TelemetryModel, error) {
	resp, err := deviceModelService.DeleteTelemetryModel(ctx, &devicePB.DeleteTelemetryModelRequest{Ids: []int64{id}})
	return &model.TelemetryModel{ID: resp.Ids[0]}, err
}

func (r *mutationResolver) CreateDevice(ctx context.Context, input model.NewDevice) (*model.Device, error) {
	resp, err := deviceService.CreateDevice(ctx, &devicePB.CreateDeviceRequest{
		DeviceModelId: input.DeviceModelID,
		Name:          input.Name,
		Ip:            ptrs.String(input.IP),
		Port:          uint64(ptrs.Int64(input.Port)),
		Username:      ptrs.String(input.Username),
		Password:      ptrs.String(input.Password),
	})
	return &model.Device{ID: resp.Id}, err
}

func (r *mutationResolver) UpdateDevice(ctx context.Context, input model.NewUpdateDevice) (*model.Device, error) {
	resp, err := deviceService.UpdateDevice(ctx, &devicePB.UpdateDeviceRequest{
		Id:       input.ID,
		Name:     input.Name,
		Ip:       ptrs.String(input.IP),
		Port:     uint64(ptrs.Int64(input.Port)),
		Username: ptrs.String(input.Username),
		Password: ptrs.String(input.Password),
	})
	return &model.Device{ID: resp.Id}, err
}

func (r *mutationResolver) CreateDeviceDashboard(ctx context.Context, input model.NewDeviceDashboard) (*model.DeviceDashboard, error) {
	resp, err := deviceDashboardService.CreateDeviceDashboard(ctx, &devicePB.CreateDeviceDashboardRequest{
		Name:       input.Name,
		DeviceType: devicePB.DeviceType(input.DeviceType),
		IsVisible:  input.IsVisible,
	})
	return &model.DeviceDashboard{ID: resp.Id}, err
}

func (r *mutationResolver) UpdateDeviceDashboard(ctx context.Context, input model.NewUpdateDeviceDashboard) (*model.DeviceDashboard, error) {
	resp, err := deviceDashboardService.UpdateDeviceDashboard(ctx, &devicePB.UpdateDeviceDashboardRequest{
		Id:        input.ID,
		Name:      input.Name,
		IsVisible: input.IsVisible,
	})
	return &model.DeviceDashboard{ID: resp.Id}, err
}

func (r *mutationResolver) DeleteDeviceDashboard(ctx context.Context, id int64) (*model.DeviceDashboard, error) {
	_, err := deviceDashboardService.DeleteDeviceDashboard(ctx, &devicePB.DeleteDeviceDashboardRequest{Ids: []int64{id}})
	return &model.DeviceDashboard{ID: 0}, err
}

func (r *mutationResolver) AddDeviceDashboardTelemetry(ctx context.Context, input model.NewDeviceDashboardTelemetry) (*model.DeviceDashboard, error) {
	_, err := deviceDashboardService.AddTelemetries(ctx, &devicePB.AddTelemetriesRequest{
		DeviceDashboardId: input.DeviceDashboardID,
		TelemetryIds:      input.TelemetryIDs,
	})
	return &model.DeviceDashboard{ID: 0}, err
}

func (r *mutationResolver) RemoveDeviceDashboardTelemetry(ctx context.Context, ids []int64) (*model.DeviceDashboard, error) {
	_, err := deviceDashboardService.RemoveTelemetries(ctx, &devicePB.RemoveTelemetriesRequest{Ids: ids})
	return &model.DeviceDashboard{ID: 0}, err
}

func (r *mutationResolver) AddDeviceDashboardCamera(ctx context.Context, input model.NewDeviceDashboardCamera) (*model.DeviceDashboard, error) {
	_, err := deviceDashboardService.AddCameras(ctx, &devicePB.AddCamerasRequest{
		DeviceDashboardId: input.DeviceDashboardID,
		DeviceIds:         input.DeviceIDs,
	})
	return &model.DeviceDashboard{ID: 0}, err
}

func (r *mutationResolver) RemoveDeviceDashboardCamera(ctx context.Context, ids []int64) (*model.DeviceDashboard, error) {
	_, err := deviceDashboardService.RemoveCameras(ctx, &devicePB.RemoveCamerasRequest{Ids: ids})
	return &model.DeviceDashboard{ID: 0}, err
}

func (r *mutationResolver) CameraCapture(ctx context.Context, deviceID int64) (string, error) {
	scheme := middleware.ForSchemeContext(ctx)
	resp, err := deviceService.CameraCapture(context.Background(),
		&devicePB.CameraCaptureRequest{DeviceId: uint32(deviceID), Scheme: scheme})
	return resp.ImageUrl, err
}

func (r *mutationResolver) SyncNotes(ctx context.Context, input model.SyncNotesInput) (*model.SyncNotesResponse, error) {
	user := middleware.ForContext(ctx)
	t, _ := ptypes.TimestampProto(time.Unix(input.LastUpdateTime, 0))
	resp, err := noteService.SyncNotes(context.Background(), &notePB.SyncNotesRequest{
		Uid:            int64(user.ID),
		LastUpdateTime: t,
		UnsyncedNotes:  note_dto.GetNotes(input.UnsyncedNotes),
	})
	return note_dto.GetSyncNotesResponseConnection(resp), err
}

func (r *mutationResolver) CreateBookShelf(ctx context.Context, input model.NewBookShelf) (*model.BookShelf, error) {
	resp, err := bookService.CreateBookShelf(ctx, &bookPB.CreateBookShelfRequest{
		Name:         input.Name,
		LayerNum:     int32(input.LayerNum),
		PartitionNum: int32(input.PartitionNum),
	})
	return &model.BookShelf{ID: resp.GetId()}, err
}

func (r *mutationResolver) UpdateBookShelf(ctx context.Context, input model.NewUpdateBookShelf) (*model.BookShelf, error) {
	resp, err := bookService.UpdateBookShelf(ctx, &bookPB.UpdateBookShelfRequest{
		Id:   input.ID,
		Name: input.Name,
	})
	return &model.BookShelf{ID: resp.GetId()}, err
}

func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	t, _ := ptypes.TimestampProto(time.Unix(input.PurchaseTime, 0))
	resp, err := bookService.CreateBook(ctx, &bookPB.CreateBookRequest{
		Isbn:            input.Isbn,
		Name:            input.Name,
		Desc:            input.Desc,
		Cover:           input.Cover,
		Author:          input.Author,
		Translator:      input.Translator,
		PublishingHouse: input.PublishingHouse,
		Edition:         input.Edition,
		PrintedTimes:    input.PrintedTimes,
		PrintedSheets:   input.PrintedSheets,
		Format:          input.Format,
		WordCount:       input.WordCount,
		Pricing:         input.Pricing,
		PurchasePrice:   input.PurchasePrice,
		PurchaseTime:    t,
		PurchaseSource:  input.PurchaseSource,
	})
	return &model.Book{ID: resp.GetId()}, err
}

func (r *mutationResolver) UpdateBook(ctx context.Context, input model.NewUpdateBook) (*model.Book, error) {
	t, _ := ptypes.TimestampProto(time.Unix(input.PurchaseTime, 0))
	resp, err := bookService.UpdateBook(ctx, &bookPB.UpdateBookRequest{
		Id:              input.ID,
		Name:            input.Name,
		Desc:            input.Desc,
		Cover:           input.Cover,
		Author:          input.Author,
		Translator:      input.Translator,
		PublishingHouse: input.PublishingHouse,
		Edition:         input.Edition,
		PrintedTimes:    input.PrintedTimes,
		PrintedSheets:   input.PrintedSheets,
		Format:          input.Format,
		WordCount:       input.WordCount,
		Pricing:         input.Pricing,
		PurchasePrice:   input.PurchasePrice,
		PurchaseTime:    t,
		PurchaseSource:  input.PurchaseSource,
	})
	return &model.Book{ID: resp.GetId()}, err
}

func (r *mutationResolver) CreateBookPosition(ctx context.Context, input model.NewBookPosition) (*model.BookPosition, error) {
	resp, err := bookService.CreateBookPosition(ctx, &bookPB.CreateBookPositionRequest{
		BookShelfId: input.BookShelfID,
		BookId:      input.BookID,
		Layer:       int32(input.Layer),
		Partition:   int32(input.Partition),
	})
	return &model.BookPosition{ID: resp.GetId()}, err
}

func (r *mutationResolver) UpdateBookPosition(ctx context.Context, input model.NewUpdateBookPosition) (*model.BookPosition, error) {
	resp, err := bookService.UpdateBookPosition(ctx, &bookPB.UpdateBookPositionRequest{
		Id:          input.ID,
		BookShelfId: input.BookShelfID,
		Layer:       int32(input.Layer),
		Partition:   int32(input.Partition),
	})
	return &model.BookPosition{ID: resp.GetId()}, err
}

func (r *mutationResolver) BorrowBook(ctx context.Context, bookID int64, uid int64) (*model.BookBorrowReturn, error) {
	resp, err := bookService.BorrowBook(ctx, &bookPB.BorrowBookRequest{
		BookId: bookID,
		Uid:    uid,
	})
	return &model.BookBorrowReturn{ID: resp.Id}, err
}

func (r *mutationResolver) BackBook(ctx context.Context, bookID int64, uid int64) (*model.BookBorrowReturn, error) {
	resp, err := bookService.ReturnBook(ctx, &bookPB.ReturnBookRequest{
		BookId: bookID,
		Uid:    uid,
	})
	return &model.BookBorrowReturn{ID: resp.Id}, err
}

func (r *queryResolver) PresignedURL(ctx context.Context, bucketName string, objectName string) (*model.PresignedURLResponse, error) {
	scheme := middleware.ForSchemeContext(ctx)
	ok, url, err := oss.GetPresignedURL(ctx, bucketName, objectName, scheme)
	resp := &model.PresignedURLResponse{
		Ok:  ok,
		URL: url,
	}
	return resp, err
}

func (r *queryResolver) Users(ctx context.Context, searchParam model.SearchParam) (*model.UserConnection, error) {
	con := new(model.UserConnection)
	scheme := middleware.ForSchemeContext(ctx)
	total, data, err := userService.ListUser(ctx, common_dto.GetSearchParam(ctx, searchParam), scheme)
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
	total, data, err := videoService.ListVideo(ctx, common_dto.GetSearchParam(ctx, searchParam), scheme, isFilterVideoSeries, episodeID)
	con.TotalCount = total
	con.Edges = data
	return con, err
}

func (r *queryResolver) VideoSerieses(ctx context.Context, searchParam model.SearchParam, episodeID *int64) (*model.VideoSeriesConnection, error) {
	con := new(model.VideoSeriesConnection)
	total, data, err := videoService.ListVideoSeries(ctx, common_dto.GetSearchParam(ctx, searchParam), episodeID)
	con.TotalCount = total
	con.Edges = data
	return con, err
}

func (r *queryResolver) SearchVideo(ctx context.Context, searchParam model.SearchParam) (*model.VideoIndexConnection, error) {
	scheme := middleware.ForSchemeContext(ctx)
	return videoService.SearchVideo(ctx, common_dto.GetSearchParam(ctx, searchParam), scheme)
}

func (r *queryResolver) SimilarVideos(ctx context.Context, searchParam model.SearchParam, episodeID int64) (*model.VideoIndexConnection, error) {
	scheme := middleware.ForSchemeContext(ctx)
	return videoService.SimilarVideos(ctx, common_dto.GetSearchParam(ctx, searchParam), episodeID, scheme)
}

func (r *queryResolver) Things(ctx context.Context, searchParam model.SearchParam) (*model.ThingConnection, error) {
	user := middleware.ForContext(ctx)
	scheme := middleware.ForSchemeContext(ctx)
	con := new(model.ThingConnection)
	total, data, err := thingService.ListThing(ctx, common_dto.GetSearchParam(ctx, searchParam), user.ID, scheme)
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
	total, data, err := historyService.ListHistory(ctx, sourceType, common_dto.GetSearchParam(ctx, searchParam), subSourceID, user.ID, scheme)
	con.TotalCount = total
	con.Edges = data
	return con, err
}

func (r *queryResolver) HistoryStatistic(ctx context.Context, sourceType *int64) (*model.HistoryStatistic, error) {
	return historyService.HistoryStatistic(ctx, sourceType)
}

func (r *queryResolver) AppHistoryStatistic(ctx context.Context, sourceType *int64) (*model.HistoryStatistic, error) {
	user := middleware.ForContext(ctx)
	return historyService.AppHistoryStatistic(ctx, sourceType, user.ID)
}

func (r *queryResolver) DeviceModels(ctx context.Context, searchParam model.SearchParam) (*model.DeviceModelConnection, error) {
	resp, err := deviceModelService.ListDeviceModel(context.Background(),
		&devicePB.ListDeviceModelRequest{SearchParam: common_dto.GetSearchParam(ctx, searchParam)})
	return device_dto.GetDeviceModelConnection(resp), err
}

func (r *queryResolver) Devices(ctx context.Context, searchParam model.SearchParam, deviceType *int64) (*model.DeviceConnection, error) {
	resp, err := deviceService.ListDevice(context.Background(), &devicePB.ListDeviceRequest{
		SearchParam: common_dto.GetSearchParam(ctx, searchParam),
		DeviceType:  deviceType,
	})
	return device_dto.GetDeviceConnection(resp), err
}

func (r *queryResolver) DeviceDashboards(ctx context.Context, searchParam model.SearchParam) (*model.DeviceDashboardConnection, error) {
	resp, err := deviceDashboardService.ListDeviceDashboards(context.Background(), &devicePB.ListDeviceDashboardRequest{
		SearchParam: common_dto.GetSearchParam(ctx, searchParam),
	})
	return device_dto.GetDeviceDashboardConnection(resp), err
}

func (r *queryResolver) AppDeviceDashboards(ctx context.Context, deviceType *int64) (*model.DeviceDashboardConnection, error) {
	resp, err := deviceDashboardService.ListAppDeviceDashboards(context.Background(), &devicePB.ListAppDeviceDashboardRequest{
		DeviceType:  deviceType,
		QueryFields: utils.GetPreloads(ctx),
	})
	return device_dto.GetAppDeviceDashboardConnection(resp), err
}

func (r *queryResolver) CameraTimeLapseVideos(ctx context.Context, deviceID int64) (*model.CameraTimeLapseVideoConnection, error) {
	scheme := middleware.ForSchemeContext(ctx)
	resp, err := deviceDashboardService.ListCameraTimeLapseVideos(context.Background(),
		&devicePB.ListCameraTimeLapseVideoRequest{
			QueryFields: utils.GetPreloads(ctx),
			DeviceID:    deviceID,
			Scheme:      scheme,
		})
	return device_dto.GetCameraTimeLapseVideoConnection(resp), err
}

func (r *queryResolver) Books(ctx context.Context, searchParam model.SearchParam) (*model.BookConnection, error) {
	resp, err := bookService.ListBook(context.Background(), &bookPB.ListBookRequest{
		SearchParam: common_dto.GetSearchParam(ctx, searchParam),
	})
	scheme := middleware.ForSchemeContext(ctx)
	return book_dto.GetBookConnection(resp, scheme), err
}

func (r *queryResolver) BookShelfs(ctx context.Context, searchParam model.SearchParam) (*model.BookShelfConnection, error) {
	resp, err := bookService.ListBookShelf(context.Background(), &bookPB.ListBookShelfRequest{
		SearchParam: common_dto.GetSearchParam(ctx, searchParam),
	})
	scheme := middleware.ForSchemeContext(ctx)
	return book_dto.GetBookShelfConnection(resp, scheme), err
}

func (r *queryResolver) BookPositions(ctx context.Context, searchParam model.SearchParam, bookID *int64, bookShelfID *int64) (*model.BookPositionConnection, error) {
	resp, err := bookService.ListBookPosition(context.Background(), &bookPB.ListBookPositionRequest{
		SearchParam: common_dto.GetSearchParam(ctx, searchParam),
		BookID:      bookID,
		BookShelfID: bookShelfID,
	})
	return book_dto.GetBookPositionConnection(resp), err
}

func (r *queryResolver) BookBorrowReturn(ctx context.Context, searchParam model.SearchParam, bookID *int64) (*model.BookBorrowReturnConnection, error) {
	resp, err := bookService.ListBookBorrowReturn(context.Background(), &bookPB.ListBookBorrowReturnRequest{
		SearchParam: common_dto.GetSearchParam(ctx, searchParam),
		BookID:      bookID,
	})
	return book_dto.GetBookBorrowReturnConnection(resp), err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
