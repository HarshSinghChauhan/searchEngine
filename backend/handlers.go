package main

import (
	"encoding/json"
	"net/http"
)

// SearchHandler handles search API requests.
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Missing search query", http.StatusBadRequest)
		return
	}

	results := SearchData(query)

	response := SearchResponse{
		Results:    results,
		MatchCount: len(results),
		SearchTime: "1ms",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
