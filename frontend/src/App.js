import React, { useState } from "react";
import SearchResults from "./components/SearchResults";
import { searchQuery } from "./api";

const App = () => {
    const [query, setQuery] = useState("");
    const [results, setResults] = useState([]);
    const [loading, setLoading] = useState(false);

    const handleSearch = async () => {
        setLoading(true);
        const data = await searchQuery(query);
        setResults(data.results);
        setLoading(false);
    };

    return (
        <div className="p-4 max-w-lg mx-auto">
            <h1 className="text-xl font-bold mb-4">Search Engine</h1>
            <input
                type="text"
                className="border p-2 w-full"
                placeholder="Enter search term..."
                value={query}
                onChange={(e) => setQuery(e.target.value)}
            />
            <button className="bg-blue-500 text-white p-2 mt-2 w-full" onClick={handleSearch}>
                Search
            </button>
            {loading ? <p>Loading...</p> : <SearchResults results={results} />}
        </div>
    );
};

export default App;
