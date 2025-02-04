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
	var msg []string
	// Parse the template
	tempDiscover, err := template.ParseFiles("templates/discover.html")
	if err != nil {
		http.Error(w, "Error while parsing the discover template", http.StatusInternalServerError)
		return
	}

	//Parse the text value
	text := r.URL.Query().Get("query")

	if text != "" {
		//Search the model.Artist with the desired text
		searchedArtists, msg := utilities.Search(artists, text)

		data := DiscoverPageData{
			Artists: searchedArtists,
			Results: msg,
		}

		// Render the template
		w.WriteHeader(http.StatusOK)
		tempDiscover.Execute(w, data)

	} else {

		data := DiscoverPageData{
			Artists: artists,
			Results: msg,
		}

		// Render the template
		w.WriteHeader(http.StatusOK)
		tempDiscover.Execute(w, data)
	}
}
