package camera

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
)

//GetRandomString 获取随机字符串
func GetRandomString(length int) string {
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

//GetDigest 获取response,无auth，无cnonce
func GetDigest(user, realm, password, nonce, url string) (string, string, string) {
	nc := "0000011"
	cnounce := GetRandomString(20)
	A1 := fmt.Sprintf("%s:%s:%s", user, realm, password)
	A2 := fmt.Sprintf("GET:%s", url)
	return nc, cnounce, getMD5(fmt.Sprintf("%s:%s:%s:%s:%s:%s", getMD5(A1), nonce, nc, cnounce, "auth", getMD5(A2)))
}

func getMD5(s string) string {
	h := md5.New()
	_, err := h.Write([]byte(s))
	if err != nil {
		log.Println("md5 write err:", err)
	}
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}
