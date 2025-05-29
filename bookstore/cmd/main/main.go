package main

import (
	"log"
	"net/http"
	"os"
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
	// http.Handle("/",r)

port := os.Getenv("PORT")
if port == "" {
    port = "9010" // fallback for local development
}

log.Printf("Server started at http://0.0.0.0:%s", port)
log.Fatal(http.ListenAndServe("0.0.0.0:"+port, handler))

}