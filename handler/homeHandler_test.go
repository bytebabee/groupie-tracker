package handler

import (
	"groupie-tracker/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestHomeHandler tests the HomeHandler function
func TestHomeHandler(t *testing.T) {
	// Mock artists data to pass to the handler
	mockArtists := []models.Artists{
		{
			ID:           1,
			Name:         "Test Artist",
			CreationDate: 2000,
			FirstAlbum:   "2005",
		},
	}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler with a wrapper function that includes the required artists argument
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		HomeHandler(w, r, mockArtists)
	})
	handler.ServeHTTP(rr, req)

	// Check if the status code is 200 OK, 400 Bad Request, or 500 Internal Server Error
	if status := rr.Code; status != http.StatusOK && status != http.StatusInternalServerError && status != http.StatusBadRequest {
		t.Errorf("Expected status code %d, %d, or %d, got %d", http.StatusOK, http.StatusInternalServerError, http.StatusBadRequest, status)
	}

	// Handle different response statuses
	body := rr.Body.String()

	if rr.Code == http.StatusInternalServerError {
		expectedError := "Error while parsing the 500 template"
		if !strings.Contains(body, expectedError) {
			t.Errorf("Expected error message '%s', got '%s'", expectedError, body)
		}
	} else if rr.Code == http.StatusBadRequest {
		expectedError := "Bad request"
		if !strings.Contains(body, expectedError) {
			t.Errorf("Expected error message '%s', got '%s'", expectedError, body)
		}
	} else {
		// If the status code is 200, check if the content type is HTML
		expectedContentType := "text/html; charset=utf-8"
		if rr.Header().Get("Content-Type") != expectedContentType {
			t.Errorf("Expected content-type '%s', got %v", expectedContentType, rr.Header().Get("Content-Type"))
		}

		// Ensure the response body contains expected HTML content
		expectedBody := "<html>" // Modify based on actual template content
		if !strings.Contains(body, expectedBody) {
			t.Errorf("Expected body to contain %v, got %v", expectedBody, body)
		}
	}
}

// Utility function to check if a string contains a substring
func Contains(str, substr string) bool {
	return strings.Contains(str, substr)
}
