// Get all the elements we need
const searchForm = document.getElementById('searchForm');
const searchInput = document.getElementById('searchInput');
const suggestions = document.querySelector('.suggestions');
const searchContainer = document.querySelector('.search-container');
// When page loads, check if there's a search term in the URL
window.onload = function() {
    const searchTerm = new URLSearchParams(window.location.search).get('query');
    if (searchTerm) {
        searchInput.value = searchTerm;
    }
}
// When user types in the search box
searchInput.addEventListener('input', function() {
    const searchTerm = this.value;
    
    // If search box has text, show suggestions
    if (searchTerm.length > 0) {
        // Get suggestions from server
        fetch(`/discover?query=${searchTerm}`)
            .then(response => response.text())
            .then(html => {
                // Update suggestions list
                const newSuggestions = new DOMParser()
                    .parseFromString(html, 'text/html')
                    .querySelector('.suggestions');
                suggestions.innerHTML = newSuggestions.innerHTML;
                suggestions.style.display = 'block';
            });
    } else {
        // If search box is empty, hide suggestions
        suggestions.style.display = 'none';
    }
});
// When user clicks on search box
searchInput.addEventListener('click', function() {
    // If there's text, show suggestions
    if (this.value.length > 0) {
        suggestions.style.display = 'block';
    }
});
// When user clicks outside search area
document.addEventListener('click', function(event) {
    if (!searchContainer.contains(event.target)) {
        suggestions.style.display = 'none';
    }
});
// When user submits the search
searchForm.addEventListener('submit', function(event) {
    event.preventDefault(); // Don't submit form normally
    // Go to search results page
    window.location.href = `/discover?query=${searchInput.value}`;
});