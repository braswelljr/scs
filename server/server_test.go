package server

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServerRunAndClose(t *testing.T) {
	// Create a test server with a mocked listener.
	mockListener, _ := net.Listen("tcp", "localhost:0")
	server := &Server{
		Listener: mockListener,
		Http: http.Server{
			// Configure your server settings here.
		},
	}

	// Create a context with a cancellation mechanism.
	ctx, cancel := context.WithCancel(context.Background())

	// Run the server in a goroutine.
	go func() {
		err := server.Run(ctx)
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			t.Errorf("Server.Run returned an unexpected error: %v", err)
		}
	}()

	// Give the server some time to start.
	time.Sleep(100 * time.Millisecond)

	// Send a test HTTP request to the server.
	testRequest := httptest.NewRequest("GET", "/", nil)
	testResponseRecorder := httptest.NewRecorder()
	go func() {
		server.Http.Handler.ServeHTTP(testResponseRecorder, testRequest)
	}()

	// Wait for the response to be recorded.
	time.Sleep(100 * time.Millisecond)

	// Perform assertions on the response.
	if testResponseRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, testResponseRecorder.Code)
	}

	expectedResponseBody := "Hello, World!" // Replace with your expected response body.
	if body := testResponseRecorder.Body.String(); body != expectedResponseBody {
		t.Errorf("Expected response body '%s', but got '%s'", expectedResponseBody, body)
	}

	expectedHeader := "application/json" // Replace with your expected header value.
	if contentType := testResponseRecorder.Header().Get("Content-Type"); contentType != expectedHeader {
		t.Errorf("Expected Content-Type header '%s', but got '%s'", expectedHeader, contentType)
	}

	// Close the server.
	if err := server.Close(); err != nil {
		t.Errorf("Server.Close returned an unexpected error: %v", err)
	}

	// Cancel the context to stop the server gracefully.
	cancel()

	// Wait for the server to shut down.
	select {
	case <-time.After(2 * time.Second): // Adjust the timeout as needed.
		t.Errorf("Server did not shut down gracefully within the timeout.")
	}
}
