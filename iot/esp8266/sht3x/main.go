package main

import (
	"fmt"

	"github.com/9d77v/pdc/iot/sdk"
	"github.com/9d77v/pdc/iot/sensors"

	"gobot.io/x/gobot/platforms/firmata"
)

func main() {
	iotSDK := sdk.NewIotSDK()
	iotSDK.Run([]func(){
		func() {
			device := iotSDK.DeviceInfo
			if device == nil {
				return
			}
			r := firmata.NewTCPAdaptor(fmt.Sprintf("%s:%d", device.IP, device.Port))
			sensors.SHT3x(iotSDK, r)
		},
	})
}
