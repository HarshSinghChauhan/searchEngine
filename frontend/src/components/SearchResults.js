import React from "react";

const SearchResults = ({ results }) => {
    return (
        <div className="mt-4">
            {results.length === 0 ? (
                <p>No results found.</p>
            ) : (
                <ul className="border p-2">
                    {results.map((res, index) => (
                        <li key={index} className="p-2 border-b">
                            <strong>{res.message}</strong> from {res.sender} (Event: {res.event})
                        </li>
                    ))}
                </ul>
            )}
        </div>
    );
};

export default SearchResults;
