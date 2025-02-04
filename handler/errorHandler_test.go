package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestErrorHandler tests the ErrorHandler function
func TestErrorHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/nonexistent", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ErrorHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound && status != http.StatusInternalServerError {
		t.Errorf("Expected status code %d or %d, got %d", http.StatusNotFound, http.StatusInternalServerError, status)
	}

	if rr.Code == http.StatusInternalServerError {
		expectedError := "Error while parsing the 500 template"
		if !strings.Contains(rr.Body.String(), expectedError) {
			t.Errorf("Expected error message '%s', got '%s'", expectedError, rr.Body.String())
		}
	} else {
		expectedContentType := "text/html; charset=utf-8"
		if rr.Header().Get("Content-Type") != expectedContentType {
			t.Errorf("Expected content-type '%s', got %v", expectedContentType, rr.Header().Get("Content-Type"))
		}

		expectedBody := "<html>" // Modify this based on your template content
		if !strings.Contains(rr.Body.String(), expectedBody) {
			t.Errorf("Expected body to contain %v, got %v", expectedBody, rr.Body.String())
		}
	}
}

// TestErrorFiveHandler tests the ErrorFiveHandler function
func TestErrorFiveHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/error", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ErrorFiveHandler(w, r, err)
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, status)
	}

	expectedError := "Error while parsing the 500 template"
	if !strings.Contains(rr.Body.String(), expectedError) {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, rr.Body.String())
	}
}

// TestBadRequestHandler tests the BadRequestHandler function
func TestBadRequestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/badrequest", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BadRequestHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest && status != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, status)
	}

}

// TestBadFetchHandler tests the BadFetchHandler function
func TestBadFetchHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/fetchissue", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BadFetchHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, status)
	}
}
