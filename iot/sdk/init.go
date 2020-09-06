package sdk

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nats-io/stan.go"
)

var (
	natsURL = getEnvStr("NATS_URL", "domain.local:4222")
)

var (
	//Client ..
	client stan.Conn
)

//mq constantsb
const (
	subjectDevice                = "device.config"
	subjectDeviceData            = "device.data"
	subjectDeviceTelemetryPrefix = "device.telemetry"

	groupDevice     = "group.device"
	groupDeviceData = "group.device.data"
)

func init() {
	client, _ = stan.Connect("test-cluster",
		fmt.Sprintf("iot-client-%d", time.Now().Unix()),
		stan.Pings(10, 5),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}),
		stan.NatsURL("nats://"+natsURL),
	)
}

func getEnvStr(key, value string) string {
	data := os.Getenv(key)
	if data == "" {
		return value
	}
	return data
}

//ackHandler for nats ack
func ackHandler(ackedNuid string, err error) {
	if err != nil {
		log.Printf("Warning: error publishing msg id %s: %v\n", ackedNuid, err.Error())
	} else {
		log.Printf("Received ack for msg id %s\n", ackedNuid)
	}
}
