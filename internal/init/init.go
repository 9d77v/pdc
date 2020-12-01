package init

import (
	"log"
	"os"
	"os/signal"

	"github.com/nats-io/stan.go"

	"github.com/9d77v/pdc/internal/db/clickhouse"
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/db/mq"
	"github.com/9d77v/pdc/internal/db/oss"
	device_consumers "github.com/9d77v/pdc/internal/module/device-service/consumers"
	device "github.com/9d77v/pdc/internal/module/device-service/models"
	history "github.com/9d77v/pdc/internal/module/history-service/models"
	thing "github.com/9d77v/pdc/internal/module/thing-service/models"
	user "github.com/9d77v/pdc/internal/module/user-service/models"
	video_consumers "github.com/9d77v/pdc/internal/module/video-service/consumers"
	video "github.com/9d77v/pdc/internal/module/video-service/models"
	"github.com/9d77v/pdc/internal/utils"
)

var (
	ownerName     = utils.GetEnvStr("ADMIN_NAME", "admin")
	ownerPassword = utils.GetEnvStr("ADMIN_PASSWORD", "123456")
)

func init() {
	autoMergeTables()
	new(user.User).GenerateAdminAccount(ownerName, ownerPassword)
	clickhouse.CreateTables()
	oss.InitMinioBuckets()
	initSubscribe()
	initConsumers()
}

func autoMergeTables() {
	err := db.GetDB().AutoMigrate(
		//device
		&device.DeviceModel{},
		&device.TelemetryModel{},
		&device.AttributeModel{},
		&device.Device{},
		&device.Attribute{},
		&device.Telemetry{},
		&device.DeviceDashboard{},
		&device.DeviceDashboardTelemetry{},
		&device.DeviceDashboardCamera{},
		&device.CameraTimeLapseVideo{},
		//history
		&history.History{},
		//thing
		&thing.Thing{},
		//user
		&user.User{},
		//video
		&video.Video{},
		&video.Episode{},
		&video.Subtitle{},
		&video.VideoSeries{},
		&video.VideoSeriesItem{},
	)
	if err != nil {
		log.Println("auto migrate error:", err)
	}
}

func initSubscribe() {
	go func() {
		qsubs := initSubScriptions()
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)
		<-interrupt
		unSubscribeMQQueues(qsubs)
	}()
}

func initSubScriptions() []stan.Subscription {
	qsub1, err := mq.GetClient().QueueSubscribe(mq.SubjectVideo,
		mq.GroupVideo, video_consumers.HandleVideoMSG)
	if err != nil {
		log.Panicln("QueueSubscribe error:", err)
	}
	qsub2, err := mq.GetClient().QueueSubscribe(mq.SubjectDeviceData, mq.GroupSaveDeviceData,
		device_consumers.HandleDeviceMsg, stan.DurableName("dur"))
	if err != nil {
		log.Panicln("SubscribeDeviceAttribute error:", err)
	}
	qsub3, err := mq.GetClient().QueueSubscribe(mq.SubjectDeviceData, mq.GroupPublishDeviceData,
		device_consumers.PublishDeviceData, stan.DurableName("dur"))
	if err != nil {
		log.Panicln("SubscribeDeviceAttribute error:", err)
	}
	return []stan.Subscription{qsub1, qsub2, qsub3}
}

func unSubscribeMQQueues(qsubs []stan.Subscription) {
	for _, qsub := range qsubs {
		err := qsub.Unsubscribe()
		if err != nil {
			log.Println("qsub Unsubscribe error:", err)
		}
	}
}

func initConsumers() {
	go device_consumers.SaveDeviceTelemetry()
	go device_consumers.SaveDeviceHealth()
}
