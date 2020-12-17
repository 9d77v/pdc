package models

import (
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/module/base"
)

//CameraTimeLapseVideo 摄像头延时视频，定时每1分钟抓拍图片，第二天1点合成视频
type CameraTimeLapseVideo struct {
	base.DefaultModel
	DeviceID uint
	Date     string `gorm:"size:20"`
	VideoURL string `gorm:"size:100"`
}

//NewCameraTimeLapseVideo ..
func NewCameraTimeLapseVideo() *CameraTimeLapseVideo {
	vs := &CameraTimeLapseVideo{}
	vs.SetDB(db.GetDB())
	return vs
}
