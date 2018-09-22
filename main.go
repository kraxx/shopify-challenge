package main

import (
	"./models"
	"./seed"
	// "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
)

const (
	PORT    = "8080"
	DB_TYPE = "sqlite3"
	DB_PATH = "foo.db"
)

var db *gorm.DB
var err error

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Waaaaaaaaaaaaaaaah itz a shop")
}

func ReseedHandler(w http.ResponseWriter, r *http.Request) {
	seed.DropAndReseedData(db)
}

func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"shop":       GetShopQuery(),
		"createShop": GetCreateShopMutation(),
	}
}
func GetCreateShopMutation() *graphql.Field {
	return &graphql.Field{
		Type:        ShopType,
		Description: "Create new shop",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			shop := models.Shop{
				Name: params.Args["name"].(string),
			}
			db.Create(&shop)
			return shop, nil
		},
	}
}
func GetShopQuery() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(ShopType),
		Description: "Queries for Shop stuff",
		Args: graphql.FieldConfigArgument{
			"id":   &graphql.ArgumentConfig{Type: graphql.ID},
			"name": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var shops []models.Shop
			db.Where(params.Args).Find(&shops)
			return shops, nil
		},
	}
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
				var products []models.Product
				db.Model(params.Source.(models.Shop)).Related(&products)
				return products, nil
			},
		},
		"orders": &graphql.Field{
			Type: graphql.NewList(OrderType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var orders []models.Order
				db.Model(params.Source.(models.Shop)).Related(&orders)
				return orders, nil
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
				var lineItems []models.LineItem
				db.Model(params.Source.(models.Product)).Related(&lineItems)
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
				var lineItems []models.LineItem
				db.Model(params.Source.(models.Order)).Related(&lineItems)
				return lineItems, nil
			},
		},
		"value": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var lineItems []models.LineItem
				var sum int
				db.Model(params.Source.(models.Order)).Related(&lineItems)
				for _, item := range lineItems {
					var product models.Product
					db.First(&product, item.ProductID)
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

				var product models.Product
				db.First(&product, params.Source.(models.LineItem).ProductID)
				return product.Value * params.Source.(models.LineItem).Quantity, nil
			},
		},
	},
})

func main() {
	db, err = gorm.Open(DB_TYPE, DB_PATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&models.Shop{}, &models.Product{}, &models.Order{}, &models.LineItem{})

	// seed.DropAndReseedData(db)

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

	GraphqlHandler := handler.New(&handler.Config{
		Schema: &schema,
	})

	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler).Methods("GET")
	router.HandleFunc("/reseed", ReseedHandler).Methods("GET")
	router.Handle("/graphql", GraphqlHandler)

	log.Println("Starting up service...")
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
