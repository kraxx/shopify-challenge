package models

import "github.com/graphql-go/graphql"

type Shop struct {
	ID       uint      `json:"id"`
	Products []Product `json:"products"`
	Orders   []Order   `json:"orders"`
	Name     string    `json:"name" gorm:"unique;not null"`
}
type Product struct {
	ID        uint       `json:"id"`
	LineItems []LineItem `json:"lineItems"`
	ShopID    uint       `json:"shopId"`
	Name      string     `json:"name"`
	Value     int        `json:"value"`
	Quantity  int        `json:"quantity"`
}
type Order struct {
	ID        uint       `json:"id"`
	LineItems []LineItem `json:"lineItems"`
	ShopID    uint       `json:"shopId"`
}
type LineItem struct {
	ID        uint `json:"id"`
	ProductID uint `json:"productId"`
	OrderID   uint `json:"orderId"`
	Quantity  int  `json:"quantity"`
}

var ShopType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Shop",
	Description: "Shops have Products and Orders.",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"name": &graphql.Field{Type: graphql.String},
		"products": &graphql.Field{
			Type: graphql.NewList(ProductType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return GetChildrenGeneric(params.Source.(Shop), &[]Product{})
			},
		},
		"orders": &graphql.Field{
			Type: graphql.NewList(OrderType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return GetChildrenGeneric(params.Source.(Shop), &[]Order{})
			},
		},
	},
})
var ProductType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Product",
	Description: "One shop's products. Products are related to LineItems currently on an Order.",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"shopId":   &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"name":     &graphql.Field{Type: graphql.String},
		"value":    &graphql.Field{Type: graphql.Int},
		"quantity": &graphql.Field{Type: graphql.Int},
		"lineItems": &graphql.Field{
			Type: graphql.NewList(LineItemType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return GetChildrenGeneric(params.Source.(Product), &[]LineItem{})
			},
		},
	},
})
var OrderType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Order",
	Description: "Orders consist of LineItems from one shop.",
	Fields: graphql.Fields{
		"id":     &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"shopId": &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"lineItems": &graphql.Field{
			Type: graphql.NewList(LineItemType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return GetChildrenGeneric(params.Source.(Order), &[]LineItem{})
			},
		},
		"value": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return GetTotalOrderPrice(params.Source.(Order)), nil
			},
		},
	},
})
var LineItemType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "LineItem",
	Description: "A list item in an order. Product type and amount.",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"productId": &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"orderId":   &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"quantity":  &graphql.Field{Type: graphql.Int},
		"value": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return GetTotalLineItemPrice(params.Source.(LineItem).ProductID, params.Source.(LineItem).Quantity), nil
			},
		},
	},
})
