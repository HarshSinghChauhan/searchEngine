package main

// SearchResult represents a search result structure.
type SearchResult struct {
	Message       string `json:"message"`
	Sender        string `json:"sender"`
	Event         string `json:"event"`
	NanoTimeStamp int64  `json:"nanoTimeStamp"`
}

// SearchResponse represents the API response format.
type SearchResponse struct {
	Results    []SearchResult `json:"results"`
	MatchCount int            `json:"match_count"`
	SearchTime string         `json:"search_time"`
}
