package sdk

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/stan.go"
)

var (
	natsURL = GetEnvStr("NATS_URL", "domain.local:4222")
)

var (
	natsConn stan.Conn
)

//mq constantsb
const (
	subjectDevice          = "device.config"
	subjectDeviceData      = "device.data"
	groupDevice            = "group.device"
	groupSaveDeviceData    = "group.device.data.save"
	groupPublishDeviceData = "group.device.data.pub"
)

func init() {
	natsConn, _ = stan.Connect("test-cluster",
		fmt.Sprintf("iot-client-%d", time.Now().Unix()),
		stan.Pings(10, 5),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}),
		stan.NatsURL("nats://"+natsURL),
	)
}

//GetEnvStr ..
func GetEnvStr(key, value string) string {
	data := os.Getenv(key)
	if data == "" {
		return value
	}
	return data
}

//AckHandler for nats ack
func AckHandler(ackedNuid string, err error) {
	if err != nil {
		log.Printf("Warning: error publishing msg id %s: %v\n", ackedNuid, err.Error())
	} else {
		log.Printf("Received ack for msg id %s\n", ackedNuid)
	}
}