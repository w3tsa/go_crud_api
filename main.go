package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

// Book Struct (Modal)
type Book struct {
	ID 			string `json:"id"`
	Isbn 		string `json:"isbn"`
	Title 		string `json:"title"`
	Author		Author `json:"author"`
}

// Author struct

type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Init books var as a slice Book struct
var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
params := mux.Vars(r) //Get params
// loop through books and find with id
for _,book := range books {
	if(book.ID == params["id"]) {
		json.NewEncoder(w).Encode(book)
		return
	}
}
json.NewEncoder(w).Encode(&Book{})
}
// Create a book
func createBook(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
var book Book
_ = json.NewDecoder(r.Body).Decode(&book)
book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID - not safe for production
books = append(books, book)
json.NewEncoder(w).Encode(book)
}
// Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
params := mux.Vars(r)
for index, book := range books {
	if(book.ID == params["id"]) {
	books = append(books[:index], books[index+1:]...)
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = params["id"]
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
	return
	}
}
json.NewEncoder(w).Encode(books)
}
// Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
params := mux.Vars(r)
for index, book := range books {
	if(book.ID == params["id"]) {
	books = append(books[:index], books[index+1:]...)
	break
	}
}
json.NewEncoder(w).Encode(books)
}


func main() {
	// init Router
	router := mux.NewRouter()

	// Mock Data
	books = append(books, Book{ID: "1", Isbn: "a55WWW454", Title: "Book one", Author: Author{Firstname: "Rahim", Lastname: "Beg"}})
	books = append(books, Book{ID: "2", Isbn: "a55WWW457", Title: "Book two", Author: Author{Firstname: "Karim", Lastname: "Beg"}})

	// Route handlers / endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}