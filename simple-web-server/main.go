package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(writer http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(writer, "404 not found", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(writer, "This method is not supported", http.StatusNotAcceptable)
	}

	fmt.Fprintln(writer, "Hello!")
}

func formHandler(writer http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Fprintf(writer, "Parse Form error: %v", err)
	}

	fmt.Fprintln(writer, "POST request successful")

	name := req.FormValue("name")
	email := req.FormValue("email")
	phone := req.FormValue("phone")

	fmt.Println("The user was correctly register with the following data")
	fmt.Println(name, email, phone)
}

func main() {
	// creating a file server in order to use it in one of the routes
	fileServer := http.FileServer(http.Dir("./html"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form-handler", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
