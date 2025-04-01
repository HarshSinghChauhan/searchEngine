# Parquet Search Engine

## Overview

This project implements a simple search engine that loads data from Parquet files and serves search results through an HTTP API. The backend is implemented in Go, and the frontend is built with React (if required). The main functionality is to load Parquet files into memory and perform full-text search queries across various fields such as `Message`, `Sender`, `Event`, and more.

## Features

- **Load Data from Parquet Files:** Load large datasets stored in Parquet files into memory for quick searching.
- **Search Engine:** Perform search queries on various fields such as `Message`, `Sender`, and `Event`.
- **HTTP API:** The search engine exposes an API that allows users to search for records by query.

## Installation

### Prerequisites

1. **Go (1.18 or later):** The backend is implemented using Go.
2. **Python (for generating Parquet files):** The project requires Parquet files. You can generate them using the provided Python script.
3. **Parquet File:** You must have a Parquet file (e.g., `sample.parquet`) with the required schema.

### Steps

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/yourusername/parquet-search-engine.git
   cd parquet-search-engine

2. **Install Go Dependencies:**

   ```bash
   go mod tidy
   
3. **Run the Backend:**

   ```bash
   go run main.go

4. **Test the API:**

    ```bash
   curl "http://localhost:8080/search?query=User"
   
5. **Example Response:**

Here’s an example of the JSON response when performing a search query:

```json
{
  "results": [
    {
      "Message": "User login successful",
      "Sender": "AuthService",
      "Event": "UserLogin",
      "NanoTimeStamp": 1710000001000
    }
  ],
  "match_count": 1,
  "search_time": "1ms"
}
```

## Project Structure

```plaintext
searchEngine/
│
├── backend/                # Golang backend
│   ├── parquet_samples/     # Sample Parquet files
│   ├── benchmark.go         # Benchmarking code
│   ├── go.mod               # Go module definition
│   ├── handlers.go          # HTTP request handlers
│   ├── main.go              # Entry point for the backend
│   ├── models.go            # Data models
│   ├── parquet_loader.go    # Code for loading and reading Parquet files
│   ├── routes.go            # API routes definition
│   ├── search.go            # Search functionality
│
├── frontend/               # React frontend
│   ├── src/                 # React source code
│   │   ├── components/      # Reusable components (e.g., SearchResults)
│   │   ├── api.js           # API request logic
│   │   ├── App.js           # Main React component
│   │   ├── index.js         # Entry point for React app
│   │   ├── styles.css       # Styles for the React app
│   ├── package.json         # React dependencies and scripts
│   ├── tailwind.config.js   # Tailwind CSS configuration
│   ├── vite.config.js       # Vite configuration for building React app
│
├── .gitignore              # Git ignore configuration
├── README.md               # Project documentation
```

### License
This project is open-source and available under the MIT License.
