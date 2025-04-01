export const searchQuery = async (query) => {
    try {
        const response = await fetch(`http://localhost:8080/search?query=${query}`);
        return response.json();
    } catch (error) {
        console.error("Search API Error:", error);
        return { results: [] };
    }
};
