package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"

	"github.com/9d77v/pdc/iot/sdk"
	"github.com/golang/protobuf/ptypes"
	cron "github.com/robfig/cron/v3"
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
	deviceMap := make(map[uint32]string)
	for _, v := range addresses {
		deviceID, _ := strconv.Atoi(v)
		device, err := iotSDK.GetDeviceInfo(uint32(deviceID))
		if err != nil {
			log.Println("GetDeviceInfo Failed:", err)
			continue
		}
		deviceMap[uint32(deviceID)] = device.GetIP()
	}
	cr := cron.New()
	cr.AddFunc("*/1 * * * *", func() {
		pingCheck(deviceMap)
	})
	cr.Start()
	log.Printf("exiting (%v)", <-c)
	cr.Stop()
}

func pingCheck(deviceMap map[uint32]string) {
	for k, v := range deviceMap {
		k, v := k, v
		go func(k uint32, v string) {
			now := ptypes.TimestampNow()
			var flag uint32
			if ping(v) {
				flag = 1
			}
			iotSDK.SetDeviceHealth(k, flag, now)
		}(k, v)
	}
}

func ping(ip string) bool {
	var buf bytes.Buffer
	var errorBuf bytes.Buffer
	cmd := exec.Command("ping", "-i", "1", "-c", "3", ip)
	cmd.Stdout = &buf
	cmd.Stderr = &errorBuf
	cmd.Run()
	if buf.String() != "" {
		data := buf.String()
		dataArr := strings.Split(data, "\n")
		if len(dataArr) != 9 {
			return false
		}
		statisticsStr := dataArr[6]
		statisticsArr := strings.Split(statisticsStr, ",")
		if len(statisticsArr) != 4 {
			return false
		}
		if statisticsArr[2] != " 0% packet loss" {
			return false
		}
		return true
	}
	return false
}
