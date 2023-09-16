package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewResponseRecorder(t *testing.T) {
	// Test with different HTTP statuses.
	statuses := []struct {
		status   string
		expected int
	}{
		{"StatusOK", http.StatusOK},
		{"StatusNotFound", http.StatusNotFound},
		{"StatusForbidden", http.StatusForbidden},
		{"StatusCreated", http.StatusCreated},
		{"StatusNoContent", http.StatusNoContent},
	}

	for _, status := range statuses {
		// Reset the mockResponseWriter for each iteration.
		mockResponseWriter := httptest.NewRecorder()

		// Create a new ResponseRecorder using NewResponseRecorder.
		recorder := NewResponseRecorder(mockResponseWriter)

		// Set the Status field to the expected status.
		recorder.Status = status.expected

		// Check if the Status field is set correctly.
		if recorder.Status != status.expected {
			t.Errorf("Expected Status to be %d, but got %d", status.expected, recorder.Status)
		} else if mockResponseWriter.Code != status.expected {
			t.Errorf("Expected Code to be %d, but got %d", status.expected, mockResponseWriter.Code)
		} else if recorder.ResponseWriter != mockResponseWriter {
			t.Errorf("Expected ResponseWriter to be %v, but got %v", mockResponseWriter, recorder.ResponseWriter)
		}
	}
}
