CRUD Operations Using GoLang
This project implements a RESTful API in GoLang for performing CRUD (Create, Read, Update, Delete) operations on a database. It provides endpoints to interact with the database via HTTP requests.

Features
Create: Add new entries to the database.
Read: Retrieve existing entries from the database.
Update: Modify existing entries in the database.
Delete: Remove entries from the database.
Technologies Used
Go (Golang) - Programming language
MongoDB - Database system
Gorilla Mux - HTTP router and dispatcher for Go
Installation
Set up your MongoDB connection string in the server.go file.

Run the server:

bash
Copy code
go run server.go
Access the API at http://localhost:8080.

Usage
Use HTTP methods (GET, POST, PUT, DELETE) to interact with the API endpoints.
Example usage:
To create a new entry: POST /movies
To retrieve all entries: GET /movies
To retrieve a specific entry: GET /movies/{id}
To update an entry: PUT /movies/{id}
To delete an entry: DELETE /movies/{id}

