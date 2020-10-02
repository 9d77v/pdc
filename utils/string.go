package utils

import (
	"math/rand"
	"time"

	hashids "github.com/speps/go-hashids"
)

const (
	accessKeyLen = 12
	secretKeyLen = 32
	salt         = "pdc"
)

// RandString 生成随机字符串
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

//GenerateAccessKey ..
func GenerateAccessKey(id uint) string {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = accessKeyLen
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{int(id)})
	return e
}

//GenerateSecretKey ..
func GenerateSecretKey() string {
	return RandString(secretKeyLen)
}
