package utilities

import (
	"groupie-tracker/models"
)

const (
	BaseUrl = "https://groupietrackers.herokuapp.com/api"
)

func FetchArtistsData() ([]models.Artists, error) {
	// Initialize the API client
	client := New()

	// Fetch artists data from the API
	artists, err := client.Fetcher(BaseUrl)
	if err != nil {
		return nil, err
	}

	return artists, nil
}
