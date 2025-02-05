package utilities

import (
	"groupie-tracker/models"
	"strconv"
	"strings"
)

func searchSlice(text string, data []string) []string {
	var sliceOfInfo []string
	for _, info := range data {
		if strings.Contains(strings.ToLower(info), text) {
			sliceOfInfo = append(sliceOfInfo, info)
		}
	}
	return sliceOfInfo
}

func Search(artists []models.Artists, text string) ([]models.Artists, []string) {
	var sliceOfArtists []models.Artists
	var msg []string

	// Convert the text to lowercase for case-insensitive matching
	text = strings.ToLower(text)

	for _, artist := range artists {

		// Prepare the variables for matching
		members := searchSlice(text, artist.Members)
		nbr := strconv.Itoa(artist.CreationDate)
		locations := searchSlice(text, artist.Locations)

		switch {
		case strings.Contains(strings.ToLower(artist.Name), text): // Name match
			sliceOfArtists = append(sliceOfArtists, artist)
			msg = append(msg, artist.Name+"-Name")

		case members != nil: // Member match
			for _, member := range members {
				msg = append(msg, member+"-Artist/Band member")
			}
			sliceOfArtists = append(sliceOfArtists, artist)

		case strings.Contains(artist.FirstAlbum, text):
			sliceOfArtists = append(sliceOfArtists, artist)
			msg = append(msg, artist.FirstAlbum+"-First album of"+artist.Name)

		case strings.Contains(nbr, text): // Creation Date match
			sliceOfArtists = append(sliceOfArtists, artist)
			msg = append(msg, nbr+"-Creation Date of "+artist.Name)

		case locations != nil: // Location match
			sliceOfArtists = append(sliceOfArtists, artist)
			for _, location := range locations {
				msg = append(msg, FormatLocation(location)+"-Location concert of "+artist.Name)
			}
		}
	}

	return sliceOfArtists, msg
}
