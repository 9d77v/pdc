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

//BH1750 get light data
func BH1750(iotSDK *sdk.IotSDK, r Adaptor, bus ...int) {
	device := iotSDK.DeviceInfo
	if device == nil {
		return
	}
	bh1750 := i2c.NewBH1750Driver(r, i2c.WithBus(getBus(bus...)))
	work := func() {
		err := bh1750.Start()
		if err != nil {
			log.Println("bh1750 start error:", err)
			return
		}
		samplingFrequency := 1
		attributeMap := make(map[uint32]string, 0)
		bh1750Hz := device.AttributeConfig["bh1750_hz"]
		if bh1750Hz != 0 {
			attributeMap[bh1750Hz] = fmt.Sprintf("%.2fHz", 1.0/float64(samplingFrequency))
		}
		if len(attributeMap) > 0 {
			iotSDK.SetDeviceAttributes(attributeMap)
		}
		gobot.Every(time.Duration(samplingFrequency)*time.Second, func() {
			illu, err := bh1750.Lux()
			if err != nil {
				log.Println("get bh1750 illuminance error:", err)
			} else {
				telemetryMap := make(map[uint32]float64, 0)
				illuminance := device.TelemetryConfig["illuminance"]
				if illuminance != 0 {
					telemetryMap[illuminance] = float64(illu)
				}
				if len(telemetryMap) > 0 {
					now := ptypes.TimestampNow()
					iotSDK.SetDeviceTelemetries(telemetryMap, now)
				}
			}
		})

	}

	robot := gobot.NewRobot("BH1750bot",
		[]gobot.Connection{r},
		[]gobot.Device{bh1750},
		work,
	)

	err := robot.Start()
	if err != nil {
		log.Println("BH1750bot start error")
	}
}
