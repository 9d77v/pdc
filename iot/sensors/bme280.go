package sensors

import (
	"fmt"
	"time"

	"github.com/9d77v/pdc/iot/sdk"
	"github.com/golang/protobuf/ptypes"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
)

//BME280 get temperature and humidity
func BME280(iotSDK *sdk.IotSDK, r Adaptor, bus ...int) {
	device := iotSDK.DeviceInfo
	if device == nil {
		return
	}
	bme280 := i2c.NewBME280Driver(r, i2c.WithBus(getBus(bus...)), i2c.WithAddress(0x76))
	work := func() {
		samplingFrequency := 1
		bme280Hz := device.AttributeConfig["bme280_hz"]
		attributeMap := make(map[uint32]string, 0)
		if bme280Hz != 0 {
			attributeMap[bme280Hz] = fmt.Sprintf("%.2fHz", 1.0/float64(samplingFrequency))
		}
		if len(attributeMap) > 0 {
			iotSDK.SetDeviceAttributes(attributeMap)
		}
		gobot.Every(time.Duration(samplingFrequency)*time.Second, func() {
			ok := true
			t, err1 := bme280.Temperature()
			if err1 != nil {
				ok = false
				fmt.Println(err1)
			}
			h, err4 := bme280.Humidity()
			if err4 != nil {
				ok = false
				fmt.Println(err4)
			}
			p, err2 := bme280.Pressure()
			if err2 != nil {
				ok = false
				fmt.Println(err2)
			}
			a, err3 := bme280.Altitude()
			if err3 != nil {
				ok = false
				fmt.Println(err3)
			}
			if ok {
				telemetryMap := make(map[uint32]float64, 0)
				temperature := device.TelemetryConfig["temperature"]
				if temperature != 0 {
					telemetryMap[temperature] = float64(t)
				}
				humidity := device.TelemetryConfig["humidity"]
				if humidity != 0 {
					telemetryMap[humidity] = float64(h)
				}
				pressure := device.TelemetryConfig["pressure"]
				if pressure != 0 {
					telemetryMap[pressure] = float64(p)
				}
				altitude := device.TelemetryConfig["altitude"]
				if altitude != 0 {
					telemetryMap[altitude] = float64(a)
				}
				if len(telemetryMap) > 0 {
					now := ptypes.TimestampNow()
					iotSDK.SetDeviceTelemetries(telemetryMap, now)
				}
			}
		})

	}

	robot := gobot.NewRobot("BME280bot",
		[]gobot.Connection{r},
		[]gobot.Device{bme280},
		work,
	)

	robot.Start()
}
