package utils

import (
	"os"
	"strconv"
)

//GetEnvStr ...
func GetEnvStr(key, value string) string {
	data := os.Getenv(key)
	if data == "" {
		return value
	}
	return data
}

//GetEnvInt ...
func GetEnvInt(key string, value int) int {
	data := os.Getenv(key)
	if data == "" {
		return value
	}
	parseData, err := strconv.Atoi(key)
	if err != nil {
		return value
	}
	return parseData
}

//GetEnvBool ...
func GetEnvBool(key string, value bool) bool {
	data := os.Getenv(key)
	if data == "" {
		return value
	}
	parseData, err := strconv.ParseBool(data)
	if err != nil {
		return value
	}
	return parseData
}
