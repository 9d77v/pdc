package camera

import (
	"reflect"
	"testing"
)

func TestDahua_Capture(t *testing.T) {
	tests := []struct {
		name string
		d    *Dahua
		want []byte
	}{
		{"test capture", NewDahua("", "", ""), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Capture(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dahua.Capture() = %v, want %v", got, tt.want)
			}
		})
	}
}
