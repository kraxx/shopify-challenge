/*
** Entry point for the API
 */

package main

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/kraxx/shopify-challenge/models"
	"github.com/kraxx/shopify-challenge/seed"
	"log"
	"net/http"
	"os"
)

const (
	// 	PORT = os.Getenv("APP_PORT")
	PORT = "8080"
)

// var PORT string = os.Getenv("APP_PORT")

// Reference to our DB struct
var db *gorm.DB

// Landing page
var serveIndex = http.FileServer(http.Dir("./static"))

// Just hit this endpoint to reseed
func reseedHandler(w http.ResponseWriter, r *http.Request) {
	seed.DropAndReseedData(db)
	http.Redirect(w, r, "/", 301)
}

// Setup our routes, run the app.
func main() {
	db = models.DB
	defer db.Close()

	router := mux.NewRouter()
	router.Handle("/", serveIndex).Methods("GET")
	router.HandleFunc("/reseed", reseedHandler).Methods("GET")
	router.Handle("/graphql", models.GraphqlHandler)

	port := os.Getenv("PORT")
	if port == nil {
		port = PORT
	}

	log.Printf("Starting up shop API service on port:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
