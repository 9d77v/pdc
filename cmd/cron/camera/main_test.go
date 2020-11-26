package main

import (
	"reflect"
	"testing"
)

func Test_getCameraIDs(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{"测试获取摄像头id数组", []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCameraIDs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCameraIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getYesterdayStr(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"测试获取昨天日期", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getYesterdayStr(); got != tt.want {
				t.Errorf("getYesterdayStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCameraPictureTmpDir(t *testing.T) {
	type args struct {
		id  uint
		day string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"测试获取摄像头图片临时目录", args{4, "2020-11-23"}, "picture/4/2020-11-23/tmp/"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCameraPictureTmpDir(tt.args.id, tt.args.day); got != tt.want {
				t.Errorf("getCameraPictureTmpDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getVideoPath(t *testing.T) {
	type args struct {
		id  uint
		day string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"测试获取摄像头图片合成的视频文件地址", args{4, "2020-11-23"}, "picture/4/2020-11-23/2020-11-23.mp4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getVideoPath(tt.args.id, tt.args.day); got != tt.want {
				t.Errorf("getVideoPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateVideo(t *testing.T) {
	type args struct {
		pictureTmpDir string
		videoFilePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"测试合成视频",
			args{"picture/4/2020-11-23/tmp/", "picture/4/2020-11-23/2020-11-23.mp4"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := generateVideo(tt.args.pictureTmpDir, tt.args.videoFilePath); (err != nil) != tt.wantErr {
				t.Errorf("generateVideo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_renamePictureFileNames(t *testing.T) {
	type args struct {
		pictureTmpDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"测试图片文件重命名", args{"picture/4/2020-11-23/tmp/"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := renamePictureFileNames(tt.args.pictureTmpDir); (err != nil) != tt.wantErr {
				t.Errorf("renamePictureFileNames() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getNewFileName(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"测试<10", args{0}, "0000"},
		{"测试<10", args{4}, "0004"},
		{"测试<100", args{33}, "0033"},
		{"测试<1000", args{344}, "0344"},
		{"测试<10000", args{1200}, "1200"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNewFileName(tt.args.i); got != tt.want {
				t.Errorf("getNewFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_saveVideoPath(t *testing.T) {
	type args struct {
		deviceID      uint
		date          string
		videoFilePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"测试保存视频", args{4, "2020-11-26", "/camera/picture/4/2020-11-25/2020-11-25.mp4"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := saveVideoPath(tt.args.deviceID, tt.args.date, tt.args.videoFilePath); (err != nil) != tt.wantErr {
				t.Errorf("saveVideoPath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
