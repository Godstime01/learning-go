package routes

import (
	"github.com/godstime01/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}/", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}/", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}/", controllers.DeleteBook).Methods("DELETE")
}
