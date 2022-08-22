package handler

import (
	"go-hexagonal/model"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetCart(c *fiber.Ctx) error {
	cart, err := h.os.Cart()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "msg": err, "data": ""})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK", "msg": "", "data": cart})
}

func (h Handler) AddCart(c *fiber.Ctx) error {
	addOrder := model.AddOrder{}
	err := c.BodyParser(&addOrder)
	qty := addOrder.Quantity
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "msg": err, "data": ""})
	}
	if qty < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "msg": "Please send positive number", "data": ""})
	}
	order, err := h.os.GetOrder(addOrder.ItemID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "msg": err, "data": ""})
	}
	stock, error := h.is.GetStock(addOrder.ItemID)
	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "msg": err, "data": ""})
	}
	//Check if order > Stock

	if stock < int(qty) || stock < int(qty)+order.Quantity {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "msg": "Out of stock", "data": ""})
	}
	_, err = h.os.AddItem(addOrder.ItemID, int(qty))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "msg": err, "data": ""})
	}
	cart, err := h.os.Cart()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "msg": err, "data": ""})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK", "msg": "", "data": cart})
}

func (h Handler) Buy(c *fiber.Ctx) error {
	order, err := h.os.Buy()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "msg": err, "data": ""})
	}
	//ลด stock item
	for _, r := range order.Items {
		_, err = h.is.UpdateItem(r.Item.ItemID, r.Quantity)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Error", "msg": err, "data": ""})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "OK", "msg": "", "data": order})

}
