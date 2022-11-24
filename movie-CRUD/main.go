package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func initializeMoviesList() {
	director1 := Director{
		FirstName: "John",
		LastName:  "Doe",
	}
	director2 := Director{
		FirstName: "Guillermo",
		LastName:  "Del Toro",
	}
	movies = append(movies,
		Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &director1},
		Movie{"2", "445679", "Movie Two", &director1}, // Adding movie by only filling the fields
		Movie{"3", "483394", "Movie Three", &director2},
	)

	fmt.Println("Some movies added to the list")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testEndpoint(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("trigering the root endpoint")
	fmt.Fprintf(writer, "Hello world")
}

func sendNotFoundJsonResponse(writer http.ResponseWriter, notFoundMsg string) {
	notFoundJson := map[string]string{
		"message": notFoundMsg,
	}

	writer.Header().Set("Content-Type", "application/json") // Setting the content type header
	writer.WriteHeader(http.StatusNotFound)                 // Setting the response code to 404
	json.NewEncoder(writer).Encode(notFoundJson)            // Sending the response
}

/* ********************** MOVIE FUNCTIONS ********************** */
func getMovies(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	fmt.Println("Getting all movies")
	err := json.NewEncoder(writer).Encode(movies)
	checkError(err)
}

func getMovie(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)

	idToSearch := mux.Vars(req)["id"]
	fmt.Printf("Getting movie with id: %v\n", idToSearch)

	for _, movie := range movies {
		if movie.ID == idToSearch {
			encoder.Encode(movie)
			return
		}
	}

	sendNotFoundJsonResponse(writer, "Movie not found")
}

func deleteMovie(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	idToDelete := params["id"]
	fmt.Printf("Deleting movie with id: %v\n", idToDelete)

	for index, movie := range movies {
		if movie.ID == idToDelete {
			movies = append(movies[:index], movies[index+1:]...) // Deleting the element using the index
			message := fmt.Sprintf(`{"message": "Movie %v deleted"}`, movie.ID)
			json.NewEncoder(writer).Encode(message)
			return
		}
	}

	sendNotFoundJsonResponse(writer, "The movie was not deleted, movie not found")
}

func createMovie(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("Creating new movie")
	writer.Header().Set("Content-Type", "application/json")

	// Creating a new movie object from the decoded json sent by the browser
	var newMovie Movie
	err := json.NewDecoder(req.Body).Decode(&newMovie)
	checkError(err)
	newMovie.ID = strconv.Itoa(rand.Intn(1000))

	// Appending the new movie to the existing list
	movies = append(movies, newMovie)

	// Notifying the user that the movie was correctly created
	json.NewEncoder(writer).Encode(newMovie)
}

func updateMovie(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)
	idToUpdate := params["id"]
	fmt.Printf("Updating movie with id: %v\n", idToUpdate)

	var updatedMovie Movie
	json.NewDecoder(req.Body).Decode(&updatedMovie)
	fmt.Println(updatedMovie)

	for index, movie := range movies {
		if movie.ID == idToUpdate {
			movies = append(movies[:index], movies[index+1:]...) // Deleting the element using the index

			movies = append(movies, updatedMovie) // Adding the updated movie data to the list

			json.NewEncoder(writer).Encode(`{"message": "movie updated"}`)
			return
		}
	}

	sendNotFoundJsonResponse(writer, "Movie not updated (id not found)")
}

func main() {
	router := mux.NewRouter()
	initializeMoviesList()

	// Creating the initial structure of the server functions

	router.HandleFunc("/", testEndpoint).Methods("GET")
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
