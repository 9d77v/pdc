package camera

import (
	"reflect"
	"testing"
)

func TestGetDeviceInfo(t *testing.T) {
	type args struct {
		cameraType int
		ip         string
		user       string
		password   string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// {name: "测试海康摄像头抓拍", args: args{cameraType: CameraTypeHikvision, ip: "ip", user: "admin", password: "password"}},
		// {name: "测试大华摄像头抓拍", args: args{cameraType: CameraTypeDaHua, ip: "ip", user: "admin", password: "password"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDeviceInfo(tt.args.cameraType, tt.args.ip, tt.args.user, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeviceInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
