package sensors

import (
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"

	"github.com/9d77v/pdc/pkg/iot/sdk"
)

//SHT3x get temperature and humidity
func SHT3x(iotSDK *sdk.IotSDK, r Adaptor, bus ...int) {
	device := iotSDK.DeviceInfo
	if device == nil {
		return
	}
	sht3x := i2c.NewSHT3xDriver(r, i2c.WithBus(getBus(bus...)))
	work := func() {
		sht3x.Units = "C"
		err := sht3x.Start()
		if err != nil {
			log.Println("sht3x start error:", err)
			return
		}
		samplingFrequency := 5
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

		sht30Hz := device.AttributeConfig["sht30_hz"]
		if sht30Hz != 0 {
			attributeMap[sht30Hz] = fmt.Sprintf("%.2fHz", 1.0/float64(samplingFrequency))
		}
		if len(attributeMap) > 0 {
			iotSDK.SetDeviceAttributes(attributeMap)
		}
		gobot.Every(time.Duration(samplingFrequency)*time.Second, func() {
			temp, rh, err := sht3x.Sample()
			log.Printf("温度：%f,湿度：%f\n", temp, rh)
			if err != nil {
				log.Println("get sht3x telemetries error:", err)
			} else {
				telemetryMap := make(map[uint32]float64, 0)
				temperature := device.TelemetryConfig["temperature"]
				if temperature != 0 {
					telemetryMap[temperature] = float64(temp)
				}
				humidity := device.TelemetryConfig["humidity"]
				if humidity != 0 {
					telemetryMap[humidity] = float64(rh)
				}
				if len(telemetryMap) > 0 {
					now := ptypes.TimestampNow()
					iotSDK.SetDeviceTelemetries(telemetryMap, now)
				}
			}
		})

	}

	robot := gobot.NewRobot("SHT3xbot",
		[]gobot.Connection{r},
		[]gobot.Device{sht3x},
		work,
	)

	err := robot.Start()
	if err != nil {
		log.Println("sht3xbot start error")
	}
}
