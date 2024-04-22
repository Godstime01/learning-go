package models

import (
	"github.com/godstime01/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model

	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect() // to connect to the database

	db = config.GetDB()

	// auto migrate
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	// create new entry
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book

	db.Find(&Books)
	return Books
}

func GetBook(id int) (Book, *gorm.DB) {
	var Book Book

	db := db.Where("id=?", id).Find(&Book)
	return Book, db
}

func DeleteBook(id int) Book {
	var Book Book
	db.Where("id=?", id).Find(&Book)

	return Book
}
