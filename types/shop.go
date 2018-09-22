package types

import "github.com/graphql-go/graphql"

type Shop struct {
	Id       int        `db:"id" json:"id,omitempty"`
	Products []*Product `json:"products"`
	Orders   []*Order   `json:"orders"`
	Name     string     `db:"name" json:"name"`
}

var ShopType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Shop",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.Int},
		"name": &graphql.Field{Type: graphql.String},
    "products": &graphql.Field{
      Type: graphql.NewList(ProductType),
      Resolve: func(params graphql.ResolveParams) (interface{}, error) }{
        var products []Product

        //shopId := params.Source.(Product).Id
        // logic to get products
        return products, nil
      }
    },
	},
})

//

/*


type Order struct {
	Id    int `json:"id,omitempty"`
	*Shop `json:"shop"`
	Items []*LineItem `json:"items"`
}

type LineItem struct {
	Id       int `json:"id,omitempty"`
	*Product `json:"product"`
	*Order   `json:"order"`
	Quantity int `json:"quantity"`
}



*/
