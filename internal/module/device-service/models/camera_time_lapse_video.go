package models

import (
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/db/oss"
	"github.com/9d77v/pdc/internal/module/base"
	"github.com/9d77v/pdc/internal/module/device-service/pb"
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
	m := &CameraTimeLapseVideo{}
	m.SetDB(db.GetDB())
	return m
}

//ToCameraTimeLapseVideoPBs ..
func (m *CameraTimeLapseVideo) ToCameraTimeLapseVideoPBs(data []*CameraTimeLapseVideo,
	scheme string) []*pb.CameraTimeLapseVideo {
	result := make([]*pb.CameraTimeLapseVideo, 0, len(data))
	for _, v := range data {
		r := m.toCameraTimeLapseVideoPB(v, scheme)
		result = append(result, r)
	}
	return result
}

func (m *CameraTimeLapseVideo) toCameraTimeLapseVideoPB(video *CameraTimeLapseVideo,
	scheme string) *pb.CameraTimeLapseVideo {
	return &pb.CameraTimeLapseVideo{
		Id:       int64(video.ID),
		DeviceId: int64(video.DeviceID),
		Date:     video.Date,
		VideoUrl: oss.GetOSSPrefixByScheme(scheme) + video.VideoURL,
	}
}
