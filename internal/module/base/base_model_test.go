package base

import (
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestModel_camelToSnack(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"test camel to snack", fields{}, args{"ssBBAccd"}, "ss_b_b_accd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Model{
				db: tt.fields.db,
			}
			if got := m.camelToSnack(tt.args.s); got != tt.want {
				t.Errorf("Model.camelToSnack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModel_ToDBFields(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		fields     []string
		omitFields []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{"test fieldmap to dbfield", fields{}, args{[]string{"title", "createAt", "episodes"}, []string{"episodes"}}, []string{"\"title\"", "\"create_at\""}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Model{
				db: tt.fields.db,
			}
			if got := m.toDBFields(tt.args.fields, tt.args.omitFields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Model.ToDBFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
