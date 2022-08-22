package handler

import (
	"go-hexagonal/service"
	"reflect"
	"testing"
)

func TestNewHandler(t *testing.T) {
	type args struct {
		cs service.CustomerService
		is service.ItemService
		os service.OrderService
	}
	tests := []struct {
		name string
		args args
		want Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHandler(tt.args.cs, tt.args.is, tt.args.os); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
