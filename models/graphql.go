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
func CustomGetQuery(objectType graphql.Output, objectArgs graphql.FieldConfigArgument, object interface{}) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(objectType),
		Args: objectArgs,
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return GetDatabaseEntry(params.Args, object)
		},
	}
}
func CustomUpdateMutation(objectType graphql.Output, objectArgs graphql.FieldConfigArgument, object interface{}) *graphql.Field {
	return &graphql.Field{
		Type: objectType,
		Args: objectArgs,
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return UpdateDatabaseEntry(params.Args, object)
		},
	}
}
func CustomDeleteMutation(objectType graphql.Output, objectArgs graphql.FieldConfigArgument, object interface{}) *graphql.Field {
	return &graphql.Field{
		Type: objectType,
		Args: objectArgs,
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			return DeleteDatabaseEntry(params.Args, object)
		},
	}
}
func GetShopQuery() *graphql.Field {
	args := graphql.FieldConfigArgument{
		"id":   &graphql.ArgumentConfig{Type: graphql.ID},
		"name": &graphql.ArgumentConfig{Type: graphql.String},
	}
	return CustomGetQuery(ShopType, args, &[]Shop{})
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
			CreateDatabaseEntry(&shop)
			return shop, nil
		},
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

			// Have to instantiate outside DB call to properly return children

			// return UpdateDatabaseEntry(params.Args, &Shop{})

			// var shop Shop
			// return UpdateDatabaseEntry(params.Args, &shop)

			var entry Shop
			DB.First(&entry, params.Args["id"]).Updates(params.Args)
			return entry, nil
		},
	}
	// args := graphql.FieldConfigArgument{
	// 	"id":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
	// 	"name": &graphql.ArgumentConfig{Type: graphql.String},
	// }
	// return CustomUpdateMutation(ShopType, args, &Shop{})
}
func DeleteShopMutation() *graphql.Field {
	return &graphql.Field{
		Type:        ShopType,
		Description: "Delete existing shop",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			// return DeleteDatabaseEntry(params.Args, &Shop{})

			var entry Shop
			DB.Where(params.Args).Delete(&entry)
			return entry, nil
		},
	}
	// args := graphql.FieldConfigArgument{
	// 	"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
	// }
	// return CustomDeleteMutation(ShopType, args, &Shop{})
}

func GetDatabaseEntry(args map[string]interface{}, entry interface{}) (interface{}, error) {
	DB.Where(args).Find(entry)
	return entry, nil
}
func CreateDatabaseEntry(entry interface{}) {
	DB.Create(entry)
}
func UpdateDatabaseEntry(args map[string]interface{}, entry interface{}) (interface{}, error) {
	DB.First(entry, args["id"]).Updates(args)
	return entry, nil
}
func DeleteDatabaseEntry(args map[string]interface{}, entry interface{}) (interface{}, error) {
	DB.Where(args).Delete(entry)
	return entry, nil
}

func GetChildrenGeneric(parent, children interface{}) (interface{}, error) {
	DB.Model(parent).Related(children)
	return children, nil
}

func GetTotalOrderPrice(order Order) int {
	var lineItems []LineItem
	var sum int
	DB.Model(order).Related(&lineItems)
	for _, item := range lineItems {
		sum += GetTotalLineItemPrice(item.ProductID, item.Quantity)
	}
	return sum
}

func GetTotalLineItemPrice(id uint, quantity int) int {
	var product Product
	DB.First(&product, id)
	return product.Value * quantity
}

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
