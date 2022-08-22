package handler

import (
	"go-hexagonal/service"
	"testing"

	"github.com/gofiber/fiber/v2"
)


func TestHandler_UpdateItem(t *testing.T) {
	type fields struct {
		cs service.CustomerService
		is service.ItemService
		os service.OrderService
	}
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Handler{
				cs: tt.fields.cs,
				is: tt.fields.is,
				os: tt.fields.os,
			}
			if err := h.UpdateItem(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Handler.UpdateItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
