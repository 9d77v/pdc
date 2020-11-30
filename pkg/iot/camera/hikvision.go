package camera

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strings"
)

//Hikvision ..
type Hikvision struct {
	camera
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

//NewHikvision with ip,user,password
func NewHikvision(ip, user, password string) *Hikvision {
	hik := &Hikvision{
		camera: camera{
			ip:       ip,
			user:     user,
			password: password,
		},
	}
	hik.requester = hik
	return hik
}

//Capture ..
func (h *Hikvision) Capture() []byte {
	cameraURI := "/ISAPI/Streaming/channels/101/picture"
	body, err := h.fetch(cameraURI)
	if err != nil {
		log.Println("fetch error:", err)
		return nil
	}
	return body
}

//GetDeviceInfo ..
func (h *Hikvision) GetDeviceInfo() map[string]string {
	cameraURI := "/ISAPI/System/deviceInfo"
	body, err := h.fetch(cameraURI)
	if err != nil {
		log.Println("fetch error:", err)
		return nil
	}
	return h.parseDeviceInfo(body)
}

func (h *Hikvision) getAuthorization(header http.Header, cameraURI string) string {
	dataStr := header["Www-Authenticate"][0]

	dataStr = strings.ReplaceAll(dataStr, "Digest ", "")
	arr := strings.Split(dataStr, ", ")

	nonce := strings.ReplaceAll(arr[2], "nonce=\"", "")
	nonce = strings.ReplaceAll(nonce, "\"", "")

	realm := strings.ReplaceAll(arr[1], "realm=", "")
	realm = strings.ReplaceAll(realm, "\"", "")

	nc, cnounce, response := h.getDigest(realm, nonce, cameraURI)
	return fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s",qop=auth,nc=%s,cnonce="%s", response="%s"`,
		h.user, realm, nonce, cameraURI, nc, cnounce, response)
}

func (h *Hikvision) parseDeviceInfo(body []byte) map[string]string {
	deviceInfoMap := make(map[string]string, 0)

	deviceInfo := new(Hikvision)
	err := xml.Unmarshal(body, deviceInfo)
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
	return deviceInfoMap
}
