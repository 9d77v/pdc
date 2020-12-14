package sdk

import (
	"log"
	"os"

	"github.com/gorilla/websocket"
)

var (
	//WSURL websocket服务地址
	WSURL     = GetEnvStr("WS_URL", "ws://localhost:8080/ws/iot/device")
	accessKey = GetEnvStr("PDC_DEVICE_ACCESSKEY", "")
	secretKey = GetEnvStr("PDC_DEVICE_SECRETKEY", "")
)

var (
	wsConn *websocket.Conn
)

func init() {
	c, _, err := websocket.DefaultDialer.Dial(WSURL, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	wsConn = c
}

//GetEnvStr ..
func GetEnvStr(key, value string) string {
	data := os.Getenv(key)
	if data == "" {
		return value
	}
	return data
}
