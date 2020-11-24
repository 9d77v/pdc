package mq

import (
	"fmt"
	"log"
	"time"

	"github.com/9d77v/pdc/internal/utils"
	"github.com/nats-io/stan.go"
)

var (
	natsURL = utils.GetEnvStr("NATS_URL", "domain.local:4222")
)
var (
	//Client ..
	Client stan.Conn
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

func init() {
	var err error
	Client, err = stan.Connect("test-cluster",
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
}
