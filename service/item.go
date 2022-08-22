package service

import "go-hexagonal/model"


type ItemResponse struct {
	ItemID int  `json:"item_id"`
	Name string	`json:"name"`
	Price float32	`json:"price"`
	Stock int `json:"stock"`
}

type ItemService interface {
	GetItems() ([]ItemResponse,error)
	GetItem(id uint) (*model.Item,error)
	CreateItem(user *model.CreateItem) (*model.Item, error)
	GetStock(uint) (int,error)
	UpdateItem(uint,int) (*model.Item,error)
	UpdateItemDetail(*model.UpdateItem,uint) (*model.Item,error)
}