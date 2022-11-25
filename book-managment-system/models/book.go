package models

import (
	"book-management/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	// Inserting the book into the database
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var foundBooks []Book

	// Getting all the books from the database
	db.Find(&foundBooks)

	return foundBooks
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var bookFound Book

	db := db.Where("ID=?", Id).Find(&bookFound)

	return &bookFound, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
