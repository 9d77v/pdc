package esp8266

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/9d77v/pdc/iot/sdk"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/firmata"
)

var iotSDK *sdk.IotSDK = sdk.NewIotSDK()

//CollectData ...
func CollectData(ctx context.Context) {
	addr := os.Getenv("DEVICE_ESP8266")
	if len(addr) == 0 {
		return
	}
	addresses := strings.Split(addr, ",")
	for _, v := range addresses {
		deviceID, _ := strconv.Atoi(v)
		go Sht3x(ctx, deviceID)
	}

	ticker := time.NewTicker(500 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}

//Sht3x get temperature and humidity
func Sht3x(ctx context.Context, deviceID int) {
	device, err := iotSDK.GetDeviceInfo(deviceID)
	if err != nil {
		log.Println("GetDeviceInfo Failed:", err)
		return
	}
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
		samplingFrequency := 5
		attributeMap := map[string]string{
			"sht30_sn": fmt.Sprintf("0x%08x", sn),
			"sht30_hz": fmt.Sprintf("%dHz", samplingFrequency/5),
		}
		iotSDK.SetDeviceAttributes(deviceID, attributeMap)
		gobot.Every(time.Duration(samplingFrequency)*time.Second, func() {
			temp, rh, err := sht3x.Sample()
			if err != nil {
				log.Println("get sht3x telemetries error:", err)
			} else {
				telemetryMap := map[string]float64{
					"temperature": float64(temp),
					"humidity":    float64(rh),
				}
				iotSDK.UploadDeviceTelemetries(deviceID, telemetryMap)
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
