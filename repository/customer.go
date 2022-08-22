package repository

import (
	"go-hexagonal/model"
)


type CustomerRepository interface {
	CreateCustomer(string, string, string, string, int)(*model.Customer, error)
	GetAll() ([]model.Customer, error)
	GetByID(uint) (*model.Customer, error) //struct can't return nil
}

