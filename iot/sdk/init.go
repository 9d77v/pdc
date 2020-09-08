package sdk

import (
	"fmt"
	"log"
	"os"
	"time"

	redis "github.com/go-redis/redis/v8"
	"github.com/nats-io/stan.go"
)

var (
	natsURL       = getEnvStr("NATS_URL", "domain.local:4222")
	redisAddress  = getEnvStr("REDIS_ADDRESS", "domain.local:6379")
	redisPassword = getEnvStr("REDIS_PASSWORD", "")
)

var (
	natsClient  stan.Conn
	redisClient *redis.Client
)

//mq constantsb
const (
	subjectDevice                      = "device.config"
	subjectDeviceData                  = "device.data"
	subjectDeviceTelemetryPrefix       = "device.telemetry"
	subjectDeviceChangeTelemetryPrefix = "device.change_telemetry"
	groupDevice                        = "group.device"
	groupDeviceData                    = "group.device.data"
)

func init() {
	initNatsStreaming()
	initRedis()
}

func initNatsStreaming() {
	natsClient, _ = stan.Connect("test-cluster",
		fmt.Sprintf("iot-client-%d", time.Now().Unix()),
		stan.Pings(10, 5),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}),
		stan.NatsURL("nats://"+natsURL),
	)
}

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       0, // use default DB
	})
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
