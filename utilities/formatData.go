package utilities

import (
	"groupie-tracker/models"
	"sort"
	"strconv"
	"strings"
)

// extractArtistID extracts and converts the artist ID from the URL path.
func ExtractArtistID(path string) (int, error) {
	idStr := strings.TrimPrefix(path, "/artist/")
	return strconv.Atoi(idStr)
}

// findArtistByID finds an artist by their ID in the given list of artists.
func FindArtistByID(artists []models.Artists, id int) (models.Artists, bool) {
	for _, artist := range artists {
		if artist.ID == id {
			return artist, true
		}
	}
	return models.Artists{}, false
}

// formatRelations formats the artist's relations data for display.
func FormatRelations(relations map[string][]string) []string {
	locationMap := make(map[string][]string)
	for location, dates := range relations {
		// Use the variadic operator to append all dates at once
		locationMap[location] = append(locationMap[location], dates...)
	}
	// Create formatted relationsData slice
	var relationsData []string
	for location, dates := range locationMap {
		// Format the location
		formattedLocation := FormatLocation(location)
		// Join the dates with a dot separator
		relation := formattedLocation + " " + strings.Join(dates, " • ")
		relationsData = append(relationsData, relation)
	}
	sort.Strings(relationsData)
	return relationsData
}

// formatData processes and formats a list of strings (locations or dates).
func FormatData(info []string) []string {
	var formatted []string
	for _, item := range info {
		if strings.HasPrefix(item, "*") {
			formatted = append(formatted, strings.TrimPrefix(item, "*"))
		} else {
			formatted = append(formatted, FormatLocation(item))
		}
	}
	return formatted
}

// Helper function to format the location name
func FormatLocation(location string) string {
	// Replace underscores and hyphens with spaces
	location = strings.ReplaceAll(location, "_", " ")
	location = strings.ReplaceAll(location, "-", ", ")
	// Capitalize the first letter of each word manually
	words := strings.Fields(location)
	for i, word := range words {
		if len(word) > 0 && (strings.Contains(word, "uk") || strings.Contains(word, "usa")) {
			words[i] = strings.ToUpper(string(word))
		} else if len(word) > 0 {
			// Convert the first character to uppercase and the rest to lowercase
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	// Join the words back together
	return strings.Join(words, " ")
}

// Format only dates
func FormatDates(dates []string) []string {
	var formatted []string
	for _, date := range dates {
		if strings.HasPrefix(date, "*") {
			formatted = append(formatted, strings.TrimPrefix(date, "*"))
		}
	}
	return formatted
}

func ExtractGroupName(text string) string {
	// Look for "of "/"Band/Artist" and check if there’s an ampersand to trim the rest
	if idx := strings.Index(text, "of "); idx != -1 {
		// Find the first occurrence of "&" after "of "
		return strings.ToLower(text[idx+3:])
	}
	if idx := strings.Index(text, "- Band/Artist"); idx != -1 {
		return strings.ToLower(text[:idx])
	}
	// Return the original text if no match
	return strings.ToLower(text)
}

// Format the string to match the json format
func InputFormat(text string) string {
	text = strings.ReplaceAll(text, ", ", "-")
	text = strings.ReplaceAll(text, ",", "-")
	text = strings.ReplaceAll(text, " ", "_")
	return text
}

func ContainsAny(query string, substrings []string) bool {
	for _, substring := range substrings {
		if strings.Contains(query, substring) {
			return true
		}
	}
	return false
}
