// Get all the elements we need
const searchForm = document.getElementById('searchForm');
const searchInput = document.getElementById('searchInput');
const suggestions = document.querySelector('.suggestions');
const searchContainer = document.querySelector('.search-container');

// On page load, check if there's a search term in the URL
window.onload = () => {
    const searchTerm = new URLSearchParams(window.location.search).get('query');
    if (searchTerm) searchInput.value = searchTerm;
};

// Function to fetch and display suggestions
const fetchSuggestions = (searchTerm) => {
    if (searchTerm.length > 0) {
        fetch(`/discover?query=${encodeURIComponent(searchTerm)}`)
            .then(response => response.text())
            .then(html => {
                const newSuggestions = new DOMParser()
                    .parseFromString(html, 'text/html')
                    .querySelector('.suggestions');
                if (newSuggestions) {
                    suggestions.innerHTML = newSuggestions.innerHTML;
                    suggestions.style.display = 'block';
                }
            });
    } else {
        suggestions.style.display = 'none';
    }
};

// When user types in the search box
searchInput.addEventListener('input', () => fetchSuggestions(searchInput.value));

// When user clicks on search box, show suggestions if available
searchInput.addEventListener('focus', () => {
    if (searchInput.value.length > 0) suggestions.style.display = 'block';
});

// When user clicks outside search area, hide suggestions
document.addEventListener('click', (event) => {
    if (!searchContainer.contains(event.target)) suggestions.style.display = 'none';
});

// When user submits the search
searchForm.addEventListener('submit', (event) => {
    event.preventDefault();
    window.location.href = `/discover?query=${encodeURIComponent(searchInput.value)}`;
});
