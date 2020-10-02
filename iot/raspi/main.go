package main

import (
	"github.com/9d77v/pdc/iot/sdk"
	"github.com/9d77v/pdc/iot/sensors"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	iotSDK := sdk.NewIotSDK()
	r := raspi.NewAdaptor()
	iotSDK.Run([]func(){
		func() {
			sensors.BME280(iotSDK, r)
		},
		func() {
			sensors.BH1750(iotSDK, r, 0)
		},
	})
}
