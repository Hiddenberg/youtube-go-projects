package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	database, err := gorm.Open("mysql", "root:password@/libraryDB?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}

	db = database
}

// The purpose of this function is only to return the database externally
func GetDB() *gorm.DB {
	return db
}
