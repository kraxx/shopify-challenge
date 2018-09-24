package main

import (
	"./models"
	"./seed"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

const (
	PORT = "8080"
)

var db *gorm.DB

func ReseedHandler(w http.ResponseWriter, r *http.Request) {
	seed.DropAndReseedData(db)
}

func main() {
	db = models.DB
	defer db.Close()

	router := mux.NewRouter()
	router.Handle("/", http.FileServer(http.Dir("./static"))).Methods("GET")
	router.HandleFunc("/reseed", ReseedHandler).Methods("GET")
	router.Handle("/graphql", models.GraphqlHandler)

	log.Println("Starting up service...")
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
