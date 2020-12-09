package models

import (
	"testing"

	"github.com/lib/pq"
)

func TestVideoIndex_GetByID(t *testing.T) {
	type fields struct {
		ID             uint
		Title          string
		Desc           string
		PubDate        int64
		Cover          string
		TotalNum       int32
		Tags           pq.StringArray
		EpisodeID      uint
		IsShow         bool
		IsHideOnMobile bool
		SeriesID       uint
		SeriesName     string
		SeriesAlias    string
		SeriesNum      uint
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"test get by id", fields{}, args{1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VideoIndex{
				ID:             tt.fields.ID,
				Title:          tt.fields.Title,
				Desc:           tt.fields.Desc,
				PubDate:        tt.fields.PubDate,
				Cover:          tt.fields.Cover,
				TotalNum:       tt.fields.TotalNum,
				Tags:           tt.fields.Tags,
				EpisodeID:      tt.fields.EpisodeID,
				IsShow:         tt.fields.IsShow,
				IsHideOnMobile: tt.fields.IsHideOnMobile,
				SeriesID:       tt.fields.SeriesID,
				SeriesName:     tt.fields.SeriesName,
				SeriesAlias:    tt.fields.SeriesAlias,
				SeriesNum:      tt.fields.SeriesNum,
			}
			if err := v.GetByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("VideoIndex.GetByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestVideoIndex_Find(t *testing.T) {
	type fields struct {
		ID             uint
		Title          string
		Desc           string
		PubDate        int64
		Cover          string
		TotalNum       int32
		Tags           pq.StringArray
		EpisodeID      uint
		IsShow         bool
		IsHideOnMobile bool
		SeriesID       uint
		SeriesName     string
		SeriesAlias    string
		SeriesNum      uint
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"test get by id", fields{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VideoIndex{
				ID:             tt.fields.ID,
				Title:          tt.fields.Title,
				Desc:           tt.fields.Desc,
				PubDate:        tt.fields.PubDate,
				Cover:          tt.fields.Cover,
				TotalNum:       tt.fields.TotalNum,
				Tags:           tt.fields.Tags,
				EpisodeID:      tt.fields.EpisodeID,
				IsShow:         tt.fields.IsShow,
				IsHideOnMobile: tt.fields.IsHideOnMobile,
				SeriesID:       tt.fields.SeriesID,
				SeriesName:     tt.fields.SeriesName,
				SeriesAlias:    tt.fields.SeriesAlias,
				SeriesNum:      tt.fields.SeriesNum,
			}
			_, err := v.Find()
			if (err != nil) != tt.wantErr {
				t.Errorf("VideoIndex.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
