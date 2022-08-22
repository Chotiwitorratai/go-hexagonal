package handler

import (
	"go-hexagonal/model"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) CreateItem(c *fiber.Ctx) error {
	item := &model.CreateItem{}
	c.BodyParser(item)
	resItem, err := h.is.CreateItem(item)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "message": err, "data": ""})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK", "message": err, "data": resItem})
}

func (h Handler) GetCategory(c *fiber.Ctx) error {
	items, err := h.is.GetItems()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "message": err, "data": ""})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK", "message": err, "data": items})
}
func (h Handler) UpdateItem(c *fiber.Ctx) error {
	item := &model.Item{}
	c.BodyParser(item)
	updateItem := &model.UpdateItem{}
	updateItem.Name = item.Name
	updateItem.Price = item.Price
	updateItem.Stock = item.Stock

	items, err := h.is.UpdateItemDetail(updateItem, uint(item.ItemID))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "message": err, "data": ""})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK", "message": err, "data": items})
}
