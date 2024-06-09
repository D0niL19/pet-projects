package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	movie := Movie{}
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			movie := Movie{}
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, []Movie{Movie{
		ID:    "1",
		Isbn:  "53426",
		Title: "Moive 1",
		Director: &Director{
			Firstname: "D0niL",
			Lastname:  "Korotkov",
		}},
		Movie{
			ID:    "2",
			Isbn:  "6536654",
			Title: "Movie 2",
			Director: &Director{
				Firstname: "Denzel",
				Lastname:  "Pisya",
			},
		},
	}...)
	r.HandleFunc("/movies", getMovies).Methods(http.MethodGet)
	r.HandleFunc("/movies/{id}", getMovie).Methods(http.MethodGet)
	r.HandleFunc("/movies", createMovie).Methods(http.MethodPost)
	r.HandleFunc("/movies/{id}", updateMovie).Methods(http.MethodPut)
	r.HandleFunc("/movie/{id}", deleteMovie).Methods(http.MethodDelete)

	fmt.Printf("Start server at port :8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))

}
