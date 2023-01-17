package config

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		env string
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(tt.args.env); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
