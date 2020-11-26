package models

import "gorm.io/gorm"

//CameraTimeLapseVideo 摄像头延时视频，定时每1分钟抓拍图片，第二天1点合成视频
type CameraTimeLapseVideo struct {
	gorm.Model
	DeviceID uint
	Date     string `gorm:"size:20"`
	VideoURL string `gorm:"size:100"`
}

//NewCameraTimeLapseVideo ..
func NewCameraTimeLapseVideo(deviceID uint, date, videoURL string) *CameraTimeLapseVideo {
	return &CameraTimeLapseVideo{
		DeviceID: deviceID,
		Date:     date,
		VideoURL: videoURL,
	}
}
