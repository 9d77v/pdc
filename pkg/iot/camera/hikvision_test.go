package camera

import (
	"reflect"
	"testing"
)

func TestHikvision_Capture(t *testing.T) {
	tests := []struct {
		name string
		h    *Hikvision
		want []byte
	}{
		{"test capture", NewHikvision("", "", ""), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.h.Capture()
			if got != nil {
				tt.h.saveImage(got)
			}
		})
	}
}

func TestHikvision_GetDeviceInfo(t *testing.T) {
	tests := []struct {
		name string
		h    *Hikvision
		want map[string]string
	}{
		{"test GetDeviceInfo", NewHikvision("", "", ""), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.GetDeviceInfo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hikvision.GetDeviceInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
