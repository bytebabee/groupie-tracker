document.getElementById("searchInput").addEventListener("input", function() {
    let query = this.value.replace(/^\s+/, ''); // Trim only leading spaces
    let suggestionsDiv = document.getElementById("suggestions");

    if (query.length < 1) {
        suggestionsDiv.style.display = "none";
        return;
    }

    fetch(`/discover?query=${encodeURIComponent(query)}`, { headers: { "X-Requested-With": "XMLHttpRequest" } })
        .then(response => response.json())
        .then(data => {
            suggestionsDiv.innerHTML = ""; // Clear previous suggestions

            if (data.length > 0) {
                data.forEach(item => {
                    let div = document.createElement("div");
                    div.textContent = item; // Display the suggestion

                    suggestionsDiv.appendChild(div);
                });
                suggestionsDiv.style.display = "block";
            } else {
                suggestionsDiv.style.display = "none";
            }
        });
});

document.addEventListener("click", function(event) {
    if (!event.target.closest("#searchInput") && !event.target.closest("#suggestions")) {
        document.getElementById("suggestions").style.display = "none";
    }
});
