package utils

import (
	"reflect"
	"testing"
)

func Test_CamelToSnack(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test camel to snack", args{"ssBBAccd"}, "ss_b_b_accd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelToSnack(tt.args.s); got != tt.want {
				t.Errorf("CamelToSnack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDBFields(t *testing.T) {
	type args struct {
		fields     []string
		omitFields []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"test fieldmap to dbfield", args{[]string{"title", "createAt", "episodes"}, []string{"episodes"}}, []string{"\"title\"", "\"create_at\""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToDBFields(tt.args.fields, tt.args.omitFields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToDBFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
