package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/9d77v/pdc/iot/sdk"
	"github.com/9d77v/pdc/iot/sensors"
	"gobot.io/x/gobot/platforms/raspi"
)

var iotSDK *sdk.IotSDK = sdk.NewIotSDK()

func main() {
	addr := sdk.GetEnvStr("DEVICE_RASPI", "4")
	if len(addr) == 0 {
		return
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	addresses := strings.Split(addr, ",")
	for _, v := range addresses {
		deviceID, _ := strconv.Atoi(v)
		device, err := iotSDK.GetDeviceInfo(uint32(deviceID))
		if err != nil {
			log.Println("GetDeviceInfo Failed:", err)
			return
		}
		r := raspi.NewAdaptor()
		go sensors.BME280(device, iotSDK, r)
		go sensors.BH1750(device, iotSDK, r, 0)
	}
	log.Printf("exiting (%v)", <-c)
}
