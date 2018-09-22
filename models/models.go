package models

import (
//"github.com/graphql-go/graphql"
//"github.com/jinzhu/gorm"
)

type Shop struct {
	// gorm.Model `json:"id"`
	ID       uint      `json:"id"`
	Products []Product `json:"products"`
	Orders   []Order   `json:"orders"`
	Name     string    `json:"name" gorm:"unique;not null"`
}
type Product struct {
	// gorm.Model
	ID        uint       `json:"id"`
	LineItems []LineItem `json:"lineItems"`
	ShopID    uint       `json:"shopId"`
	Name      string     `json:"name"`
	Value     int        `json:"value"`
	Quantity  int        `json:"quantity"`
}
type Order struct {
	// gorm.Model
	ID        uint       `json:"id"`
	LineItems []LineItem `json:"lineItems"`
	ShopID    uint       `json:"shopId"`
}
type LineItem struct {
	// gorm.Model
	ID        uint `json:"id"`
	ProductID uint `json:"productId"`
	OrderID   uint `json:"orderId"`
	Quantity  int  `json:"quantity"`
}
