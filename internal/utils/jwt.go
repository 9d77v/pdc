package utils

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//MyCustomClaims jwt自定义内容
type MyCustomClaims struct {
	UID      int64  `json:"uid"` //用户id
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
	jwt.StandardClaims
}

//JWT 生成jwt字符串
func JWT(signingKey []byte, uid int64, username, avatar string, gender int, tokenExpire time.Duration, issuer string) string {
	claims := MyCustomClaims{
		uid,
		username,
		avatar,
		gender,
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
