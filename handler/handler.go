package handler

import (
	"go-hexagonal/service"

	"github.com/gofiber/fiber/v2"
)

var c *fiber.App

type Handler struct {
	cs service.CustomerService
	is service.ItemService
	os service.OrderService
}

func NewHandler(cs service.CustomerService, is service.ItemService, os service.OrderService) Handler {
	return Handler{cs: cs, is: is, os: os}
}
