package utilities

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	tool *http.Client
}

// Creates a new client with expire time 10s
func New() *Client {
	return &Client{
		tool: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) FetchFunc(url string, target interface{}) error {
	resp, err := c.tool.Get(url)
	if err != nil {
		return err

	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Print("Error getting data from API\n")
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func (c *Client) Fetcher(BaseUrl string) ([]models.Artists, error) {
	var artists []models.Artists

	if err := c.FetchFunc(BaseUrl+"/artists", &artists); err != nil {
		fmt.Print("Error fetching data from API\n")
		return nil, err
	}

	var group sync.WaitGroup
	chanErr := make(chan error, len(artists))

	for i := range artists {
		group.Add(1)
		go func(artist *models.Artists) {
			defer group.Done()

			var locations models.Locations
			if err := c.FetchFunc(artist.LocationsUrl, &locations); err != nil {
				chanErr <- fmt.Errorf("error fetching locations data from artist with ID %d: %v", artist.ID, err)
				return
			}
			artist.Locations = locations.Locations

			var dates models.Dates
			if err := c.FetchFunc(artist.ConcertDatesUrl, &dates); err != nil {
				chanErr <- fmt.Errorf("error fetching dates data from artist with ID %d: %v", artist.ID, err)
				return
			}
			artist.ConcertDates = dates.ConcertDates

			var relations models.Relations
			if err := c.FetchFunc(artist.RelationsUrl, &relations); err != nil {
				chanErr <- fmt.Errorf("error fetching relations data from artist with ID %d: %v", artist.ID, err)
				return
			}
			artist.Relations = relations.Relations

		}(&artists[i])

	}
	go func() {
		group.Wait()
		close(chanErr)
	}()
	for err := range chanErr {
		if err != nil {
			return nil, err
		}
	}
	return artists, nil

}
