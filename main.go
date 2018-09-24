/*
** Entry point for the API
 */

package main

import (
	"./models"
	"./seed"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	// "os"
)

const (
	// PORT = os.Getenv["port"] || "8080"
	PORT = "8080"
)

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

	log.Println("Starting up shop API service...")
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
