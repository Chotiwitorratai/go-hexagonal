package service

import "go-hexagonal/model"

type OrderService interface {
	Buy() (*model.Cart,error)
	AddItem(uint,int) (*model.Order,error)
	Remove(uint,int) (*model.Order, error)
	Cart()(*model.Cart,error)
	GetOrder(uint)(*model.Order,error)
}