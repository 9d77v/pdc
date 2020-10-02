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
			sensors.SHT3x(iotSDK, r)
		},
	})
}
