package model

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	CustomerID int  `gorm:"primaryKey"`
	Name string
	DateOfBirth time.Time 
	City string
	ZipCode string
	Status int
	// ItemID uint
	// Items []Item
}

type Item struct {
	ItemID uint `gorm:"primaryKey"`
	Name string
	Price float32
	Stock int
}

type Order struct {
	gorm.Model
	ItemID uint 
	Item Item  `json:"item"`
	Quantity int `json:"quantity"`
}
