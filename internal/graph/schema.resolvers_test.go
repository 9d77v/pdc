package graph

import (
	"fmt"
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

func TestMain(m *testing.M) {
	login()
	m.Run()
}

func login() {
	cc.MustPost(`mutation login($username: String!, $password: String!) {
		login(username: $username, password: $password) {
		  accessToken
		  refreshToken
		}
	  }`,
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
		input model.NewHistoryInput
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test RecordHistory", args{
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
			cc.MustPost(`
			mutation recordHistory($input: NewHistoryInput!) {
				recordHistory(input: $input) {
				  subSourceID
				}
			  }
			`, &resp,
				client.Path("/api"),
				client.Operation("recordHistory"),
				client.AddHeader("Authorization", loginResponse.Login.AccessToken),
				client.Var("input", tt.args.input))
			require.NotZero(t, resp.RecordHistory.SubSourceID)
		})
	}
}

func Test_queryResolver_Histories(t *testing.T) {
	var resp struct {
		Histories model.HistoryConnection
	}
	var sourceType int64 = 1
	var page int64 = 1
	type args struct {
		sourceType  *int64
		searchParam model.SearchParam
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test histories", args{&sourceType, model.SearchParam{Page: &page}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc.MustPost(`
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
			`, &resp,
				client.Path("/api"),
				client.Operation("histories"),
				client.AddHeader("Authorization", loginResponse.Login.AccessToken),
				client.Var("searchParam", tt.args.searchParam),
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
		input model.NewVideo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test CreateVideo", args{model.NewVideo{
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
			cc.MustPost(`
			mutation createVideo($input:NewVideo!){
				createVideo(input:$input){
				  id
				}
			 }
			`, &resp,
				client.Path("/api"),
				client.Operation("createVideo"),
				client.AddHeader("Authorization", loginResponse.Login.AccessToken),
				client.Var("input", tt.args.input))
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
		input model.NewUpdateVideo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test UpdateVideo", args{model.NewUpdateVideo{
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
			cc.MustPost(`
			mutation updateVideo($input:NewUpdateVideo!){
				updateVideo(input:$input){
				  id
				}
			 }
			`, &resp,
				client.Path("/api"),
				client.Operation("updateVideo"),
				client.AddHeader("Authorization", loginResponse.Login.AccessToken),
				client.Var("input", tt.args.input))
			require.NotZero(t, resp.UpdateVideo.ID)
		})
	}
}
