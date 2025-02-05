package handler

import (
	"encoding/json"
	"groupie-tracker/models"
	"groupie-tracker/utilities"
	"html/template"
	"net/http"
)

type DiscoverPageData struct {
	Artists []models.Artists
	Results []string
}

func DiscoverHandler(w http.ResponseWriter, r *http.Request, artists []models.Artists) {
	text := r.URL.Query().Get("query")

	// Perform search only if text is not empty
	searchedArtists, msg := utilities.Search(artists, text)

	// Detect AJAX requests (from JavaScript)
	if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg) // Return only suggestions
		return
	}

	// Render the discover.html template for normal requests
	tempDiscover, err := template.ParseFiles("templates/discover.html")
	if err != nil {
		http.Error(w, "Error while parsing the discover template", http.StatusInternalServerError)
		return
	}

	data := DiscoverPageData{
		Artists: searchedArtists,
		Results: msg,
	}

	w.WriteHeader(http.StatusOK)
	tempDiscover.Execute(w, data)
}
