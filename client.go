package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Movie struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Year     int    `json:"year"`
}

func main() {
	// Create a new movie
	newMovie := Movie{
		Title:    "The Godfather",
		Director: "Francis Ford Coppola",
		Year:     1972,
	}
	createMovie(newMovie)

	// Get all movies
	getAllMovies()

	// Update a movie
	updatedMovie := Movie{
		ID:       1,
		Title:    "The Godfather",
		Director: "Francis Ford Coppola",
		Year:     1972,
	}
	updateMovie(updatedMovie)

	// Get a specific movie
	getMovie(1)

	// Delete a movie
	//deleteMovie(1)
}

func createMovie(movie Movie) {
	jsonData, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/movies", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating movie:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Movie created successfully")
}

func getAllMovies() {
	resp, err := http.Get("http://localhost:8080/movies")
	if err != nil {
		fmt.Println("Error fetching movies:", err)
		return
	}
	defer resp.Body.Close()

	var movies []Movie
	err = json.NewDecoder(resp.Body).Decode(&movies)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	fmt.Println("All movies:")
	for _, movie := range movies {
		fmt.Printf("ID: %d, Title: %s, Director: %s, Year: %d\n", movie.ID, movie.Title, movie.Director, movie.Year)
	}
}

func updateMovie(movie Movie) {
	jsonData, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/movies/%d", movie.ID), bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error updating movie:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Movie updated successfully")
}

func getMovie(id int) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/movies/%d", id))
	if err != nil {
		fmt.Println("Error fetching movie:", err)
		return
	}
	defer resp.Body.Close()

	var movie Movie
	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	fmt.Printf("Movie retrieved: ID: %d, Title: %s, Director: %s, Year: %d\n", movie.ID, movie.Title, movie.Director, movie.Year)
}

func deleteMovie(id int) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/movies/%d", id), nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error deleting movie:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Movie deleted successfully")
}
