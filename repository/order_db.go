package repository

import (
	"errors"
	"fmt"
	"go-hexagonal/model"

	"gorm.io/gorm"
)

type orderRepositoryDB struct {
	db *gorm.DB
}

func NewOrderRepositoryDB(db *gorm.DB) OrderRepository {
	return orderRepositoryDB{db:db}
}

func (r orderRepositoryDB) CreateOrder(item_id uint,qty int)(*model.Order,error) {
	order := &model.Order{}
	item := &model.Item{}
	err := r.db.First(item,item_id).Error
	if err != nil {
		return nil,err
	}
	order.Item = *item
	order.Quantity = qty
	order.ItemID = item_id
	err = r.db.Create(&order).Error
	if err != nil {
		return nil,err
	}
	return order,nil
}
func (r orderRepositoryDB) UpdateOrder( id_item uint, qty int)(*model.Order,error) {
	order := &model.Order{}
	err := r.db.First(&order,"item_id =?",id_item).Error
	if err != nil {
		return nil,err
	}
	err = r.db.Model(&order).Where("item_id =?",id_item).Update("Quantity",order.Quantity + qty).Error
	if err != nil {
		return nil,err
	}
	return order,nil
}
func (r orderRepositoryDB) DeleteOrder(item_id uint)(*model.Order,error) {
	order := &model.Order{ItemID:item_id}
	err := r.db.Delete(order,"item_id =?",item_id).Error
	if err != nil {
		return nil,err
	}
	return order,nil
}
func (r orderRepositoryDB) GetOrderByID(item_id uint)(*model.Order,error) {
	order := &model.Order{ItemID:item_id}
	err := r.db.Find(order,"item_id=?",item_id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &model.Order{Quantity:0},nil
		}
		return nil,err
	}
	return order,nil
}
func (r orderRepositoryDB) ResetOrder()(*model.Order,error) {
	orders := []model.Order{}
	order := model.Order{}
	err := r.db.Find(&orders).Error
	if err != nil {
		return nil,err
	}
	if len(orders) == 0 {
		return nil,nil
	}

	err = r.db.Preload("Item").Find(&orders).Delete(&orders).Error
	if err != nil {
		if(err == gorm.ErrRecordNotFound){
			return nil,nil
		}
		return nil,err
	}
	return &order,nil
}
func (r orderRepositoryDB) GetOrderByItemID(itemID uint)(bool,error) {
	order := &model.Order{}
	err := r.db.Where("item_id = ?",itemID).First(&order).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Error",errors.Is(err, gorm.ErrRecordNotFound))
			return false,nil
		}
		return false,err
	}
	return true,nil
}
func (r orderRepositoryDB) GetOrder()(*model.Cart,error) {
	orders := []model.Order{}
	items := []model.CartItem{}
	cart := &model.Cart{}
	err := r.db.Preload("Item").Find(&orders).Error
	if err != nil {
		return nil,err
	}
	if len(orders) == 0 {return nil,nil}
	for _, order := range orders {
		CartItem := &model.CartItem{Item:order.Item,Quantity:order.Quantity}
		items = append(items,*CartItem)
		cart.Total += float32(order.Quantity) * order.Item.Price
	}
	cart.Items = items
	return cart,nil
}