package init

import (
	"log"

	"github.com/9d77v/pdc/internal/db/clickhouse"
	"github.com/9d77v/pdc/internal/db/db"
	"github.com/9d77v/pdc/internal/db/oss"
	device "github.com/9d77v/pdc/internal/module/device-service/models"
	history "github.com/9d77v/pdc/internal/module/history-service/models"
	thing "github.com/9d77v/pdc/internal/module/thing-service/models"
	user "github.com/9d77v/pdc/internal/module/user-service/models"
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
