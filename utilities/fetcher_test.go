package utilities

import (
	"encoding/json"
	"groupie-tracker/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Fetcher(t *testing.T) {
	var server *httptest.Server
	//Create a fake/mock API with httptest newServer and try to fetch and decode the data
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/artists":
			json.NewEncoder(w).Encode([]models.Artists{
				{
					ID:              1,
					Name:            "Artist One",
					LocationsUrl:    server.URL + "/locations/1",
					ConcertDatesUrl: server.URL + "/dates/1",
					RelationsUrl:    server.URL + "/relations/1",
				},
			})
		case "/locations/1":
			json.NewEncoder(w).Encode(models.Locations{
				Locations: []string{"Location1", "Location2"},
			})
		case "/dates/1":
			json.NewEncoder(w).Encode(models.Dates{
				ConcertDates: []string{"2023-01-01", "2023-01-02"},
			})
		case "/relations/1":
			json.NewEncoder(w).Encode(models.Relations{
				Relations: map[string][]string{
					"Location1": {"Relation1"},
					"Location2": {"Relation2"},
				},
			})
		default:
			http.Error(w, "not found", http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Initialize the client
	client := New()

	// Test Fetcher function
	artists, err := client.Fetcher(server.URL)
	if err != nil {
		t.Fatalf("Fetcher failed: %v", err)
	}

	// Validate the response
	if len(artists) != 1 {
		t.Fatalf("Expected 1 artist, got %d", len(artists))
	}

	artist := artists[0]

	// Check artist fields
	if artist.Name != "Artist One" {
		t.Errorf("Expected artist name 'Artist One', got '%s'", artist.Name)
	}

	// Check nested fields
	if len(artist.Locations) != 2 {
		t.Errorf("Expected 2 locations, got %d", len(artist.Locations))
	}

	if len(artist.ConcertDates) != 2 {
		t.Errorf("Expected 2 concert dates, got %d", len(artist.ConcertDates))
	}

	if len(artist.Relations) != 2 {
		t.Errorf("Expected 2 relations, got %d", len(artist.Relations))
	}
}
