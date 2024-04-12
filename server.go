package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var collection *mongo.Collection

type Movie struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title,omitempty" bson:"title,omitempty"`
	Director string             `json:"director,omitempty" bson:"director,omitempty"`
	Year     int                `json:"year,omitempty" bson:"year,omitempty"`
}


func main() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://SuchirMV:WGBOp8GJ0uFf2NdU@cluster0.rb7ivtw.mongodb.net/"))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection = client.Database("movies").Collection("movies1") // Specify the database name and collection name

	router := mux.NewRouter()

	router.HandleFunc("/movies", getAllMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var movies []Movie
	if err := cursor.All(ctx, &movies); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var movie Movie
	ctx := context.Background()
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	res, err := collection.InsertOne(ctx, movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	movie.ID = res.InsertedID.(primitive.ObjectID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	// Convert id to primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the movie ID
	movie.ID = objectID

	ctx := context.Background()
	_, err = collection.ReplaceOne(ctx, bson.M{"_id": objectID}, movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the updated movie
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}


func deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	_, err = collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

