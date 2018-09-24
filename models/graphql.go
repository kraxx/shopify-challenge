/*
** All our logic for graphQL integration.
 */

package models

import (
	"../db"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

var DB *gorm.DB

var GraphqlHandler http.Handler

/*
** Queries and Mutations available to our graphQL endpoint
** Getters are called simply with the name of the schema.
** Create, Update, and Delete are more verbose
 */
var GraphqlRootFields = graphql.Fields{
	"shop":           GetShopQuery(),
	"createShop":     CreateShopMutation(),
	"updateShop":     UpdateShopMutation(),
	"deleteShop":     DeleteShopMutation(),
	"product":        GetProductQuery(),
	"createProduct":  CreateProductMutation(),
	"updateProduct":  UpdateProductMutation(),
	"deleteProduct":  DeleteProductMutation(),
	"order":          GetOrderQuery(),
	"createOrder":    CreateOrderMutation(),
	"updateOrder":    UpdateOrderMutation(),
	"deleteOrder":    DeleteOrderMutation(),
	"lineItem":       GetLineItemQuery(),
	"createLineItem": CreateLineItemMutation(),
	"updateLineItem": UpdateLineItemMutation(),
	"deleteLineItem": DeleteLineItemMutation(),
}

/*
** Init is run once imported.
** Set the database models to our database struct.
** Set the graphQL schemas, and return the handler for use with our graphQL endpoint.
 */
func init() {
	DB = db.DB
	DB.AutoMigrate(&Shop{}, &Product{}, &Order{}, &LineItem{})

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: GraphqlRootFields,
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: GraphqlRootFields,
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
