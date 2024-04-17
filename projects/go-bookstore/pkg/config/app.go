package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "godstime01:godstime01/simplerest?charset=utf-8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
