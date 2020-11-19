package lib

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

//MD5 获取md5字符串
func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

//GetRandomString 获取随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//GetDigest 获取response,无auth，无cnonce
func GetDigest(user, realm, password, nonce, url string) (string, string, string) {
	nc := "0000011"
	cnounce := GetRandomString(20)
	A1 := fmt.Sprintf("%s:%s:%s", user, realm, password)
	A2 := fmt.Sprintf("GET:%s", url)
	return nc, cnounce, MD5(fmt.Sprintf("%s:%s:%s:%s:%s:%s", MD5(A1), nonce, nc, cnounce, "auth", MD5(A2)))
}
