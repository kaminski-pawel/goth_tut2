package main

import (
	"log"
	"net/http"

	"github.com/kaminski-pawel/goth_tut1/handlers"
)

func main() {
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("GET /", handlers.CountryList)
	mux.HandleFunc("GET /country/{name}", handlers.CountryDetail)
	mux.HandleFunc("GET /", handlers.SearchCountries)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
