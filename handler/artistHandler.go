package handler

import (
	"groupie-tracker/models"
	"groupie-tracker/utilities"
	"html/template"
	"net/http"
)

type ArtistPageData struct {
	Artist         models.Artists
	MarqueeContent []string
}

func ArtistHandler(w http.ResponseWriter, r *http.Request, artists []models.Artists) {

	// Extract the artist ID from the URL
	id, err := utilities.ExtractArtistID(r.URL.Path)
	if err != nil {
		BadRequestHandler(w, r)
		return
	}

	if id > len(artists) || id <= 0 {
		BadRequestHandler(w, r)
		return
	}

	// Find the artist by the id ID from the URL
	artist, found := utilities.FindArtistByID(artists, id)
	if !found {
		ErrorHandler(w, r)
		return
	}

	// Convert the map data, in data of type []string
	relationsData := utilities.FormatRelations(artist.Relations)

	var MarqueeContent string
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			ErrorFiveHandler(w, r, err)
			return
		}
		MarqueeContent = r.FormValue("MarqueeContent")
	}

	// Prepare data for rendering the template
	var pageData ArtistPageData
	switch MarqueeContent {
	case "Relations":
		pageData = ArtistPageData{
			Artist:         artist,
			MarqueeContent: relationsData,
		}
	case "Locations":
		pageData = ArtistPageData{
			Artist:         artist,
			MarqueeContent: utilities.FormatData(artist.Locations),
		}
	case "ConcertDates":
		pageData = ArtistPageData{
			Artist:         artist,
			MarqueeContent: utilities.FormatDates(artist.ConcertDates),
		}
	default:
		pageData = ArtistPageData{
			Artist:         artist,
			MarqueeContent: relationsData,
		}
	}

	tempArtist, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		ErrorFiveHandler(w, r, err)
		return
	}
	if err := tempArtist.Execute(w, pageData); err != nil {

		ErrorFiveHandler(w, r, err)
	}
}
