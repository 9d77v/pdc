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
	addr := sdk.GetEnvStr("DEVICE_RASPI_SHT3X", "")
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
		go sensors.SHT3x(device, iotSDK, r)
	}
	log.Printf("exiting (%v)", <-c)
}