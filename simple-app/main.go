package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Movie struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var movies []Movie

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, item := range movies {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//body := r.Body
	id := vars["id"]
	for index, item := range movies {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = id
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, item := range movies {
		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}

func main() {
	movies = append(movies, Movie{ID: "1", Title: "Elite", Director: &Director{ID: "1", Name: "Nguyen Van A"}})
	movies = append(movies, Movie{ID: "2", Title: "Money Heist", Director: &Director{ID: "2", Name: "Professor"}})
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)

	})
	r.HandleFunc("/movies", GetAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", GetMovieById).Methods("GET")
	r.HandleFunc("/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
