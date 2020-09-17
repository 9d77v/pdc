package sensors

import (
	"fmt"
	"log"
	"time"

	"github.com/9d77v/pdc/iot/sdk"
	"github.com/9d77v/pdc/iot/sdk/pb"
	"github.com/golang/protobuf/ptypes"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
)

//BH1750 get light data
func BH1750(device *pb.DeviceInfo, iotSDK *sdk.IotSDK, r Adaptor) {
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
			iotSDK.SetDeviceAttributes(device.ID, attributeMap)
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
					iotSDK.SetDeviceTelemetries(device.ID, telemetryMap, now)
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
