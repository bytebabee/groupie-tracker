package utilities

import (
	"groupie-tracker/models"
	"strconv"
	"strings"
)

// Member slice search
func searchMembers(text string, data []string) []string {
	var sliceOfInfo []string
	for _, info := range data {

		if strings.Contains(strings.ToLower(info), text) {
			sliceOfInfo = append(sliceOfInfo, info)
		}
	}
	return sliceOfInfo
}

// Locations slice search
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

func Search(artists []models.Artists, text string) ([]models.Artists, []string, []int) {
	var sliceOfArtists []models.Artists
	var msg []string
	var id []int
	for _, artist := range artists {

		// Prepare the variables for matching and searching at the slices of data
		members := searchMembers(text, artist.Members)
		nbr := strconv.Itoa(artist.CreationDate)
		locations := searchLocations(text, artist.Locations)

		switch {
		case strings.Contains(strings.ToLower(artist.Name)+" ", text): // Name match
			sliceOfArtists = append(sliceOfArtists, artist)
			msg = append(msg, artist.Name+" - Band/Artist")
			id = append(id, artist.ID)

		case members != nil: // Member match
			for _, member := range members {
				member = FormatLocation(member)
				msg = append(msg, member+" - Member of "+artist.Name+" ")
			}
			sliceOfArtists = append(sliceOfArtists, artist)
			id = append(id, artist.ID)

		case strings.Contains(artist.FirstAlbum, text): // First Album match
			sliceOfArtists = append(sliceOfArtists, artist)
			msg = append(msg, artist.FirstAlbum+" - First album of "+artist.Name+" ")
			id = append(id, artist.ID)

		case strings.Contains(nbr, text): // Creation Date match
			sliceOfArtists = append(sliceOfArtists, artist)
			msg = append(msg, nbr+" - Creation date of "+artist.Name)
			id = append(id, artist.ID)

		case locations != nil: // Location match
			for _, location := range locations {
				msg = append(msg, FormatLocation(location)+" - Concert location of "+artist.Name+" ")
			}
			sliceOfArtists = append(sliceOfArtists, artist)
			id = append(id, artist.ID)
		}
	}

	return sliceOfArtists, msg, id
}
