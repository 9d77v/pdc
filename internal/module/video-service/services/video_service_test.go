package services

import (
	"context"
	"testing"

	"github.com/9d77v/go-lib/ptrs"
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/graph/model"
	"github.com/9d77v/pdc/internal/module/video-service/models"
	"github.com/stretchr/testify/assert"
)

func TestVideoService_CreateVideo(t *testing.T) {
	type args struct {
		ctx    context.Context
		input  model.NewVideo
		scheme string
	}
	tests := []struct {
		name    string
		s       VideoService
		args    args
		wantErr bool
	}{
		{"test ceraete video", VideoService{}, args{context.Background(), model.NewVideo{
			Title:          "title",
			Desc:           ptrs.StringPtr("desc"),
			PubDate:        ptrs.Int64Ptr(123),
			Cover:          ptrs.StringPtr("dd"),
			Tags:           []string{"d", "d"},
			IsShow:         true,
			IsHideOnMobile: true,
			Theme:          "theme",
		}, "http"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := VideoService{}
			got, err := s.CreateVideo(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("VideoService.CreateVideo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			removeVideo(t, uint(got.ID))
			assert.NotZero(t, got.ID, "id should not be zero")
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
