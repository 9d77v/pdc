package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	hashids "github.com/speps/go-hashids"
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

//GenerateHashID ..
func GenerateHashID(id uint, salt string, length int) string {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = length
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

//GetRawID ..
func GetRawID(id string, salt string, length int) uint {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = length
	h, err := hashids.NewWithData(hd)
	if err != nil {
		log.Println("NewWithData error:", err)
		return 0
	}
	nums, err := h.DecodeWithError(id)
	if err != nil {
		log.Println("hash encode error:", err)
		return 0
	}
	return uint(nums[0])
}

//GenerateSecretKey ..
func GenerateSecretKey(length int) string {
	return RandString(length)
}
