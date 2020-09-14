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
	addr := os.Getenv("DEVICE_ESP8266_SHT3X")
	if len(addr) == 0 {
		return
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	addresses := strings.Split(addr, ",")
	for _, v := range addresses {
		deviceID, _ := strconv.Atoi(v)
		go sht3x(uint32(deviceID))
	}
	log.Printf("exiting (%v)", <-c)
}

//sht3x get temperature and humidity
func sht3x(deviceID uint32) {
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
	sht3x := i2c.NewSHT3xDriver(r)
	work := func() {
		sht3x.Units = "C"
		sht3x.Start()
		sn, err := sht3x.SerialNumber()
		if err != nil {
			log.Println("get sht3x serial number error:", err)
			return
		}
		attributeMap := make(map[uint32]string, 0)
		sht30SN := device.AttributeConfig["sht30_sn"]
		if sht30SN != 0 {
			attributeMap[sht30SN] = fmt.Sprintf("0x%08x", sn)
		}
		samplingFrequency := 5
		sht30Hz := device.AttributeConfig["sht30_hz"]
		if sht30Hz != 0 {
			attributeMap[sht30Hz] = fmt.Sprintf("%.2fHz", 1.0/float64(samplingFrequency))
		}
		if len(attributeMap) > 0 {
			iotSDK.SetDeviceAttributes(deviceID, attributeMap)
		}
		gobot.Every(time.Duration(samplingFrequency)*time.Second, func() {
			temp, rh, err := sht3x.Sample()
			if err != nil {
				log.Println("get sht3x telemetries error:", err)
			} else {
				telemetryRawMap := make(map[uint32]float64, 0)
				telemetryMap := make(map[uint32]float64, 0)
				config := iotSDK.GetTelemetryConfig(deviceID)
				temperature := config["temperature"]
				if temperature != nil {
					telemetryRawMap[temperature.ID] = float64(temp)
					formatTemp, _ := strconv.ParseFloat(strconv.FormatFloat(float64(temp)*temperature.Factor, 'f', int(temperature.Scale), 64), 64)
					telemetryMap[temperature.ID] = formatTemp
				}
				humidity := config["humidity"]
				if humidity != nil {
					telemetryRawMap[humidity.ID] = float64(rh)
					formatHumidity, _ := strconv.ParseFloat(strconv.FormatFloat(float64(rh)*humidity.Factor, 'f', int(humidity.Scale), 64), 64)
					telemetryMap[temperature.ID] = formatHumidity
				}
				if len(telemetryRawMap) > 0 {
					now := ptypes.TimestampNow()
					iotSDK.UploadDeviceTelemetries(deviceID, telemetryRawMap, now)
				}
			}
		})

	}

	robot := gobot.NewRobot("SHT3xbot",
		[]gobot.Connection{r},
		[]gobot.Device{sht3x},
		work,
	)

	robot.Start()
}
