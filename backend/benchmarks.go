package main

import (
	"log"
	"time"
)

// BenchmarkSearch tests the search execution time.
func BenchmarkSearch(query string) {
	start := time.Now()
	_ = SearchData(query)
	elapsed := time.Since(start)
	log.Printf("Search query '%s' took %s\n", query, elapsed)
}
