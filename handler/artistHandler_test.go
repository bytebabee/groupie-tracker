package handler

import (
	"groupie-tracker/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestArtistHandler(t *testing.T) {
	// Mock artists data to pass to the handler
	mockArtists := []models.Artists{
		{
			ID:           1,
			Name:         "Test Artist",
			CreationDate: 2000,
			FirstAlbum:   "2005",
		},
	}

	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/artist/2", nil) // Using ID 2 to trigger BadRequestHandler
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Call the handler with the mock artists data
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ArtistHandler(w, r, mockArtists)
	})
	handler.ServeHTTP(w, req)

	// Check if the status code is valid
	if status := w.Code; status != http.StatusOK && status != http.StatusInternalServerError && status != http.StatusBadRequest {
		t.Errorf("Expected status code %d, %d, or %d, got %d", http.StatusOK, http.StatusInternalServerError, http.StatusBadRequest, status)
	}

	// Handling 500 Internal Server Error case
	if w.Code == http.StatusInternalServerError {
		expectedError := "Error while parsing the 500 template"
		if !strings.Contains(w.Body.String(), expectedError) {
			t.Errorf("Expected error message '%s', got '%s'", expectedError, w.Body.String())
		}
	} else if w.Code == http.StatusBadRequest {
		expectedError := "Invalid artist ID" // Make sure BadRequestHandler sends this message
		if !strings.Contains(w.Body.String(), expectedError) {
			t.Errorf("Expected error message '%s', got '%s'", expectedError, w.Body.String())
		}
	} else {
		// If the status code is 200, check if the content type is HTML
		contentType := w.Header().Get("Content-Type")
		if contentType != "text/html" {
			t.Errorf("Expected content-type 'text/html', got %v", contentType)
		}

		// Check if the mock artist name is in the response body
		expectedArtist := "Test Artist"
		if !strings.Contains(w.Body.String(), expectedArtist) {
			t.Errorf("Expected response body to contain artist name '%s', but it does not. Got: %s", expectedArtist, w.Body.String())
		}

		expectedAlbum := "2005"
		if !strings.Contains(w.Body.String(), expectedAlbum) {
			t.Errorf("Expected response body to contain first album '%s', but it does not. Got: %s", expectedAlbum, w.Body.String())
		}

		expectedBody := "<html>" // Modify this based on your template content
		if !strings.Contains(w.Body.String(), expectedBody) {
			t.Errorf("Expected body to contain %v, got %v", expectedBody, w.Body.String())
		}
	}
}
