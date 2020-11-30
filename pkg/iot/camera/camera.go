package camera

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"
)

//CameraCompany 摄像头厂家
const (
	CameraCompanyHikvision = iota
	CameraCompanyDaHua
)

//Cameraer ..
type Cameraer interface {
	Capture() []byte
	GetDeviceInfo() map[string]string
}

type camera struct {
	ip       string
	user     string
	password string
}

func (c *camera) request(cameraURI string, authorization ...string) (*http.Response, error) {
	url := fmt.Sprintf("http://%s%s", c.ip, cameraURI)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if len(authorization) == 1 && authorization[0] != "" {
		req.Header.Set("Authorization", authorization[0])
	}
	client := &http.Client{Timeout: time.Second * 3}
	resp, err := client.Do(req)
	return resp, err
}

func (c *camera) getHeaderFromRequest(cameraURI string, authorization ...string) (http.Header, error) {
	resp, err := c.request(cameraURI, authorization...)
	if err != nil {
		return nil, err
	}
	return resp.Header, nil
}

func (c *camera) getBodyFromRequest(cameraURI string, authorization ...string) ([]byte, error) {
	resp, err := c.request(cameraURI, authorization...)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (c *camera) getDigest(realm, nonce, url string) (string, string, string) {
	nc := "0000011"
	cnounce := c.getRandomString(20)
	A1 := fmt.Sprintf("%s:%s:%s", c.user, realm, c.password)
	A2 := fmt.Sprintf("GET:%s", url)
	return nc, cnounce, c.getMD5(fmt.Sprintf("%s:%s:%s:%s:%s:%s", c.getMD5(A1), nonce, nc, cnounce, "auth", c.getMD5(A2)))
}

func (c *camera) getRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r, err := rand.Int(rand.Reader, big.NewInt(int64(len(bytes))))
	if err != nil {
		log.Printf("error occurred while generating random: %s", err.Error())
		return ""
	}
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Int64()])
	}
	return string(result)
}

func (c *camera) getMD5(s string) string {
	h := md5.New()
	_, err := h.Write([]byte(s))
	if err != nil {
		log.Println("md5 write err:", err)
	}
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

func (c *camera) saveImage(data []byte) {
	now := time.Now().Unix()
	imageURL := fmt.Sprintf("tmp/%s_%d.jpg", c.ip, now)
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
