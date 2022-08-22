package handler

import (
	"encoding/json"
	"fmt"
	"go-hexagonal/model"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) GetCustomers(c *fiber.Ctx)error{
	customers ,err := h.cs.GetCustomers()
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"Error","msg":err,"data":""})
	}
	
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"OK","msg":"","data":customers})
}

func (h Handler) GetCustomer(c *fiber.Ctx)error{
	CustomerID := c.Params("CustomerID")
	id, err := strconv.ParseUint(CustomerID, 10, 32)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"Error","msg":"Customer ID must be number","data":""})
	}
	uid := uint(id)
	customer,err := h.cs.GetCustomer(uid)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"Error","msg":err,"data":""})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status":"OK","msg":"","data":customer})
}

func (h Handler) CreateCustomer(c *fiber.Ctx)error {
	user := &model.CreateUser{}
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"Error","msg":err,"data":""})
	}
	result,err := h.cs.CreateCustomer(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status":"Error","msg":err,"data":""})
	}
	body := c.Body()
	json.Unmarshal(body, user)
	fmt.Println(user)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status":"OK","msg":"","data":result})

}