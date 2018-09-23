package main

import (
	// myDB "./db"
	"./models"
	"./seed"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

const (
	PORT = "8080"
)

var db *gorm.DB

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's a shopping API!")
}

func ReseedHandler(w http.ResponseWriter, r *http.Request) {
	seed.DropAndReseedData(db)
}

func main() {
	db = models.DB
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", IndexHandler).Methods("GET")
	router.HandleFunc("/reseed", ReseedHandler).Methods("GET")
	router.Handle("/graphql", models.GraphqlHandler)

	log.Println("Starting up service...")
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
