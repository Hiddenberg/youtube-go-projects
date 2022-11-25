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

func CreateBook(w http.ResponseWriter, r *http.Request)

func DeleteBook(w http.ResponseWriter, r *http.Request)

func UpdateBook(w http.ResponseWriter, r *http.Request)
