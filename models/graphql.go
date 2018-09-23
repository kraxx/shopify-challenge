package models

import (
	"../db"
	// "fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

var DB *gorm.DB

func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"shop":       GetShopQuery(),
		"createShop": CreateShopMutation(),
		"updateShop": UpdateShopMutation(),
		"deleteShop": DeleteShopMutation(),
	}
}
func UpdateShopMutation() *graphql.Field {
	return &graphql.Field{
		Type:        ShopType,
		Description: "Update existing shop by ID",
		Args: graphql.FieldConfigArgument{
			"id":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
			"name": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			// var shop Shop
			return UpdateGeneric(params.Args, &Shop{})
			// return shop, nil
		},
	}
}
func DeleteShopMutation() *graphql.Field {
	return &graphql.Field{
		Type:        ShopType,
		Description: "Delete existing shop",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			// DeleteGeneric(params.Args, &Shop{})
			// return Shop{}, nil
			return DeleteGeneric(params.Args, &Shop{})
		},
	}
}
func CreateShopMutation() *graphql.Field {
	return &graphql.Field{
		Type:        ShopType,
		Description: "Create new shop",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			shop := Shop{
				Name: params.Args["name"].(string),
			}
			CreateGeneric(&shop)
			return shop, nil
		},
	}
}
func GetShopQuery() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(ShopType),
		Description: "Fetch existing shop",
		Args: graphql.FieldConfigArgument{
			"id":   &graphql.ArgumentConfig{Type: graphql.ID},
			"name": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			// var shops []Shop
			// GetGeneric(params.Args, &shops)
			// return shops, nil
			return GetGeneric(params.Args, &[]Shop{})
		},
	}
}

// func GetGeneric(args map[string]interface{}, entry interface{}) {
// 	DB.Where(args).Find(entry)
// }
func GetGeneric(args map[string]interface{}, entry interface{}) (interface{}, error) {
	DB.Where(args).Find(entry)
	return entry, nil
}
func GetChildrenGeneric(parent, children interface{}) (interface{}, error) {
	DB.Model(parent).Related(children)
	return children, nil
}
func CreateGeneric(entry interface{}) {
	DB.Create(entry)
}
func UpdateGeneric(args map[string]interface{}, entry interface{}) (interface{}, error) {
	DB.First(entry, args["id"]).Updates(args)
	return entry, nil
}
func DeleteGeneric(args map[string]interface{}, entry interface{}) (interface{}, error) {
	DB.Where(args).Delete(entry)
	return entry, nil
}

var ShopType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Shop",
	Description: "Shops have Products and Orders",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"name": &graphql.Field{Type: graphql.String},
		"products": &graphql.Field{
			Type: graphql.NewList(ProductType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// var products []Product
				// DB.Model(params.Source.(Shop)).Related(&products)
				// GetChildrenGeneric(params.Source.(Shop), &products)
				// return products, nil
				return GetChildrenGeneric(params.Source.(Shop), &[]Product{})
			},
		},
		"orders": &graphql.Field{
			Type: graphql.NewList(OrderType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// var orders []Order
				// DB.Model(params.Source.(Shop)).Related(&orders)
				// return orders, nil
				return GetChildrenGeneric(params.Source.(Shop), &[]Order{})
			},
		},
	},
})
var ProductType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Product",
	Description: "Products have LineItems",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"shopId":   &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"name":     &graphql.Field{Type: graphql.String},
		"value":    &graphql.Field{Type: graphql.Int},
		"quantity": &graphql.Field{Type: graphql.Int},
		"lineItems": &graphql.Field{
			Type: graphql.NewList(LineItemType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var lineItems []LineItem
				DB.Model(params.Source.(Product)).Related(&lineItems)
				return lineItems, nil
			},
		},
	},
})
var OrderType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Order",
	Description: "Orders have LineItems.",
	Fields: graphql.Fields{
		"id":     &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"shopId": &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"lineItems": &graphql.Field{
			Type: graphql.NewList(LineItemType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var lineItems []LineItem
				DB.Model(params.Source.(Order)).Related(&lineItems)
				return lineItems, nil
			},
		},
		"value": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var lineItems []LineItem
				var sum int
				DB.Model(params.Source.(Order)).Related(&lineItems)
				for _, item := range lineItems {
					var product Product
					DB.First(&product, item.ProductID)
					sum += product.Value * item.Quantity
				}
				return sum, nil
			},
		},
	},
})
var LineItemType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "LineItem",
	Description: "A list item in an order. Product type and amount",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"productId": &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"orderId":   &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
		"quantity":  &graphql.Field{Type: graphql.Int},
		"value": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var product Product
				DB.First(&product, params.Source.(LineItem).ProductID)
				return product.Value * params.Source.(LineItem).Quantity, nil
			},
		},
	},
})

var GraphqlHandler http.Handler

func init() {
	DB = db.DB
	DB.AutoMigrate(&Shop{}, &Product{}, &Order{}, &LineItem{})

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: GetRootFields(),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: GetRootFields(),
		}),
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create new schema, error: %v", err)
	}
	GraphqlHandler = handler.New(&handler.Config{
		Schema: &schema,
	})
}
