package utilities

import (
	"groupie-tracker/models"
	"strconv"
	"strings"
)

func searchSlice(text string, data []string) []string {
	var sliceOfInfo []string
	for _, info := range data {
		info = FormatLocation(info)
		if strings.Contains(strings.ToLower(info), text) {
			sliceOfInfo = append(sliceOfInfo, info)
		}
	}
	return sliceOfInfo
}

func Search(artists []models.Artists, text string) ([]models.Artists, []string, bool) {
	var sliceOfArtists []models.Artists
	var msg []string

	// Convert the text to lowercase for case-insensitive matching
	text = strings.ToLower(text)
	matchFound := false
	for _, artist := range artists {

		// Prepare the variables for matching
		members := searchSlice(text, artist.Members)
		nbr := strconv.Itoa(artist.CreationDate)
		locations := searchSlice(text, artist.Locations)

		switch {
		case strings.Contains(strings.ToLower(artist.Name), text): // Name match
			sliceOfArtists = append(sliceOfArtists, artist)
			msg = append(msg, artist.Name+" - Artist/Band name")
			matchFound = true

		case members != nil: // Member match
			for _, member := range members {
				msg = append(msg, member+" - Artist/Band member")
			}
			sliceOfArtists = append(sliceOfArtists, artist)
			matchFound = true

		case strings.Contains(artist.FirstAlbum, text):
			sliceOfArtists = append(sliceOfArtists, artist)
			msg = append(msg, artist.FirstAlbum+" - First album of "+artist.Name)
			matchFound = true

		case strings.Contains(nbr, text): // Creation Date match
			sliceOfArtists = append(sliceOfArtists, artist)
			msg = append(msg, nbr+" - Creation Date of "+artist.Name)
			matchFound = true

		case locations != nil: // Location match
			for _, location := range locations {
				msg = append(msg, FormatLocation(location)+" - Concert location of "+artist.Name)
			}
			sliceOfArtists = append(sliceOfArtists, artist)
			matchFound = true
		}
	}

	return sliceOfArtists, msg, matchFound
}
