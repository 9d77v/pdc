package init

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/9d77v/go-pkg/env"
	"github.com/9d77v/pdc/internal/db/clickhouse"
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/db/mq"
	"github.com/9d77v/pdc/internal/db/oss"

	ch_device "github.com/9d77v/pdc/internal/module/device-service/chmodels"
	ch_history "github.com/9d77v/pdc/internal/module/history-service/chmodels"

	device_consumers "github.com/9d77v/pdc/internal/module/device-service/consumers"
	device "github.com/9d77v/pdc/internal/module/device-service/models"

	book "github.com/9d77v/pdc/internal/module/book-service/models"
	history "github.com/9d77v/pdc/internal/module/history-service/models"
	note "github.com/9d77v/pdc/internal/module/note-service/models"
	thing "github.com/9d77v/pdc/internal/module/thing-service/models"
	user "github.com/9d77v/pdc/internal/module/user-service/models"
	video_consumers "github.com/9d77v/pdc/internal/module/video-service/consumers"
	video "github.com/9d77v/pdc/internal/module/video-service/models"
)

var (
	ownerName     = env.GetEnvStr("ADMIN_NAME", "admin")
	ownerPassword = env.GetEnvStr("ADMIN_PASSWORD", "123456")
)

func init() {
	autoMergePostgresTables()
	autoMergeClickhouseTables()
	new(user.User).GenerateAdminAccount(ownerName, ownerPassword)
	oss.InitMinioBuckets()
	initSubscribe()
	initConsumers()
	initElasticSearchIndexes()
}

func autoMergePostgresTables() {
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
		//note
		&note.Note{},
		&note.NoteHistory{},
		//book
		&book.Book{},
		&book.Bookshelf{},
		&book.BookPosition{},
		&book.BookBorrowReturn{},
	)
	if err != nil {
		log.Println("auto migrate error:", err)
	}
}

func autoMergeClickhouseTables() {
	err := clickhouse.GetDB().Set("gorm:table_options",
		"engine=MergeTree() ORDER BY (device_id,telemetry_id,action_time) PARTITION BY (device_id)").
		AutoMigrate(&ch_device.DeviceTelemetry{})
	if err != nil {
		log.Println("auto migrate error:", err)
	}
	err = clickhouse.GetDB().Set("gorm:table_options",
		"engine=MergeTree() ORDER BY (device_id,action_time) PARTITION BY (device_id)").
		AutoMigrate(&ch_device.DeviceHealth{})
	if err != nil {
		log.Println("auto migrate error:", err)
	}
	err = clickhouse.GetDB().Set("gorm:table_options",
		"engine=MergeTree() ORDER BY (source_type,uid,source_id,sub_source_id,server_ts) PARTITION BY (source_type,uid)").
		AutoMigrate(&ch_history.HistoryLog{})
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

func initSubScriptions() []*nats.Subscription {
	js, err := mq.GetJetStream()
	if err != nil {
		log.Println("get jetstream failed:", err)
	}
	//VIDEO
	s, err := js.StreamInfo(mq.StreamVideo)
	if s == nil {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     mq.StreamVideo,
			Subjects: []string{mq.SubjectVideo},
		})
		if err != nil {
			log.Printf("AddStream %s failed:%v\n", mq.SubjectVideo, err)
		}
	}
	c, err := js.ConsumerInfo(mq.StreamVideo, "MONITOR")
	if c == nil {
		_, err = js.AddConsumer(mq.StreamVideo, &nats.ConsumerConfig{
			Durable:   "MONITOR",
			AckPolicy: nats.AckExplicitPolicy,
		})
		if err != nil {
			log.Printf("AddConsumer %s failed:%v\n", mq.SubjectVideo, err)
		}
	}
	//DEVICE
	s, err = js.StreamInfo(mq.StreamDevice)
	if s == nil {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     mq.StreamDevice,
			Subjects: []string{mq.SubjectDeviceData},
		})
		if err != nil {
			log.Printf("AddStream %s failed:%v\n", mq.SubjectDeviceData, err)
		}
	}
	c, err = js.ConsumerInfo(mq.StreamDevice, "MONITOR")
	if c == nil {
		_, err = js.AddConsumer(mq.StreamDevice, &nats.ConsumerConfig{
			Durable:   "MONITOR",
			AckPolicy: nats.AckExplicitPolicy,
		})
		if err != nil {
			log.Printf("AddConsumer %s failed:%v\n", mq.SubjectDeviceData, err)
		}
	}
	//QUEUE SUBSCRIBE
	qsub1, err := js.QueueSubscribe(mq.SubjectVideo,
		mq.GroupVideo, video_consumers.HandleVideoMSG, nats.DeliverLast())
	if err != nil {
		log.Panicln("QueueSubscribe error:", err)
	}
	qsub2, err := js.QueueSubscribe(mq.SubjectDeviceData, mq.GroupSaveDeviceData,
		device_consumers.HandleDeviceMsg, nats.DeliverLast())
	if err != nil {
		log.Panicln("SubscribeDeviceAttribute error:", err)
	}
	qsub3, err := js.QueueSubscribe(mq.SubjectDeviceData, mq.GroupPublishDeviceData,
		device_consumers.PublishDeviceData, nats.DeliverLast())
	if err != nil {
		log.Panicln("SubscribeDeviceAttribute error:", err)
	}
	return []*nats.Subscription{qsub1, qsub2, qsub3}
}

func unSubscribeMQQueues(qsubs []*nats.Subscription) {
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

func initElasticSearchIndexes() {
	js, err := mq.GetJetStream()
	if err != nil {
		log.Println("get jetstream failed:", err)
	}
	guid, err := js.PublishAsync(mq.SubjectVideo, []byte("0"))
	select {
	case <-js.PublishAsyncComplete():
	case <-time.After(5 * time.Second):
		fmt.Println("Did not resolve in time")
	}
	if err != nil {
		log.Println("mq publish failed,guid:", guid, " error:", err)
	}
}
