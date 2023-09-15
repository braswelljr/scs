package server

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

// Server is the http server for the api.
type Server struct {
	Logger   *zerolog.Logger
	Listener net.Listener
	Http     http.Server
}

// Run starts the server that host webapp and api endpoints.
func (s *Server) Run(ctx context.Context) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	var group errgroup.Group

	group.Go(func() error {
		<-ctx.Done()
		return s.Http.Shutdown(context.Background())
	})

	group.Go(func() error {
		defer cancel()
		err := s.Http.Serve(s.Listener)
		if err == context.Canceled || errors.Is(err, http.ErrServerClosed) {
			err = nil
		}
		return err
	})

	return group.Wait()
}

// Close closes server and underlying listener.
func (s *Server) Close() error {
	return s.Http.Close()
}

// LoggingMiddlewareFunc is a middleware function to log incoming requests.
func LoggingMiddlewareFunc(logger *zerolog.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Create a response recorder to capture the status code.
			rr := NewResponseRecorder(w)

			// Call the next handler in the chain.
			next.ServeHTTP(rr, r)

			// Log information about the incoming request and response status.
			logger.Info().Str("method", r.Method).Str("path", r.URL.Path).Int("status", rr.Status).Msg("Request received")
		})
	}
}

// ResponseRecorder is a custom ResponseWriter that captures the status code.
type ResponseRecorder struct {
	http.ResponseWriter
	Status int
}

// NewResponseRecorder creates a new ResponseRecorder.
func NewResponseRecorder(w http.ResponseWriter) *ResponseRecorder {
	return &ResponseRecorder{
		ResponseWriter: w,
		Status:         http.StatusOK, // Default to 200 OK.
	}
}

// WriteHeader captures the status code.
func (r *ResponseRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}
