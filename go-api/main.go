package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//! Book structs
type Book struct {
	ID     string  `json:"id"`
	ISBN   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {

}
func createBook(w http.ResponseWriter, r *http.Request) {

}
func updateBook(w http.ResponseWriter, r *http.Request) {

}
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter()

	books = append(books, Book{ID: "1", ISBN: "134533", Title: "Test book", Author: &Author{Firstname: "John", Lastname: "Doe"}})

	books = append(books, Book{ID: "2", ISBN: "143223", Title: "Test book 2", Author: &Author{Firstname: "Arthur", Lastname: "Morgan"}})

	//! Route handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
