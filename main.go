package main

import (
	// "github.com/graphql-go/graphql"
	// "github.com/graphql-go/handler"
	"./models"
	"./seed"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

func main() {
	log.Print("Here we go!")
	db, err = gorm.Open("sqlite3", "foo.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// db.AutoMigrate(&Shop{}, &Product{}, &Order{}, &ListItem{})

	seed.DropAndReseedData(db)

	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler).Methods("GET")
	router.HandleFunc("/reseed", ReseedHandler).Methods("GET")
	router.HandleFunc("/shop", GetShops).Methods("GET")
	router.HandleFunc("/shop/{id}", GetShopById).Methods("GET")
	router.HandleFunc("/shop", CreateShop).Methods("POST")
	router.HandleFunc("/shop/{id}", DeleteShopById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+PORT, router))
	/*
		  schemaConfig := graphql.SchemaConfig{
		    Query: graphql.NewObject(graphql.ObjectConfig{
		      Name:   "RootQuery",
		      Fields: queries.GetRootFields(),
		    }),
		    Mutation: graphql.NewObject(graphql.ObjectConfig{
		      Name:   "RootMutation",
		      Fields: mutations.GetRootFields(),
		    }),
		  }

		  schema, err := graphql.NewSchema(schemaConfig)
		  if err != nil {
		    log.Fatalf("Error creating new schema, %v", err)
		  }

		  http.Handler := handler.New(&handler.Config{
		    Schema: &schema
		    })

			http.HandleFunc("/graphql", graphqlHandler(schema))
			log.Println("Listenin and servin baby")
			log.Fatal(http.ListenAndServe(":8080", nil))
	*/
}
