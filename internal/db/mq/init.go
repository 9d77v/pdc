package mq

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/9d77v/go-pkg/env"
	"github.com/nats-io/nats.go"
)

var (
	natsURL = env.GetEnvStr("NATS_URL", "nats://domain.local:4222")
)
var (
	client *nats.Conn
	once   sync.Once
)

//video subject
const (
	StreamVideo          = "VIDEO"
	SubjectVideo         = "video"
	SubjectGetDeviceInfo = "GetDeviceInfo"
	GroupVideo           = "video-group"
)

//device subject
const (
	StreamDevice                 = "DEVICE"
	StreamDeviceData             = "DEVICE_DATA"
	StreamDeviceTelemetry        = "DEVICE_TELEMETRY"
	StreamDeviceHealth           = "DEVICE_HEALTH"
	SubjectDevicPrefix           = "device."
	SubjectDeviceData            = "device_data"
	SubjectDeviceTelemetryPrefix = "device.telemetry."
	SubjectDeviceHealthPrefix    = "device.health."
	GroupDevice                  = "group.device"
	GroupSaveDeviceData          = "group.device.data.save"
	GroupPublishDeviceData       = "group.device.data.pub"
)

//GetClient get mq connection
func GetClient() *nats.Conn {
	once.Do(func() {
		client = initClient()
	})
	return client
}

//GetClient get mq connection
func GetJetStream() (nats.JetStreamContext, error) {
	once.Do(func() {
		client = initClient()
	})
	return client.JetStream(nats.PublishAsyncMaxPending(256))
}

func initClient() *nats.Conn {
	conn, err := nats.Connect(natsURL,
		nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(5),
		nats.ReconnectWait(2*time.Second),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			fmt.Printf("Got disconnected! Reason: %q\n", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			fmt.Printf("Got reconnected to %v!\n", nc.ConnectedUrl())
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			fmt.Printf("Connection closed. Reason: %q\n", nc.LastError())
		}))
	if err != nil {
		log.Println("nats connect error:", err)
	}
	return conn
}
