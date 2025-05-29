package main

import (
	"log"
	"net/http"

	"github.com/eklavya/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)

	    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"}, // React dev server
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"*"},
    })

	handler := c.Handler(r)

	// Start the HTTP server on port  and log any errors
	http.Handle("/",r)
log.Println("Server started at http://localhost:9010")
log.Fatal(http.ListenAndServe(":9010", handler))

}