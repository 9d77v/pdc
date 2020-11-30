package camera

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

//Dahua ..
type Dahua struct {
	camera
}

//NewDahua with ip,user,password
func NewDahua(ip, user, password string) *Dahua {
	d := &Dahua{
		camera: camera{
			ip:       ip,
			user:     user,
			password: password,
		},
	}
	d.requester = d
	return d
}

//Capture ..
func (d *Dahua) Capture() []byte {
	cameraURI := "/cgi-bin/snapshot.cgi"
	body, err := d.fetch(cameraURI)
	if err != nil {
		log.Println("fetch error:", err)
		return nil
	}
	return body
}

//GetDeviceInfo ..
func (d *Dahua) GetDeviceInfo() map[string]string {
	deviceInfoMap := make(map[string]string, 0)
	return deviceInfoMap
}

func (d *Dahua) getAuthorization(header http.Header, cameraURI string) string {
	dataStr := header["Www-Authenticate"][0]
	dataStr = strings.ReplaceAll(dataStr, "Digest ", "")
	arr := strings.Split(dataStr, ", ")

	nonce := strings.ReplaceAll(arr[2], "nonce=\"", "")
	nonce = strings.ReplaceAll(nonce, "\"", "")

	realm := strings.ReplaceAll(arr[0], "realm=", "")
	realm = strings.ReplaceAll(realm, "\"", "")
	opaque := strings.ReplaceAll(arr[3], "opaque=", "")
	opaque = strings.ReplaceAll(realm, "\"", "")

	nc, cnounce, response := d.getDigest(realm, nonce, cameraURI)
	return fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s",qop=auth,nc=%s,cnonce="%s", response="%s" opaque="%s"`,
		d.user, realm, nonce, cameraURI, nc, cnounce, response, opaque)
}
