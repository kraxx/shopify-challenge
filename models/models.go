package models

import (
	"github.com/jinzhu/gorm"
)

type Shop struct {
	gorm.Model
	Products []Product `json:"products"`
	Orders   []Order   `json:"orders"`

	Name string `json:"name" gorm:"unique;not null"`
}
type Product struct {
	gorm.Model
	// Shop      *Shop       `json:"shop"`
	LineItems []LineItem `json:"lineItems"`

	ShopID   uint   `json:"shopId"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}
type Order struct {
	gorm.Model
	// Shop      *Shop       `json:"shop"`
	LineItems []LineItem `json:"lineItems"`

	ShopID uint `json:"shopId"`
}
type LineItem struct {
	gorm.Model
	// Product *Product `json:"product"`
	// Order   *Order   `json:"order"`

	ProductID uint `json:"productId"`
	OrderID   uint `json:"orderId"`
	Quantity  int  `json:"quantity"`
}
