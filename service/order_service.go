package service

import (
	"fmt"
	"go-hexagonal/model"
	"go-hexagonal/repository"
)
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}
var e *errorString
type orderService struct {
	or repository.OrderRepository 
}

func NewOrderService (or repository.OrderRepository) OrderService {
	return orderService{or: or}
} 

func (os orderService) Buy() (*model.Cart, error) {
	cart, err := os.or.GetOrder()
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}
	_ , err = os.or.ResetOrder()
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}
	return cart, nil
}
func (os orderService) AddItem(item_id uint,qty int) (*model.Order,error) {
	// item := &model.Item{}
	var order *model.Order
	have, err := os.or.GetOrderByItemID(item_id)
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}
	if have {
		order, err = os.or.UpdateOrder(item_id,qty)
		if err != nil {
		// log.Fatal(err)
		return nil, err
		}
	} else {
		order, err = os.or.CreateOrder(item_id,qty)
		if err != nil {
			// log.Fatal(err)
			return nil, err
		}
	}
	return order, nil
}
func (os orderService) Remove(id uint,qty int) (*model.Order, error) {
	order, err :=os.or.UpdateOrder(id,qty)
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}
	return order, nil
}
func (os orderService) Cart()(*model.Cart,error) {
	cart, err := os.or.GetOrder()
	test := "record not found"
	if err != nil {
		// log.Fatal(err)
		err := e.Error()
		if err == test{
		// fmt.Println(err)
		}
		return nil, nil
	}
	return cart, nil
}
func (os orderService) GetOrder(item_id uint)(*model.Order,error) {
	order, err :=os.or.GetOrderByID(item_id)
	if err != nil {
		fmt.Println(err)
		// log.Fatal(err)
		return nil, err
	}
	return order, nil
}
