package lib

//CameraType 摄像头类型
const (
	CameraCompanyHikvision = iota
	CameraCompanyDaHua
)

//CameraDeviceInfoURLMap 摄像头设备信息map
var CameraDeviceInfoURLMap = map[int]string{
	CameraCompanyHikvision: "/ISAPI/System/deviceInfo",
	CameraCompanyDaHua:     "",
}

//CameraCaptureURLMap 摄像头抓拍map
var CameraCaptureURLMap = map[int]string{
	CameraCompanyHikvision: "/ISAPI/Streaming/channels/101/picture",
	CameraCompanyDaHua:     "/cgi-bin/snapshot.cgi",
}
