package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Define a simple route for testing
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Parquet Reader!"))
	}).Methods("GET")

	return r
}
