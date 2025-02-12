package handler

import (
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
	text = utilities.ExtractGroupName(text)
	// text = utilities.InputFormat(text)

	// Perform search only if text is not empty
	searchedArtists, msg := utilities.Search(artists, text)

	tempDiscover, err := template.ParseFiles("templates/discover.html")
	if err != nil {
		http.Error(w, "Error while parsing the discover template", http.StatusInternalServerError)
		return
	}

	if searchedArtists == nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	data := DiscoverPageData{
		Artists: searchedArtists,
		Results: msg,
	}

	w.WriteHeader(http.StatusOK)
	tempDiscover.Execute(w, data)
}
