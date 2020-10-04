package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	hashids "github.com/speps/go-hashids"
)

const (
	accessKeyLen = 12
	secretKeyLen = 32
	salt         = "pdc"
)

// RandString 生成随机字符串
func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b, err := rand.Int(rand.Reader, big.NewInt(26))
		if err != nil {
			fmt.Println("rand.Int：", b, b.BitLen())
		}
		bytes[i] = byte(b.Int64() + 65)
	}
	return string(bytes)
}

//GenerateAccessKey ..
func GenerateAccessKey(id uint) string {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = accessKeyLen
	h, err := hashids.NewWithData(hd)
	if err != nil {
		log.Println("NewWithData error:", err)
		return ""
	}
	e, err := h.Encode([]int{int(id)})
	if err != nil {
		log.Println("hash encode error:", err)
		return ""
	}
	return e
}

//GenerateSecretKey ..
func GenerateSecretKey() string {
	return RandString(secretKeyLen)
}
