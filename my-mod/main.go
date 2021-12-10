package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello, world")

	r := mux.NewRouter()

	r.HandleFunc("/", serverHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))

}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, world</h1>"))
}
