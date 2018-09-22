package types

import "github.com/graphql-go/graphql"

type Product struct {
	Id       int `db:"id" json:"id,omitempty"`
	*Shop    `json:"shop"`
	Items    []*LineItem `json:"items"`
	Name     string      `db:"name" json:"name"`
	Price    int         `db:"price" json:"price"`
	Quantity int         `db:"quantity" json:"quantity"`
}

var ProductType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.Int},
		"name":     &graphql.Field{Type: graphql.String},
		"price":    &graphql.Field{Type: graphql.Int},
		"quantity": &graphql.Field{Type: graphql.Int},
	},
})
