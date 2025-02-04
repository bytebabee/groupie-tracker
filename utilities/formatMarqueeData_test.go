package utilities

import (
	"groupie-tracker/models"
	"sort"
	"testing"
)

func TestExtractArtistID(t *testing.T) {
	// Test cases
	tests := []struct {
		path     string
		expected int
		hasError bool
	}{
		{"/artist/123", 123, false},
		{"/artist/abc", 0, true},
		{"/artist/456", 456, false},
	}

	for _, test := range tests {
		t.Run(test.path, func(t *testing.T) {
			id, err := ExtractArtistID(test.path)
			if (err != nil) != test.hasError {
				t.Errorf("Expected error: %v, got: %v", test.hasError, err)
			}
			if id != test.expected {
				t.Errorf("Expected: %d, got: %d", test.expected, id)
			}
		})
	}
}

func TestFindArtistByID(t *testing.T) {
	// Test data
	artists := []models.Artists{
		{ID: 1, Name: "Artist 1"},
		{ID: 2, Name: "Artist 2"},
	}

	// Test cases
	tests := []struct {
		id       int
		expected models.Artists
		found    bool
	}{
		{1, models.Artists{ID: 1, Name: "Artist 1"}, true},
		{2, models.Artists{ID: 2, Name: "Artist 2"}, true},
		{3, models.Artists{}, false},
	}

	for _, test := range tests {
		t.Run("Find Artist", func(t *testing.T) {
			artist, found := FindArtistByID(artists, test.id)
			if found != test.found {
				t.Errorf("Expected found: %v, got: %v", test.found, found)
			}
			if artist.ID != test.expected.ID {
				t.Errorf("Expected: %v, got: %v", test.expected, artist)
			}
		})
	}
}

func TestFormatRelations(t *testing.T) {
	// Test data
	relations := map[string][]string{
		"New York":    {"2021-08-01", "2021-08-02"},
		"Los Angeles": {"2021-09-01"},
	}

	expected := []string{
		"New York 2021-08-01 • 2021-08-02",
		"Los Angeles 2021-09-01",
	}

	// Call FormatRelations function
	result := FormatRelations(relations)
	sort.Strings(result)
	sort.Strings(expected)
	// Check if the result matches the expected output
	for i, relation := range expected {
		if result[i] != relation {
			t.Errorf("Expected: %s, got: %s", relation, result[i])
		}
	}
}

func TestFormatData(t *testing.T) {
	// Test data
	tests := []struct {
		info     []string
		expected []string
	}{
		{
			[]string{"*Special Event", "location_1", "location-2"},
			[]string{"Special Event", "Location 1", "Location, 2"},
		},
		{
			[]string{"location-1", "*Important Event"},
			[]string{"Location, 1", "Important Event"},
		},
	}

	for _, test := range tests {
		t.Run("Format Data", func(t *testing.T) {
			result := FormatData(test.info)
			for i, formatted := range result {
				if formatted != test.expected[i] {
					t.Errorf("Expected: %s, got: %s", test.expected[i], formatted)
				}
			}
		})
	}
}

func TestFormatLocation(t *testing.T) {
	// Test cases
	tests := []struct {
		location string
		expected string
	}{
		{"new_york", "New York"},
		{"los_angeles", "Los Angeles"},
		{"san_francisco-usa", "San Francisco, USA"},
		{"chicago", "Chicago"},
	}

	for _, test := range tests {
		t.Run(test.location, func(t *testing.T) {
			result := FormatLocation(test.location)
			if result != test.expected {
				t.Errorf("Expected: %s, got: %s", test.expected, result)
			}
		})
	}
}
