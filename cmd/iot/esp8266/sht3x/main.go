package main

import (
	"fmt"

	"gobot.io/x/gobot/platforms/firmata"

	"github.com/9d77v/pdc/pkg/iot/sdk"
	"github.com/9d77v/pdc/pkg/iot/sensors"
)

func main() {
	iotSDK := sdk.NewIotSDK()
	works := []func(){
		func() {
			device := iotSDK.DeviceInfo
			if device == nil {
				return
			}
			r := firmata.NewTCPAdaptor(fmt.Sprintf("%s:%d", device.Ip, device.Port))
			sensors.SHT3x(iotSDK, r, 0)
		},
	}
	iotSDK.Run(works, nil, nil)
}
