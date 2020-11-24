package camera

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

//HikvisionDeviceInfo ..
type HikvisionDeviceInfo struct {
	DeviceName           string `xml:"deviceName" json:"device_name"`
	DeviceID             string `xml:"deviceID" json:"device_id"`
	DeviceDescription    string `xml:"deviceDescription" json:"device_description"`
	DeviceLocation       string `xml:"deviceLocation" json:"device_location"`
	SystemContact        string `xml:"systemContact" json:"system_contact"`
	Model                string `xml:"model" json:"model"`
	SerialNumber         string `xml:"serialNumber" json:"serial_number"`
	MacAddress           string `xml:"macAddress" json:"mac_address"`
	FirmwareVersion      string `xml:"firmwareVersion" json:"firmware_version"`
	FirmwareReleasedDate string `xml:"firmwareReleasedDate" json:"firmware_releasedDate"`
	EncoderVersion       string `xml:"encoderVersion" json:"encoder_version"`
	EncoderReleasedDate  string `xml:"encoderReleasedDate" json:"encoder_released_date"`
	BootVersion          string `xml:"bootVersion" json:"boot_version"`
	BootReleasedDate     string `xml:"bootReleasedDate" json:"boot_released_date"`
	HardwareVersion      string `xml:"hardwareVersion" json:"hardware_version"`
	DeviceType           string `xml:"deviceType" json:"device_type"`
	TelecontrolID        string `xml:"telecontrolID" json:"telecontrol_id"`
	SupportBeep          string `xml:"supportBeep" json:"support_beep"`
	SupportVideoLoss     string `xml:"supportVideoLoss" json:"support_video_loss"`
	FirmwareVersionInfo  string `xml:"firmwareVersionInfo" json:"firmware_version_info"`
}

//GetDeviceInfo 获取设备信息
func GetDeviceInfo(cameraType int, ip, user, password string) map[string]string {
	deviceInfoMap := make(map[string]string, 0)
	cameraURI := CameraDeviceInfoURLMap[cameraType]
	var resp *http.Response
	var err error
	url := fmt.Sprintf("http://%s%s", ip, cameraURI)
	log.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error reading request. ", err)
	}
	client := &http.Client{Timeout: time.Second * 3}
	resp, err = client.Do(req)
	if err != nil {
		log.Println("Error reading response. ", err)
		return deviceInfoMap
	}
	if resp.StatusCode == http.StatusUnauthorized {
		dataStr := resp.Header["Www-Authenticate"][0]
		fmt.Println(dataStr)
		dataStr = strings.ReplaceAll(dataStr, "Digest ", "")
		arr := strings.Split(dataStr, ", ")

		nonce := strings.ReplaceAll(arr[2], "nonce=\"", "")
		nonce = strings.ReplaceAll(nonce, "\"", "")

		authorization := ""
		if cameraType == CameraCompanyHikvision {
			realm := strings.ReplaceAll(arr[1], "realm=", "")
			realm = strings.ReplaceAll(realm, "\"", "")

			nc, cnounce, response := GetDigest(user, realm, password, nonce, cameraURI)
			authorization = fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s",qop=auth,nc=%s,cnonce="%s", response="%s"`, user, realm, nonce, cameraURI, nc, cnounce, response)
		} else if cameraType == CameraCompanyDaHua {
			return deviceInfoMap
			// realm := strings.ReplaceAll(arr[0], "realm=", "")
			// realm = strings.ReplaceAll(realm, "\"", "")
			// opaque := strings.ReplaceAll(arr[3], "opaque=", "")
			// opaque = strings.ReplaceAll(realm, "\"", "")

			// nc, cnounce, response := GetDigest(user, realm, password, nonce, cameraURI)
			// authorization = fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s",qop=auth,nc=%s,cnonce="%s", response="%s" opaque="%s"`,
			// 	user, realm, nonce, cameraURI, nc, cnounce, response, opaque)
		}
		req.Header.Set("Authorization", authorization)
		resp, err = client.Do(req)
		if err != nil {
			log.Println("Error reading response with Authorization:", err)
			return deviceInfoMap
		}
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error ReadAll:", err)
		return deviceInfoMap
	}
	if cameraType == CameraCompanyHikvision {
		deviceInfo := new(HikvisionDeviceInfo)
		err := xml.Unmarshal(result, deviceInfo)
		if err != nil {
			log.Println("Error xml Unmarshal.:", err)
			return deviceInfoMap
		}
		jsonData, err := json.Marshal(deviceInfo)
		if err != nil {
			log.Println("json marshal error", err)
			return deviceInfoMap
		}
		err = json.Unmarshal(jsonData, &deviceInfoMap)
		if err != nil {
			log.Println("json Unmarshal error", err)
			return deviceInfoMap
		}
	}
	return deviceInfoMap
}
