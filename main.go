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

// movie struct

type Movie struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Director *Director `json:"director"`
}

// Director structure

type Director struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// create a slice movies
var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get all vars in the r
	for _, item := range movies {
		movie_id, _ := strconv.Atoi(params["id"])
		if item.ID == movie_id {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = rand.Intn(10000)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get all vars in the r
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	for index, item := range movies {
		movie_id, _ := strconv.Atoi(params["id"])
		if item.ID == movie_id {
			// update the movie with id movie_id
			movie.ID = movie_id
			movies[index] = movie
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get all vars in the r
	for index, item := range movies {
		movie_id, _ := strconv.Atoi(params["id"])
		if item.ID == movie_id {
			// update the movie with id movie_id
			movies = append(movies[:index], movies[index+1:]...)
		}
	}
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: 1, Name: "Movie 1", Director: &Director{Id: 1, Name: "Jon"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie", createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("Starting web server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
