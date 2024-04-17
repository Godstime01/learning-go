package routes

import (
	"github.com/gorilla/mux"
	"github.com/godstime01/go-bookstore/controllers"
) 

var RegisterBookStoreRoutes = func (routes *mux.Router)  {
	router.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}/", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}/", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}/", controllers.DeleteBook).Methods("DELETE")
}