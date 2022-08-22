package repository

import (
	"go-hexagonal/model"
)


type ItemRepository interface {
	CreateItem(*model.CreateItem)(*model.Item, error) //สร้าง
	GetAllItem() ([]model.Item, error) //Get All
	GetItemByID(uint) (*model.Item, error) //Get /w ID
	UpdateItem(uint,int) (*model.Item, error) //ลด stock
	GetStockByID(uint) (int, error) //Get Stock
	UpdateItemDetail(*model.UpdateItem,uint) (*model.Item, error) //Update Item Detail
}