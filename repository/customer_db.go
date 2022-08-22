package repository

import (
	"go-hexagonal/model"
	"time"

	"gorm.io/gorm"
)
const (
    layoutISO = "2006-01-02"
)
type customerRepositoryDB struct {
	db *gorm.DB
}

func NewCustomerRepositoryDB(db *gorm.DB) CustomerRepository { 
	return customerRepositoryDB{db:db} 
}

func (r customerRepositoryDB)CreateCustomer(Name string, Birthday string,City string,ZipCode string,Status int)(*model.Customer, error) {
	customer := model.Customer{}
	time , _ := time.Parse(layoutISO, Birthday)
	customer.Name = Name
	customer.DateOfBirth = time
	customer.City = City
	customer.ZipCode = ZipCode
	customer.Status = Status
	err := r.db.Create(&customer).Error
	if err != nil {
		return nil , err
	}
	return &customer, nil
} 

func (r customerRepositoryDB)GetAll() ([]model.Customer, error) {
	customers := []model.Customer{}
	err := r.db.Find(&customers).Error
	if err != nil {
		return nil , err
	}
	return customers, nil

} 

func (r customerRepositoryDB)GetByID(id uint) (*model.Customer, error) {
	customer := model.Customer{}
	err := r.db.Find(&customer,id).Error
	if err != nil {
		return nil , err
	}
	return &customer, nil
} 