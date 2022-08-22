package router

import (
	"go-hexagonal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App, h handler.Handler) {
	customer := app.Group("/customer")
	item := app.Group("/item")
	order := app.Group("/order")

	customer.Get("/list", h.GetCustomers)
	customer.Get("/user/:CustomerID", h.GetCustomer)
	customer.Post("/create", h.CreateCustomer)

	item.Post("/create", h.CreateItem)
	item.Get("/get/all", h.GetCategory)
	item.Put("/update", h.UpdateItem)

	order.Get("/cart", h.GetCart)
	order.Post("/add", h.AddCart)
	order.Post("/buy", h.Buy)
}
