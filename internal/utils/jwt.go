package utils

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//MyCustomClaims jwt自定义内容
type MyCustomClaims struct {
	UID string `json:"uid"` //用户id
	jwt.StandardClaims
}

//JWT 生成jwt字符串
func JWT(signingKey []byte, uid string, tokenExpire time.Duration, issuer string) string {
	claims := MyCustomClaims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpire).Unix(),
			Issuer:    issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		log.Println("generate token error:", err)
	}
	return tokenString
}

//ParseJWT 解析jwt字符串
func ParseJWT(signingKey []byte, token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
}