package utilities

import (
	"groupie-tracker/models"
	"strconv"
	"strings"
)

func searchSlice(text string, data []string) string {
	for _, info := range data {
		if strings.Contains(strings.ToLower(info), text) {
			return info
		}
	}
	return ""
}

func Search(artists []models.Artists, text string) ([]models.Artists, []string) {
	var sliceOfArtists []models.Artists
	var msg []string

	// Convert the text to lowercase for case-insensitive matching
	text = strings.ToLower(text)

	for _, artist := range artists {
		var matchFound bool // Flag to check if any match was found

		// Prepare the variables for matching
		member := searchSlice(text, artist.Members)
		nbr := strconv.Itoa(artist.CreationDate)
		location := searchSlice(text, artist.Locations)
		date := searchSlice(text, artist.ConcertDates)

		// Use switch to check the conditions
		switch {
		case strings.Contains(strings.ToLower(artist.Name), text): // Name match
			sliceOfArtists = append(sliceOfArtists, artist)
			msg = append(msg, artist.Name+" - Name")
			matchFound = true

		case member != "": // Member match
			sliceOfArtists = append(sliceOfArtists, artist)
			msg = append(msg, member+" - Member")
			matchFound = true

		case strings.Contains(nbr, text): // Creation Date match
			sliceOfArtists = append(sliceOfArtists, artist)
			msg = append(msg, nbr+" - Creation Date")
			matchFound = true

		case location != "": // Location match
			sliceOfArtists = append(sliceOfArtists, artist)
			msg = append(msg, location+" - Location")
			matchFound = true

		case date != "": // Concert Date match
			sliceOfArtists = append(sliceOfArtists, artist)
			msg = append(msg, date+" - Date")
			matchFound = true
		}

		// If any match was found, update sliceOfArtists and msg
		if matchFound {
			// You can add any additional processing here if needed for each match
		}
	}

	return sliceOfArtists, msg
}
