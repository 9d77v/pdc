package graph

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/graph/generated"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/middleware"
	"github.com/9d77v/pdc/internal/module/video-service/models"
	"github.com/stretchr/testify/require"
)

var cc = client.New(middleware.Auth()(apiHandler))

var apiHandler = handler.NewDefaultServer(
	generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &Resolver{},
		},
	))
var loginResponse struct {
	Login model.LoginResponse
}

func post(query string, response interface{}, operation string, etcOptions ...client.Option) {
	options := []client.Option{
		client.Path("/api"),
		client.Operation(operation),
		client.AddHeader("Authorization", loginResponse.Login.AccessToken),
	}
	cc.MustPost(query, &response, append(options, etcOptions...)...)
}

func postInput(query string, response interface{}, operation string, input interface{}) {
	post(query, response, operation, client.Var("input", input))
}

func postSearch(query string, response interface{}, operation string, searchParam interface{},
	etcOptions ...client.Option) {
	options := []client.Option{client.Var("searchParam", searchParam)}
	post(query, response, operation, append(options, etcOptions...)...)
}

func TestMain(m *testing.M) {
	login()
	m.Run()
}

func login() {
	query := `
	mutation login($username: String!, $password: String!) {
		login(username: $username, password: $password) {
		  accessToken
		  refreshToken
		}
	  }
	  `
	cc.MustPost(query,
		&loginResponse,
		client.Var("username", "admin"),
		client.Var("password", "1234567890"),
	)
	fmt.Println("accessToken is ", loginResponse.Login.AccessToken)
}

func Test_mutationResolver_RecordHistory(t *testing.T) {
	var resp struct {
		RecordHistory model.History
	}
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		query string
		input model.NewHistoryInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test RecordHistory", args{`
			mutation recordHistory($input: NewHistoryInput!) {
				recordHistory(input: $input) {
				  subSourceID
				}
			  }
			`,
			model.NewHistoryInput{
				SourceType:    1,
				SourceID:      1,
				SubSourceID:   1,
				Platform:      "desktop",
				CurrentTime:   3.192658,
				RemainingTime: 1413.8713420000001,
			}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postInput(tt.args.query, &resp, "recordHistory", tt.args.input)
			require.NotZero(t, resp.RecordHistory.SubSourceID)
		})
	}
}

func Test_queryResolver_Histories(t *testing.T) {
	var resp struct {
		Histories model.HistoryConnection
	}
	type args struct {
		query       string
		sourceType  *int64
		searchParam model.SearchParam
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test histories", args{`
			query histories($sourceType: Int, $searchParam: SearchParam!) {
				histories(sourceType: $sourceType, searchParam: $searchParam) {
				  totalCount
				  edges {
					sourceType
					sourceID
					title
					num
					subTitle
					cover
					subSourceID
					platform
					currentTime
					remainingTime
					updatedAt
				  }
				}
			  }
			`, ptrs.Int64Ptr(1), model.SearchParam{Page: ptrs.Int64Ptr(1)}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postSearch(tt.args.query, &resp, "histories", tt.args.searchParam,
				client.Var("sourceType", tt.args.sourceType))
			require.NotZero(t, resp.Histories.TotalCount)
		})
	}
}

func Test_mutationResolver_CreateVideo(t *testing.T) {
	var resp struct {
		CreateVideo model.Video
	}
	type args struct {
		query string
		input model.NewVideo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test CreateVideo", args{`
			mutation createVideo($input:NewVideo!){
				createVideo(input:$input){
				  id
				}
			 }
			`, model.NewVideo{
			IsHideOnMobile: false,
			Title:          "title",
			Desc:           ptrs.StringPtr("desc"),
			PubDate:        ptrs.Int64Ptr(0),
			IsShow:         true,
			Theme:          "",
			Cover:          ptrs.StringPtr("cover"),
			Tags:           []string{"a", "b"},
		}}, false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postInput(tt.args.query, &resp, "createVideo", tt.args.input)
			require.NotZero(t, resp.CreateVideo.ID)
			removeVideo(t, uint(resp.CreateVideo.ID))
		})
	}
}

func removeVideo(t *testing.T, videoID uint) {
	video := new(models.Video)
	video.ID = videoID
	err := db.GetDB().Unscoped().Delete(video).Error
	if err != nil {
		t.Error("remove video failed:", err)
	}
}

func Test_mutationResolver_UpdateVideo(t *testing.T) {
	var resp struct {
		UpdateVideo model.Video
	}
	type args struct {
		query string
		input model.NewUpdateVideo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test UpdateVideo", args{`
			mutation updateVideo($input:NewUpdateVideo!){
			updateVideo(input:$input){
			  id
			}
		 }
		`, model.NewUpdateVideo{
			ID:             1,
			IsHideOnMobile: ptrs.BoolPtr(false),
			Title:          ptrs.StringPtr("title"),
			Desc:           ptrs.StringPtr("desc"),
			PubDate:        ptrs.Int64Ptr(0),
			IsShow:         ptrs.BoolPtr(true),
			Theme:          "",
			Cover:          ptrs.StringPtr("cover"),
			Tags:           []string{"a", "b"},
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postInput(tt.args.query, &resp, "updateVideo", tt.args.input)
			require.NotZero(t, resp.UpdateVideo.ID)
		})
	}
}

func Test_queryResolver_DeviceModels(t *testing.T) {
	var resp struct {
		DeviceModels model.DeviceModelConnection
	}
	type args struct {
		query       string
		searchParam model.SearchParam
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test DeviceModels", args{`
		query deviceModels($searchParam:SearchParam!) {
			deviceModels(searchParam: $searchParam) {
			  edges {
				id
				name
				desc
				deviceType
				cameraCompany
				attributeModels{
				  id
				  key
				  name
				  createdAt
				  updatedAt
				}
				telemetryModels{
				  id
				  key
				  name
				  factor
				  unit
				  unitName
				  scale
				  createdAt
				  updatedAt
				}
				createdAt
				updatedAt
			  }
			}
		  }
		`, model.SearchParam{Page: ptrs.Int64Ptr(1)}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postSearch(tt.args.query, &resp, "deviceModels", tt.args.searchParam)
			require.Zero(t, resp.DeviceModels.TotalCount)
			require.NotZero(t, len(resp.DeviceModels.Edges))
		})
	}
}

func Test_queryResolver_Devices(t *testing.T) {
	var resp struct {
		Devices model.DeviceConnection
	}
	type args struct {
		query       string
		searchParam model.SearchParam
		deviceType  *int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test Devices without deviceType", args{`
		query devices( $searchParam:SearchParam!) {
			devices(searchParam: $searchParam) {
			  edges {
				id
				name
				ip
				port
				accessKey
				secretKey
				username
				password
				deviceModelID
				deviceModelName
				deviceModelDesc
				deviceModelDeviceType
				deviceModelCameraCompany
				attributes{
				  id
				  key
				  name
				  value
				  createdAt
				  updatedAt
				}
				telemetries{
				  id
				  key
				  name
				  value
				  unit
				  unitName
				  factor
				  scale
				  createdAt
				  updatedAt
				}
				createdAt
				updatedAt
			  }
			}
		  }
		`, model.SearchParam{Page: ptrs.Int64Ptr(1)}, nil}, false},
		{"test Devices  with deviceType", args{`
		query devices($deviceType:Int!, $searchParam:SearchParam!) {
			devices(deviceType:$deviceType,searchParam: $searchParam) {
			  edges {
				id
				name
				telemetries{
				  id
				  key
				  name
				}
			  }
			}
		  }
		`, model.SearchParam{Page: ptrs.Int64Ptr(1)}, ptrs.Int64Ptr(1)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			options := []client.Option{}
			if tt.args.deviceType != nil {
				options = append(options, client.Var("deviceType", tt.args.deviceType))
			}
			postSearch(tt.args.query, &resp, "devices", tt.args.searchParam, options...)
			require.Zero(t, resp.Devices.TotalCount)
			require.NotZero(t, len(resp.Devices.Edges))
		})
	}
}

func Test_queryResolver_DeviceDashboards(t *testing.T) {
	var resp struct {
		DeviceDashboards model.DeviceDashboardConnection
	}
	type args struct {
		query       string
		searchParam model.SearchParam
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test DeviceDashboards", args{`
		query deviceDashboards($searchParam:SearchParam!) {
			deviceDashboards(searchParam: $searchParam) {
			  totalCount
			  edges {
				id
				name
				isVisible
				deviceType
				telemetries{
				  id
				  deviceID
				  deviceName
				  telemetryID
				  name
				  value
				  factor
				  scale
				  unit
				}
				cameras{
				  id
				  deviceID
				  deviceName
				}
			  }
			}
		  }
		`, model.SearchParam{Page: ptrs.Int64Ptr(1)}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postSearch(tt.args.query, &resp, "deviceDashboards", tt.args.searchParam)
			require.NotZero(t, resp.DeviceDashboards.TotalCount)
		})
	}
}

func Test_queryResolver_Things(t *testing.T) {
	var resp struct {
		Things model.ThingConnection
	}
	type args struct {
		query       string
		searchParam model.SearchParam
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test Things", args{`
		query things( $searchParam:SearchParam!) {
			things(searchParam: $searchParam){
				 totalCount
				 edges{
					 id
					 uid
					 name
					 num
					 brandName
					 pics
					 unitPrice
					 unit
					 specifications
					 category
					 consumerExpenditure
					 location
					 status
					 purchaseDate
					 purchasePlatform
					 refOrderID
					 rubbishCategory
					 createdAt
					 updatedAt
				}
			}
		   }
		`, model.SearchParam{Page: ptrs.Int64Ptr(1)}}, false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postSearch(tt.args.query, &resp, "things", tt.args.searchParam)
			require.NotZero(t, resp.Things.TotalCount)
		})
	}
}

func Test_queryResolver_Users(t *testing.T) {
	var resp struct {
		Users model.UserConnection
	}
	type args struct {
		query       string
		searchParam model.SearchParam
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test Users", args{`
		query users($searchParam:SearchParam!) {
			users(searchParam: $searchParam){
				 totalCount
				 edges{
					 id
					 name
					 avatar
					 roleID
					 gender
					 color
					 birthDate
					 ip
					 createdAt
					 updatedAt
				}
			}
		   }
		`, model.SearchParam{Page: ptrs.Int64Ptr(1)}}, false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postSearch(tt.args.query, &resp, "users", tt.args.searchParam)
			require.NotZero(t, resp.Users.TotalCount)
		})
	}
}

func Test_queryResolver_Videos(t *testing.T) {
	var resp struct {
		Videos model.VideoConnection
	}
	type args struct {
		query               string
		searchParam         model.SearchParam
		isFilterVideoSeries *bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test Videos not filter", args{`
		query videos($searchParam:SearchParam!) {
			videos(searchParam: $searchParam){
				totalCount
				edges{
					  id
					 title
					 desc
					 cover
					 pubDate
					 episodes{
					   id
					   num
					   title
					   desc
					   cover
					   url
					   subtitles{
						   name
						   url
					   }
					 createdAt
					 updatedAt
					 }
					 isShow
					 isHideOnMobile
					 theme
					 tags
					 createdAt
					 updatedAt
				}
			}
		   }
		`, model.SearchParam{Page: ptrs.Int64Ptr(1)}, ptrs.BoolPtr(false)}, false},
		{"test Videos filter", args{`
		query videos($searchParam:SearchParam!,$isFilterVideoSeries:Boolean=true) {
			videos(searchParam:$searchParam,isFilterVideoSeries:$isFilterVideoSeries){
				edges{
				   id 
				   title 
				}
			}
		   }
		`, model.SearchParam{Page: ptrs.Int64Ptr(1)}, ptrs.BoolPtr(true)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postSearch(tt.args.query, &resp, "videos", tt.args.searchParam)
			require.NotZero(t, resp.Videos.TotalCount)
		})
	}
}

func Test_queryResolver_VideoSerieses(t *testing.T) {
	var resp struct {
		VideoSerieses model.VideoSeriesConnection
	}
	type args struct {
		query       string
		searchParam model.SearchParam
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test VideoSerieses", args{`
		query videoSerieses($searchParam:SearchParam!) {
			videoSerieses(searchParam:$searchParam){
				 totalCount
				 edges{
					  id
					  name
					  items{
						videoSeriesID
						videoID
						title
						alias
						num
					  }
					  createdAt
					  updatedAt
				 }
			 }
			}
		`, model.SearchParam{Page: ptrs.Int64Ptr(1)}}, false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				postSearch(tt.args.query, &resp, "videoSerieses", tt.args.searchParam)
				require.NotZero(t, resp.VideoSerieses.TotalCount)
			})
		})
	}
}

func Test_queryResolver_SearchVideo(t *testing.T) {
	var resp struct {
		SearchVideo model.VideoIndexConnection
	}
	type args struct {
		query       string
		searchParam model.SearchParam
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test SearchVideo", args{`
		query searchVideo($searchParam:SearchParam!) {
			searchVideo(searchParam:$searchParam){
				edges{
					 id
					 title
					 desc
					 cover
					 totalNum
					 episodeID
				}
				totalCount
				aggResults{
				  key
				  value
				}
			}
		   }
		`, model.SearchParam{Page: ptrs.Int64Ptr(1)}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postSearch(tt.args.query, &resp, "searchVideo", tt.args.searchParam)
			require.NotZero(t, resp.SearchVideo.TotalCount)
			require.NotZero(t, len(resp.SearchVideo.AggResults))
			require.NotZero(t, len(resp.SearchVideo.Edges))
		})
	}
}

func Test_queryResolver_SimilarVideos(t *testing.T) {
	var resp struct {
		SimilarVideos model.VideoIndexConnection
	}
	type args struct {
		query       string
		searchParam model.SearchParam
		episodeID   int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test SimilarVideos", args{`
		query similarVideos($searchParam:SearchParam!,$episodeID:ID!) {
		similarVideos(searchParam: $searchParam, episodeID: $episodeID) {
			edges {
			  id
			  title
			  desc
			  cover
			  totalNum
			  episodeID
			}
		  }
		}
		`, model.SearchParam{Page: ptrs.Int64Ptr(1)}, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postSearch(tt.args.query, &resp, "similarVideos", tt.args.searchParam,
				client.Var("episodeID", tt.args.episodeID))
			require.Zero(t, resp.SimilarVideos.TotalCount)
			require.Zero(t, len(resp.SimilarVideos.AggResults))
			require.Zero(t, len(resp.SimilarVideos.Edges))
		})
	}
}

var deviceModelID int64

func Test_mutationResolver_CreateDeviceModel(t *testing.T) {
	var resp struct {
		CreateDeviceModel model.DeviceModel
	}
	type args struct {
		query string
		input model.NewDeviceModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test CreateDeviceModel", args{`
		mutation createDeviceModel($input:NewDeviceModel!){
			createDeviceModel(input:$input){
			  id
			}
		 }
		`, model.NewDeviceModel{
			Name:          "name",
			Desc:          ptrs.StringPtr("desc"),
			DeviceType:    1,
			CameraCompany: 0,
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postInput(tt.args.query, &resp, "createDeviceModel", tt.args.input)
			require.NotZero(t, resp.CreateDeviceModel.ID)
			deviceModelID = resp.CreateDeviceModel.ID
		})
	}
}

func Test_mutationResolver_UpdateDeviceModel(t *testing.T) {
	var resp struct {
		UpdateDeviceModel model.DeviceModel
	}
	type args struct {
		query string
		input model.NewUpdateDeviceModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test UpdateDeviceModel", args{`
		mutation updateDeviceModel($input:NewUpdateDeviceModel!){
			updateDeviceModel(input:$input){
			  id
			}
		 }
		`, model.NewUpdateDeviceModel{
			ID:   deviceModelID,
			Name: "name",
			Desc: ptrs.StringPtr("desc2"),
		}}, false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postInput(tt.args.query, &resp, "updateDeviceModel", tt.args.input)
			require.NotZero(t, resp.UpdateDeviceModel.ID)
		})
	}
}

var attributeModelID int64

func Test_mutationResolver_CreateAttributeModel(t *testing.T) {
	var resp struct {
		CreateAttributeModel model.AttributeModel
	}
	type args struct {
		query string
		input model.NewAttributeModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test CreateAttributeModel", args{`
		mutation createAttributeModel($input:NewAttributeModel!){
			createAttributeModel(input:$input){
			  id
			}
		 }
		`, model.NewAttributeModel{
			DeviceModelID: deviceModelID,
			Key:           "key",
			Name:          "name",
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postInput(tt.args.query, &resp, "createAttributeModel", tt.args.input)
			require.NotZero(t, resp.CreateAttributeModel.ID)
			attributeModelID = resp.CreateAttributeModel.ID
		})
	}
}

func Test_mutationResolver_UpdateAttributeModel(t *testing.T) {
	var resp struct {
		UpdateAttributeModel model.AttributeModel
	}
	type args struct {
		query string
		input model.NewUpdateAttributeModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test UpdateAttributeModel", args{`
		mutation updateAttributeModel($input:NewUpdateAttributeModel!){
			updateAttributeModel(input:$input){
			  id
			}
		 }
		`, model.NewUpdateAttributeModel{
			ID:   attributeModelID,
			Name: "name2",
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postInput(tt.args.query, &resp, "updateAttributeModel", tt.args.input)
			require.NotZero(t, resp.UpdateAttributeModel.ID)
		})
	}
}

func Test_mutationResolver_DeleteAttributeModel(t *testing.T) {
	var resp struct {
		DeleteAttributeModel model.AttributeModel
	}
	type args struct {
		query string
		id    int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test DeleteAttributeModel", args{`
		mutation deleteAttributeModel($id:Int!){
			deleteAttributeModel(id:$id){
			  id
			}
		 }
		`, attributeModelID}, false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			post(tt.args.query, &resp, "deleteAttributeModel", client.Var("id", tt.args.id))
		})
	}
}

var telemetryModelID int64

func Test_mutationResolver_CreateTelemetryModel(t *testing.T) {
	var resp struct {
		CreateTelemetryModel model.TelemetryModel
	}
	type args struct {
		query string
		input model.NewTelemetryModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test CreateTelemetryModel", args{`
		mutation createTelemetryModel($input:NewTelemetryModel!){
			createTelemetryModel(input:$input){
			  id
			}
		 }
		`, model.NewTelemetryModel{
			DeviceModelID: 1,
			Key:           "key",
			Name:          "name",
			Factor:        1,
			Unit:          ptrs.StringPtr("A"),
			UnitName:      ptrs.StringPtr("安培"),
			Scale:         1,
		}}, false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postInput(tt.args.query, &resp, "createTelemetryModel", tt.args.input)
			require.NotZero(t, resp.CreateTelemetryModel.ID)
			telemetryModelID = resp.CreateTelemetryModel.ID
		})
	}
}

func Test_mutationResolver_UpdateTelemetryModel(t *testing.T) {
	var resp struct {
		UpdateTelemetryModel model.TelemetryModel
	}
	type args struct {
		query string
		input model.NewUpdateTelemetryModel
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test UpdateTelemetryModel", args{`
		mutation updateTelemetryModel($input:NewUpdateTelemetryModel!){
			updateTelemetryModel(input:$input){
			  id
			}
		 }
		`, model.NewUpdateTelemetryModel{
			ID:       telemetryModelID,
			Name:     "name2",
			Factor:   2,
			Unit:     "V",
			UnitName: "伏特",
			Scale:    2,
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postInput(tt.args.query, &resp, "updateTelemetryModel", tt.args.input)
			require.NotZero(t, resp.UpdateTelemetryModel.ID)
		})
	}
}

func Test_mutationResolver_DeleteTelemetryModel(t *testing.T) {
	var resp struct {
		DeleteTelemetryModel model.TelemetryModel
	}
	type args struct {
		query string
		id    int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test DeleteTelemetryModel", args{`
		mutation deleteTelemetryModel($id:Int!){
			deleteTelemetryModel(id:$id){
			  id
			}
		 }
		`, telemetryModelID}, false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			post(tt.args.query, &resp, "deleteTelemetryModel", client.Var("id", tt.args.id))
		})
	}
}

func Test_mutationResolver_CreateDevice(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		input model.NewDevice
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Device
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.CreateDevice(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("mutationResolver.CreateDevice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mutationResolver.CreateDevice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mutationResolver_UpdateDevice(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		input model.NewUpdateDevice
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Device
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.UpdateDevice(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("mutationResolver.UpdateDevice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mutationResolver.UpdateDevice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mutationResolver_CreateDeviceDashboard(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		input model.NewDeviceDashboard
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.DeviceDashboard
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.CreateDeviceDashboard(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("mutationResolver.CreateDeviceDashboard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mutationResolver.CreateDeviceDashboard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mutationResolver_UpdateDeviceDashboard(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		input model.NewUpdateDeviceDashboard
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.DeviceDashboard
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.UpdateDeviceDashboard(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("mutationResolver.UpdateDeviceDashboard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mutationResolver.UpdateDeviceDashboard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mutationResolver_DeleteDeviceDashboard(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.DeviceDashboard
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.DeleteDeviceDashboard(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("mutationResolver.DeleteDeviceDashboard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mutationResolver.DeleteDeviceDashboard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mutationResolver_AddDeviceDashboardTelemetry(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		input model.NewDeviceDashboardTelemetry
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.DeviceDashboard
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.AddDeviceDashboardTelemetry(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("mutationResolver.AddDeviceDashboardTelemetry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mutationResolver.AddDeviceDashboardTelemetry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mutationResolver_RemoveDeviceDashboardTelemetry(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx context.Context
		ids []int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.DeviceDashboard
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.RemoveDeviceDashboardTelemetry(tt.args.ctx, tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("mutationResolver.RemoveDeviceDashboardTelemetry() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mutationResolver.RemoveDeviceDashboardTelemetry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mutationResolver_AddDeviceDashboardCamera(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		input model.NewDeviceDashboardCamera
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.DeviceDashboard
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.AddDeviceDashboardCamera(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("mutationResolver.AddDeviceDashboardCamera() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mutationResolver.AddDeviceDashboardCamera() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mutationResolver_RemoveDeviceDashboardCamera(t *testing.T) {
	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx context.Context
		ids []int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.DeviceDashboard
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.RemoveDeviceDashboardCamera(tt.args.ctx, tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("mutationResolver.RemoveDeviceDashboardCamera() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mutationResolver.RemoveDeviceDashboardCamera() = %v, want %v", got, tt.want)
			}
		})
	}
}
