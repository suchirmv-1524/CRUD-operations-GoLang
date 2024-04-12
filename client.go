package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Year     int    `json:"year"`
}

func main() {
	var choice int
	for {
		fmt.Println("Menu:")
		fmt.Println("1. Create a new movie")
		fmt.Println("2. View all movies")
		fmt.Println("3. Update a movie")
		fmt.Println("4. Delete a movie")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			createMovie()
		case 2:
			getAllMovies()
		case 3:
			updateMovie()
		case 4:
			deleteMovie()
		case 5:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func createMovie() {
	var movie Movie
	fmt.Println("Enter details of the new movie:")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Title: ")
	title, _ := reader.ReadString('\n')
	movie.Title = strings.TrimSpace(title)

	fmt.Print("Director: ")
	director, _ := reader.ReadString('\n')
	movie.Director = strings.TrimSpace(director)

	fmt.Print("Year: ")
	fmt.Scanln(&movie.Year)

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
		fmt.Printf("ID: %s, Title: %s, Director: %s, Year: %d\n", movie.ID, movie.Title, movie.Director, movie.Year)
	}
}

func updateMovie() {
	var movie Movie
	fmt.Println("Enter ID of the movie to update: ")
	fmt.Scanln(&movie.ID)

	var updatedMovie Movie
	fmt.Println("Enter new details of the movie:")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Title: ")
	title, _ := reader.ReadString('\n')
	updatedMovie.Title = strings.TrimSpace(title)

	fmt.Print("Director: ")
	director, _ := reader.ReadString('\n')
	updatedMovie.Director = strings.TrimSpace(director)

	fmt.Print("Year: ")
	fmt.Scanln(&updatedMovie.Year)

	jsonData, err := json.Marshal(updatedMovie)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	url := fmt.Sprintf("http://localhost:8080/movies/%s", movie.ID)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set content type header
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error updating movie:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Movie updated successfully")
}


func deleteMovie() {
	var movie Movie
	fmt.Println("Enter ID of the movie to delete: ")
	fmt.Scanln(&movie.ID)

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/movies/%s", movie.ID), nil)
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
