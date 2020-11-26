package mq

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/9d77v/pdc/internal/utils"
	"github.com/nats-io/stan.go"
)

var (
	natsURL = utils.GetEnvStr("NATS_URL", "domain.local:4222")
)
var (
	client stan.Conn
	once   sync.Once
)

//video subject
const (
	SubjectVideo         = "video"
	SubjectGetDeviceInfo = "GetDeviceInfo"
	GroupVideo           = "video-group"
)

//device subject
const (
	SubjectDevicPrefix           = "device."
	SubjectDeviceData            = "device.data"
	SubjectDeviceTelemetryPrefix = "device.telemetry."
	SubjectDeviceHealthPrefix    = "device.health."
	GroupDevice                  = "group.device"
	GroupSaveDeviceData          = "group.device.data.save"
	GroupPublishDeviceData       = "group.device.data.pub"
)

//GetClient get mq connection
func GetClient() stan.Conn {
	once.Do(func() {
		client = initClient()
	})
	return client
}

func initClient() stan.Conn {
	conn, err := stan.Connect("test-cluster",
		fmt.Sprintf("client-%d", time.Now().Unix()),
		stan.Pings(10, 5),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}),
		stan.NatsURL("nats://"+natsURL),
	)
	if err != nil {
		log.Println("nats connect error:", err)
	}
	return conn
}
