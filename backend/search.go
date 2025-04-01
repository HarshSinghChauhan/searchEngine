package main

import (
	"strings"
)

// SearchData searches for records based on a query.
func SearchData(query string) []SearchResult {
	dataLock.RLock()
	defer dataLock.RUnlock()

	var results []SearchResult
	for _, entry := range data {
		if strings.Contains(strings.ToLower(entry.Message), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(entry.Sender), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(entry.Event), strings.ToLower(query)) {

			// Convert ParquetRecord to SearchResult
			results = append(results, SearchResult{
				Message:       entry.Message,
				Sender:        entry.Sender,
				Event:         entry.Event,
				NanoTimeStamp: entry.NanoTimeStamp,
			})
		}
	}
	return results
}
