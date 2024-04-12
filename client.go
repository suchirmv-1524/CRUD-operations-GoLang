package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Movie struct {
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
		Title:    "The Godfather",
		Director: "Francis Ford Coppola",
		Year:     1972,
	}
	updateMovie(1, updatedMovie)

	// Get a specific movie
	getMovie(1)

	// Delete a movie
	deleteMovie(1)
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
		fmt.Printf("Title: %s, Director: %s, Year: %d\n", movie.Title, movie.Director, movie.Year)
	}
}

func updateMovie(id int, movie Movie) {
    jsonData, err := json.Marshal(movie)
    if err != nil {
        fmt.Println("Error marshalling JSON:", err)
        return
    }

    req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/movies/%d", id), bytes.NewBuffer(jsonData))
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

    if resp.StatusCode != http.StatusOK {
        fmt.Printf("Error updating movie: %s\n", resp.Status)
        return
    }

// Print response body for debugging
responseBody, err := ioutil.ReadAll(resp.Body)
if err != nil {
    fmt.Println("Error reading response body:", err)
    return
	fmt.Println("Response Body:", string(responseBody))
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

	fmt.Printf("Movie retrieved: Title: %s, Director: %s, Year: %d\n", movie.Title, movie.Director, movie.Year)
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
