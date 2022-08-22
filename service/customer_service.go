package service

import (
	"go-hexagonal/model"
	"go-hexagonal/repository"
	"log"
)

type customerService struct {
	cr repository.CustomerRepository 
}

func NewCustomerService (cr repository.CustomerRepository) CustomerService {
	return customerService{cr: cr}
} 

func (s customerService) GetCustomers() ([]CustomerResponse,error) {
	customers, err := s.cr.GetAll()
	if err != nil {
		log.Fatal(err)
		return nil,err
	}
	custResponses := []CustomerResponse{}
	for _, c := range customers {
		custResponse := CustomerResponse{
			CustomerID : c.CustomerID,
			Name : c.Name , 
			Status: c.Status,
		}
		custResponses = append(custResponses,custResponse)
	}
	return custResponses,nil

}

func (s customerService) GetCustomer(id uint) (*CustomerResponse,error) {
	customer,err := s.cr.GetByID(id)
	if err != nil {
		log.Fatal(err)
		return nil,err
	}
	custResponse := CustomerResponse{CustomerID:customer.CustomerID,Name:customer.Name,Status:customer.Status}
	return &custResponse,nil
}

func (s customerService) CreateCustomer(user *model.CreateUser) (*model.Customer, error) {
	customer,err := s.cr.CreateCustomer(user.Name,user.DateOfBirth,user.City,user.ZipCode,user.Status)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
