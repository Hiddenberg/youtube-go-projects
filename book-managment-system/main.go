package main

import (
	"book-management/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
