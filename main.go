package main

import (
	"./models"
	"./seed"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// "github.com/kr/pretty"
	"log"
	"net/http"
)

const (
	PORT = "8080"
)

var db *gorm.DB
var err error

func GetShops(w http.ResponseWriter, r *http.Request) {
	var shops []models.Shop
	db.Find(&shops)
	json.NewEncoder(w).Encode(&shops)
}
func GetShopById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var shop models.Shop
	var products []models.Product
	var orders []models.Order
	db.First(&shop, params["id"])
	db.Model(&shop).Related(&products)
	db.Model(&shop).Related(&orders)
	shop.Products = products
	shop.Orders = orders

	/*

		var shopRef models.Shop
		for _, product := range products {
			db.First(&shopRef, product.ShopID)
			product.Shop = shopRef
			log.Println(shopRef)
			log.Println(product.ID)
			log.Println(product.Shop)
		}
		log.Println("squee")
		for _, order := range orders {
			// db.First(&shopRef, order.ShopID)
			order.Shop = shop
			log.Println(shop)
			log.Println(order.Shop)
		}
		log.Println("ruh roh")
		for _, product := range shop.Products {
			log.Println(product.Shop)
		}
		log.Println("fadsddddroh")
		for _, V := range shop.Orders {
			log.Println(V.Shop)
		}

	*/

	json.NewEncoder(w).Encode(&shop)
}
func CreateShop(w http.ResponseWriter, r *http.Request) {
	var shop models.Shop
	json.NewDecoder(r.Body).Decode(&shop)
	db.Create(&shop)
	json.NewEncoder(w).Encode(&shop)
}
func DeleteShopById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var shop models.Shop
	db.First(&shop, params["id"])
	db.Delete(&shop)

	GetShops(w, r)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Waaaaaaaaaaaaaaaah itz a shop")
}

func ReseedHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Waaaaaaaaaaaaaaaah itz a shop")
	seed.DropAndReseedData(db)
}

func GetRootFields() graphql.Fields {
	return graphql.Fields{
		"shop": GetShopQuery(),
	}
}
func GetShopQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(ShopType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var shops []models.Shop

			// log.Printf("%+v\n\n", params.Source)
			// log.Printf("%# v", pretty.Formatter(params.Source)) //It will print all struct details
			// ... Implement the way you want to obtain your data here.
			db.Find(&shops) // MAYBE
			//

			return shops, nil
		},
	}
}

var ShopType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Shop",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.Int},
		"name": &graphql.Field{Type: graphql.String},
		"products": &graphql.Field{
			Type: graphql.NewList(ProductType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var products []models.Product

				// userID := params.Source.(User).ID
				// Implement logic to retrieve user associated roles from user id here.

				// var shop models.Shop
				// shopID := params.Source.(models.Shop).ID
				// log.Println(params.Source.(models.Shop))
				// db.First(&shop, shopID)
				// db.Model(&shop).Related(&products)
				//
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
	Name: "Product",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.Int},
		"shopId":   &graphql.Field{Type: graphql.Int},
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
	Name: "Order",
	Fields: graphql.Fields{
		"id":     &graphql.Field{Type: graphql.Int},
		"shopId": &graphql.Field{Type: graphql.Int},
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
	Name: "LineItem",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.Int},
		"productId": &graphql.Field{Type: graphql.Int},
		"orderId":   &graphql.Field{Type: graphql.Int},
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
	log.Print("Here we go!")
	db, err = gorm.Open("sqlite3", "foo.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// db.AutoMigrate(&Shop{}, &Product{}, &Order{}, &ListItem{})

	seed.DropAndReseedData(db)

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
	// router.HandleFunc("/shop", GetShops).Methods("GET")
	// router.HandleFunc("/shop/{id}", GetShopById).Methods("GET")
	// router.HandleFunc("/shop", CreateShop).Methods("POST")
	// router.HandleFunc("/shop/{id}", DeleteShopById).Methods("DELETE")
	router.Handle("/graphql", GraphqlHandler)

	log.Fatal(http.ListenAndServe(":"+PORT, router))

	// http.ListenAndServe(":"+PORT, router)
}
