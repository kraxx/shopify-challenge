/*
** Describes our database models, as well as graphQL schemas.
 */

package models

import "github.com/graphql-go/graphql"

// Database models
type Shop struct {
	ID       int       `json:"id"`
	Products []Product `json:"products"`
	Orders   []Order   `json:"orders"`
	Name     string    `json:"name" gorm:"unique;not null"`
}
type Product struct {
	ID        int        `json:"id"`
	LineItems []LineItem `json:"line_items"`
	ShopID    int        `json:"shop_id"`
	Name      string     `json:"name"`
	Value     int        `json:"value"`
	Quantity  int        `json:"quantity"`
}
type Order struct {
	ID        int        `json:"id"`
	LineItems []LineItem `json:"line_items"`
	ShopID    int        `json:"shop_id"`
}
type LineItem struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	OrderID   int `json:"order_id"`
	Quantity  int `json:"quantity"`
}

// Schema Object Types for graphQL
var ShopType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Shop",
	Description: "Shops have Products and Orders.",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
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
		"id":       &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
		"shop_id":  &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
		"name":     &graphql.Field{Type: graphql.String},
		"value":    &graphql.Field{Type: graphql.Int},
		"quantity": &graphql.Field{Type: graphql.Int},
		"line_items": &graphql.Field{
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
		"id":      &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
		"shop_id": &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
		"line_items": &graphql.Field{
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
		"id":         &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
		"product_id": &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
		"order_id":   &graphql.Field{Type: graphql.NewNonNull(graphql.Int)},
		"quantity":   &graphql.Field{Type: graphql.Int},
		"value": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return GetTotalLineItemPrice(params.Source.(LineItem).ProductID, params.Source.(LineItem).Quantity), nil
			},
		},
	},
})
