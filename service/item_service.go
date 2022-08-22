package service

import (
	"fmt"
	"go-hexagonal/model"
	"go-hexagonal/repository"
	"log"
)

type itemService struct {
	ir repository.ItemRepository
}

func NewItemService(ir repository.ItemRepository) ItemService {
	return itemService{ir: ir}
}

func (is itemService) GetItems() ([]ItemResponse, error) {
	items, err := is.ir.GetAllItem()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println(items)
	resItems := []ItemResponse{}
	for _, item := range items {
		resItem := ItemResponse{
			ItemID: int(item.ItemID),
			Name:   item.Name,
			Price:  float32(item.Price),
			Stock:  int(item.Stock),
		}
		resItems = append(resItems, resItem)
	}
	return resItems, nil
}
func (is itemService) GetItem(id uint) (*model.Item, error) {
	item, err := is.ir.GetItemByID(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return item, nil
}
func (is itemService) CreateItem(ci *model.CreateItem) (*model.Item, error) {
	item, err := is.ir.CreateItem(ci)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return item, nil
}
func (is itemService) GetStock(id uint) (int, error) {
	stock, err := is.ir.GetStockByID(id)
	if err != nil {
		return 0, err
	}
	return stock, nil
}
func (is itemService) UpdateItem(item_id uint, qty int) (*model.Item, error) {
	stock, err := is.ir.GetStockByID(item_id)
	if err != nil {
		return nil, err
	}
	item, err := is.ir.UpdateItem(item_id, stock-qty)
	if err != nil {
		return nil, err
	}
	return item, nil
}
func (is itemService) UpdateItemDetail(update *model.UpdateItem, item_id uint) (*model.Item, error) {
	fmt.Println("item_id", item_id)
	// fmt.Println(reflect.ValueOf(item_id).Kind())
	item, err := is.ir.UpdateItemDetail(update, uint(item_id))
	if err != nil {
		return nil, err
	}
	return item, nil
}
