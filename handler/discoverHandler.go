package handler

import (
	"groupie-tracker/models"
	"groupie-tracker/utilities"
	"html/template"
	"net/http"
	"strconv"
)

type DiscoverPageData struct {
	Artists []models.Artists
	Results []string
}

func DiscoverHandler(w http.ResponseWriter, r *http.Request, artists []models.Artists) {
	uneditedURLQuery := r.URL.Query().Get("query")
	text := r.URL.Query().Get("query")
	text = utilities.ExtractGroupName(text)
	// Perform search only if text is not empty
	searchedArtists, msg, id := utilities.Search(artists, text)

	if len(id) == 1 && utilities.ContainsAny(uneditedURLQuery, []string{
		"Band/Artist",
		"Member of",
		"First album of",
		"Concert location of",
		"Creation date of",
	}) {
		http.Redirect(w, r, "/artist/"+strconv.Itoa(id[0]), http.StatusSeeOther)
		return
	} else {
		// Check if search resulted in no artists, and set the proper response status
		if searchedArtists == nil {
			BadRequestHandler(w, r)
			return
		}

		// No need to call WriteHeader here if we're doing it in the if branch
		tempDiscover, err := template.ParseFiles("templates/discover.html")
		if err != nil {
			ErrorFiveHandler(w, r, err)
			return
		}

		// Proceed with writing the response with status OK
		data := DiscoverPageData{
			Artists: searchedArtists,
			Results: msg,
		}

		// This is the only place we set the status code to OK
		w.WriteHeader(http.StatusOK)
		tempDiscover.Execute(w, data)
	}
}
