package repository

import "go-hexagonal/model"

type OrderRepository interface {
	CreateOrder(uint, int)(*model.Order,error) //Create Order
	UpdateOrder(uint, int)(*model.Order,error) //Update Order เมื่อมี Item นั้นอยู่ใน Order แล้ว
	DeleteOrder(uint)(*model.Order,error) //ใช้ในการนำ Order ออก
	ResetOrder()(*model.Order,error)//ใช้ในการ Buy
	GetOrder()(*model.Cart,error) //Get Cart
	GetOrderByID(uint)(*model.Order,error) //Get Order by ID
	GetOrderByItemID(uint)(bool,error) //Get Order by item ID
}