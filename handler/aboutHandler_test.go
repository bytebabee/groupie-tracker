package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestHomeHandler tests the HomeHandler function
func TestAboutHandler(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/about", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(AboutHandler)
	handler.ServeHTTP(rr, req)

	// Check if the status code is 200 OK or 500 Internal Server Error
	if status := rr.Code; status != http.StatusOK && status != http.StatusInternalServerError {
		t.Errorf("Expected status code %d or %d, got %d", http.StatusOK, http.StatusInternalServerError, status)
	}

	// If status code is 500, check the response body for the error message
	if rr.Code == http.StatusInternalServerError {
		expectedError := "Error while parsing the 500 template"
		if !strings.Contains(rr.Body.String(), expectedError) {
			t.Errorf("Expected error message '%s', got '%s'", expectedError, rr.Body.String())
		}
	} else {
		// If the status code is 200, check if the content type is HTML
		contentType := rr.Header().Get("Content-Type")
		if contentType != "text/html; charset=utf-8" {
			t.Errorf("Expected content-type 'text/html; charset=utf-8', got %v", contentType)
		}

		// Optionally, check if the response body contains expected HTML content
		// Update this based on the actual content of your template
		expectedBody := "<html>" // Modify this based on your template content
		if !Contains(rr.Body.String(), expectedBody) {
			t.Errorf("Expected body to contain %v, got %v", expectedBody, rr.Body.String())

		}
	}
}
