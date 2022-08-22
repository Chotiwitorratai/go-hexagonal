package model

type CreateUser struct {
	Name string `json:"name"`
	DateOfBirth string `json:"bd"`
	City string `json:"city"`
	ZipCode string `json:"zipcode"`
	Status int `json:"status"`
}

type CartItem struct {
	Item Item
	Quantity int
}

type Cart struct {
	Items []CartItem
	Total float32
}

type CreateItem struct {
	Name string
	Price float32
	Stock int
}

type UpdateItem struct {
	Name string
	Price float32
	Stock int
}

type AddOrder struct {
	ItemID uint `json:"item_id"`
	Quantity uint `json:"qty"`
}
