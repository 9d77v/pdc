package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/9d77v/pdc/iot/sdk"
	"github.com/golang/protobuf/ptypes"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/firmata"
)

var iotSDK *sdk.IotSDK = sdk.NewIotSDK()

func main() {
	addr := sdk.GetEnvStr("DEVICE_ESP8266_BH1750", "2")
	if len(addr) == 0 {
		return
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	addresses := strings.Split(addr, ",")
	for _, v := range addresses {
		deviceID, _ := strconv.Atoi(v)
		go bh1750(uint32(deviceID))
	}
	log.Printf("exiting (%v)", <-c)
}

//bh1750 get light data
func bh1750(deviceID uint32) {
	device, err := iotSDK.GetDeviceInfo(deviceID)
	if err != nil {
		log.Println("GetDeviceInfo Failed:", err)
		return
	}
	iotSDK.SetTelemetryConfig(deviceID, device.TelemetryConfig)
	qsub, err := iotSDK.SubscribeDeviceInfo(deviceID)
	if err != nil {
		log.Panicln("SubscribeDeviceInfo error:", err)
	}
	defer func() {
		qsub.Unsubscribe()
		qsub.Close()
	}()

	r := firmata.NewTCPAdaptor(fmt.Sprintf("%s:%d", device.IP, device.Port))
	bh1750 := i2c.NewBH1750Driver(r)
	work := func() {
		bh1750.Start()
		attributeMap := make(map[uint32]string, 0)
		samplingFrequency := 1
		bh1750Hz := device.AttributeConfig["bh1750_hz"]
		if bh1750Hz != 0 {
			attributeMap[bh1750Hz] = fmt.Sprintf("%.2fHz", 1.0/float64(samplingFrequency))
		}
		if len(attributeMap) > 0 {
			iotSDK.SetDeviceAttributes(deviceID, attributeMap)
		}
		gobot.Every(time.Duration(samplingFrequency)*time.Second, func() {
			illu, err := bh1750.Lux()
			if err != nil {
				log.Println("get bh1750 illuminance error:", err)
			} else {
				telemetryRawMap := make(map[uint32]float64, 0)
				telemetryMap := make(map[uint32]float64, 0)
				config := iotSDK.GetTelemetryConfig(deviceID)
				illuminance := config["illuminance"]
				if illuminance != nil {
					telemetryRawMap[illuminance.ID] = float64(illu)
					formatTemp, _ := strconv.ParseFloat(strconv.FormatFloat(float64(illu)*illuminance.Factor, 'f', int(illuminance.Scale), 64), 64)
					telemetryMap[illuminance.ID] = formatTemp
				}
				if len(telemetryRawMap) > 0 {
					now := ptypes.TimestampNow()
					iotSDK.UploadDeviceTelemetries(deviceID, telemetryRawMap, now)
				}
			}
		})

	}

	robot := gobot.NewRobot("BH1750bot",
		[]gobot.Connection{r},
		[]gobot.Device{bh1750},
		work,
	)

	robot.Start()
}
