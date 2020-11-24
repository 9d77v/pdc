package camera

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
	"time"
)

func TestCapture(t *testing.T) {
	type args struct {
		cameraType int
		ip         string
		user       string
		password   string
	}
	tests := []struct {
		name string
		args args
	}{
		// {name: "测试海康摄像头抓拍", args: args{cameraType: CameraTypeHikvision, ip: "ip", user: "admin", password: "password"}},
		// {name: "测试大华摄像头抓拍", args: args{cameraType: CameraTypeDaHua, ip: "ip", user: "admin", password: "password"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Capture(tt.args.cameraType, tt.args.ip, tt.args.user, tt.args.password)
			if result != nil {
				saveImage(result, tt.args.ip)
			}
		})
	}
}

func saveImage(data []byte, ip string) {
	now := time.Now().Unix()
	imageURL := fmt.Sprintf("tmp/%s_%d.jpg", ip, now)
	file, err := os.Create(imageURL)
	if err != nil {
		log.Println("create error", err)
		return
	}
	defer file.Close()
	_, err = io.Copy(file, bytes.NewReader(data))
	if err != nil {
		log.Println("copy error", err)
	}
}
