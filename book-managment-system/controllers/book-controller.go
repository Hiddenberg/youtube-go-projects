package controllers

import (
	"book-management/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	NewBooks := models.GetAllBooks()

	// Parsing the database response into a json format
	res, _ := json.Marshal(NewBooks)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]

	intId, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing id")
	}

	bookDetails, _ := models.GetBookById(intId)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookDetails) // Sending the json response directly using the encoder
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book

	json.NewDecoder(r.Body).Decode(&newBook)

	bookCreated := newBook.CreateBook()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookCreated)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["bookId"], 10, 0)
	if err != nil {
		fmt.Println(err)
	}
	deletedBook := models.DeleteBook(bookId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(deletedBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var bookUpdateData models.Book
	json.NewDecoder(r.Body).Decode(&bookUpdateData)

	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["bookId"], 10, 0)
	if err != nil {
		fmt.Println(err)
	}

	bookDetails, db := models.GetBookById(bookId)
	if bookUpdateData.Name != "" {
		bookDetails.Name = bookUpdateData.Name
	}
	if bookUpdateData.Author != "" {
		bookDetails.Author = bookUpdateData.Author
	}
	if bookUpdateData.Publication != "" {
		bookDetails.Publication = bookUpdateData.Publication
	}
	db.Save(&bookDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookDetails)
}
