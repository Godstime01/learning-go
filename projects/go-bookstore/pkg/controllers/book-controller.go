package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/godstime01/go-bookstore/pkg/models"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()

	res, _ := json.Marshal(books)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {

}

func CreateBook(w http.ResponseWriter, r *http.Request) {

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

}
