package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/9d77v/pdc/iot/sdk"
)

var iotSDK *sdk.IotSDK = sdk.NewIotSDK()

func main() {
	addr := sdk.GetEnvStr("DEVICE_PING", "")
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
		iotSDK.SetTelemetryConfig(uint32(deviceID), device.TelemetryConfig)
		qsub, err := iotSDK.SubscribeDeviceInfo(uint32(deviceID))
		if err != nil {
			log.Panicln("SubscribeDeviceInfo error:", err)
		}
		defer func() {
			qsub.Unsubscribe()
			qsub.Close()
		}()
	}
	go pingCheck()
	log.Printf("exiting (%v)", <-c)
}

func pingCheck() {

}
