package utilities

import (
	"groupie-tracker/models"
	"strconv"
	"strings"
)

func searchMembers(text string, data []string) []string {
	var sliceOfInfo []string
	for _, info := range data {

		if strings.Contains(strings.ToLower(info), text) {
			sliceOfInfo = append(sliceOfInfo, info)
		}
	}
	return sliceOfInfo
}

func searchLocations(text string, data []string) []string {
	var sliceOfInfo []string
	for _, info := range data {
		text = InputFormat(text)
		if strings.Contains(strings.ToLower(info), text) {
			sliceOfInfo = append(sliceOfInfo, info)
		}
	}
	return sliceOfInfo
}

func Search(artists []models.Artists, text string) ([]models.Artists, []string) {
	var sliceOfArtists []models.Artists
	var msg []string

	for _, artist := range artists {

		// Prepare the variables for matching
		members := searchMembers(text, artist.Members)
		nbr := strconv.Itoa(artist.CreationDate)
		locations := searchLocations(text, artist.Locations)

		switch {
		case strings.Contains(strings.ToLower(artist.Name)+" ", text): // Name match
			sliceOfArtists = append(sliceOfArtists, artist)
			msg = append(msg, artist.Name+" Band/Artist")

		case members != nil: // Member match
			for _, member := range members {
				member = FormatLocation(member)
				msg = append(msg, member+" - Member of "+artist.Name+" ")
			}
			sliceOfArtists = append(sliceOfArtists, artist)

		case strings.Contains(artist.FirstAlbum, text):
			sliceOfArtists = append(sliceOfArtists, artist)
			msg = append(msg, artist.FirstAlbum+" - First album of "+artist.Name+" ")

		case strings.Contains(nbr, text): // Creation Date match
			sliceOfArtists = append(sliceOfArtists, artist)
			msg = append(msg, nbr+" - Creation Date of "+artist.Name)

		case locations != nil: // Location match
			for _, location := range locations {
				msg = append(msg, FormatLocation(location)+" - Concert location of "+artist.Name+" ")
			}
			sliceOfArtists = append(sliceOfArtists, artist)
		}
	}

	return sliceOfArtists, msg
}
