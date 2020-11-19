package lib

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

//Capture 抓图
func Capture(cameraCompany int, ip, user, password string) []byte {
	cameraURI := CameraCaptureURLMap[cameraCompany]
	var resp *http.Response
	var err error
	url := fmt.Sprintf("http://%s%s", ip, cameraURI)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error reading request. ", err)
	}
	client := &http.Client{Timeout: time.Second * 3}
	resp, err = client.Do(req)
	if err != nil {
		log.Println("Error reading response. ", err)
		return nil
	}
	if resp.StatusCode == http.StatusUnauthorized {
		dataStr := resp.Header["Www-Authenticate"][0]
		fmt.Println(dataStr)
		dataStr = strings.ReplaceAll(dataStr, "Digest ", "")
		arr := strings.Split(dataStr, ", ")

		nonce := strings.ReplaceAll(arr[2], "nonce=\"", "")
		nonce = strings.ReplaceAll(nonce, "\"", "")

		authorization := ""
		if cameraCompany == CameraCompanyHikvision {
			realm := strings.ReplaceAll(arr[1], "realm=", "")
			realm = strings.ReplaceAll(realm, "\"", "")

			nc, cnounce, response := GetDigest(user, realm, password, nonce, cameraURI)
			authorization = fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s",qop=auth,nc=%s,cnonce="%s", response="%s"`, user, realm, nonce, cameraURI, nc, cnounce, response)
		} else if cameraCompany == CameraCompanyDaHua {
			realm := strings.ReplaceAll(arr[0], "realm=", "")
			realm = strings.ReplaceAll(realm, "\"", "")
			opaque := strings.ReplaceAll(arr[3], "opaque=", "")
			opaque = strings.ReplaceAll(realm, "\"", "")

			nc, cnounce, response := GetDigest(user, realm, password, nonce, cameraURI)
			authorization = fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s",qop=auth,nc=%s,cnonce="%s", response="%s" opaque="%s"`,
				user, realm, nonce, cameraURI, nc, cnounce, response, opaque)
		}
		req.Header.Set("Authorization", authorization)
		resp, err = client.Do(req)
		if err != nil {
			log.Println("Error reading response. ", err)
			return nil
		}
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response. ", err)
		return nil
	}
	return result
}
