package main

import (
	"gobot.io/x/gobot/platforms/raspi"

	"github.com/9d77v/pdc/pkg/iot/sdk"
	"github.com/9d77v/pdc/pkg/iot/sensors"
)

func main() {
	iotSDK := sdk.NewIotSDK()
	r := raspi.NewAdaptor()
	works := []func(){
		func() {
			sensors.SHT3x(iotSDK, r)
		},
	}
	iotSDK.Run(works, nil, nil)
}
