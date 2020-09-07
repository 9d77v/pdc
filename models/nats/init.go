package nats

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/stan.go"

	"github.com/9d77v/pdc/utils"
)

var (
	natsURL = utils.GetEnvStr("NATS_URL", "domain.local:4222")
)
var (
	//Client ..
	Client stan.Conn
)

//mq constants
const (
	SubjectVideo         = "video"
	SubjectGetDeviceInfo = "GetDeviceInfo"
	GroupVideo           = "video-group"
)

func init() {
	Client, _ = stan.Connect("test-cluster",
		fmt.Sprintf("client-%d", time.Now().Unix()),
		stan.Pings(10, 5),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}),
		stan.NatsURL("nats://"+natsURL),
	)
}
