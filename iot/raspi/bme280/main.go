package main

import (
	"github.com/9d77v/pdc/iot/sdk"
	"github.com/9d77v/pdc/iot/sensors"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	iotSDK := sdk.NewIotSDK()
	r := raspi.NewAdaptor()
	works := []func(){
		func() {
			sensors.BME280(iotSDK, r)
		},
	}
	iotSDK.Run(works, nil)
}
