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
	Isbn     string    `json:"isdn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "1264", Title: "Percy Jackson", Director: &Director{Firstname: "John", Lastname: "Kennedy"}})
	movies = append(movies, Movie{ID: "2", Isbn: "1164", Title: "Harry Potter", Director: &Director{Firstname: "Joe", Lastname: "Biden"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovieById).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting Server at http://localhost:8000")

	isServerUp := http.ListenAndServe(":8000", r)
	if isServerUp != nil {
		log.Fatal(isServerUp)
	}

}

func getMovies(resp http.ResponseWriter, req *http.Request) {

	resp.Header().Set("Content-type", "application/json")
	json.NewEncoder(resp).Encode(movies)
}

func getMovieById(resp http.ResponseWriter, req *http.Request) {

	resp.Header().Set("Content-type", "application/json")

	params := mux.Vars(req)
		
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(resp).Encode(item)
			return
		}

		
	}

}

func createMovie(resp http.ResponseWriter, req *http.Request) {

	resp.Header().Set("Content-type", "application/json")

	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(100))
	movies = append(movies, movie)

	json.NewEncoder(resp).Encode(movie)
}

func updateMovie(resp http.ResponseWriter, req *http.Request) {

	resp.Header().Set("Content-type", "application/json")

	params := mux.Vars(req)

	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)

			var movie Movie
			_ = json.NewDecoder(req.Body).Decode(&movie)

			movie.ID = params["id"]
			movies = append(movies, movie)

			json.NewEncoder(resp).Encode(movie)
		}
	}
}

func deleteMovie(resp http.ResponseWriter, req *http.Request) {

	resp.Header().Set("Content-type", "application/json")

	params := mux.Vars(req)

	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}

	json.NewEncoder(resp).Encode(movies)

}
