package main

import "testing"

func TestReplaceURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test replace url",
			args{"http://pdc.9d77v.me:7500/video/%5BCASO%5D%5BSuzumiya_Haruhi_no_Yuuutsu%5D%5B01%5D%5BBDRIP%5D%5B1920x1080%5D%5Bx264_FLAC_2%5D%5BC39D66D3%5D.mp4"},
			"/video/%5BCASO%5D%5BSuzumiya_Haruhi_no_Yuuutsu%5D%5B01%5D%5BBDRIP%5D%5B1920x1080%5D%5Bx264_FLAC_2%5D%5BC39D66D3%5D.mp4"}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceURL(tt.args.url); got != tt.want {
				t.Errorf("ReplaceURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
