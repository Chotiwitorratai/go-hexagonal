package service

import "go-hexagonal/model"


type CustomerResponse struct {
	CustomerID int  `json:"customer_id"`
	Name string	`json:"name"`
	Status int	`json:"status"`
}

type CustomerService interface {
	GetCustomers() ([]CustomerResponse,error)
	GetCustomer(id uint) (*CustomerResponse,error)
	CreateCustomer(user *model.CreateUser) (*model.Customer, error)
}