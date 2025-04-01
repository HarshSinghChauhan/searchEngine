package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Load sample Parquet file
	err := LoadParquetFiles("parquet_samples/sample.parquet")
	if err != nil {
		log.Fatalf("Error loading Parquet file: %v", err)
	}

	r := SetupRoutes()
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
