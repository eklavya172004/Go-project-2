package routes

import(
	"github.com/gorilla/mux"
	"github.com/eklavya/go-bookstore/pkg/controllers"
)

//this function will gonna have my all the routes
var RegisterBookstoreRoutes = func(router *mux.Router){
	//"When someone accesses this URL, call this function."
	router.HandleFunc("/book/",controllers.CreateBooks).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.GetBooksById).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.UpdateBooks).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controllers.DeleteBooks).Methods("DELETE")	
}