package api

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	"github.com/braswelljr/scs/router"
	"github.com/braswelljr/scs/server"
)

// NewServer creates a new server.
func NewServer(logger *zerolog.Logger, listener net.Listener) *server.Server {
	// Create router.
	r := mux.NewRouter()

	// Add middleware.
	r.Use(mux.CORSMethodMiddleware(r))

	// Add logging middleware.
	r.Use(mux.MiddlewareFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Create a response recorder to capture the status code.
			rr := NewResponseRecorder(w)

			// Call the next handler in the chain.
			next.ServeHTTP(rr, r)

			// Log information about the incoming request and response status.
			logger.Info().Str("method", r.Method).Str("path", r.URL.Path).Int("status", rr.Status).Msg("Request received")
		})
	}))

	// Create server.
	s := &server.Server{
		Logger:   logger,
		Listener: listener,
		Http: http.Server{
			Handler: r,
		},
	}

	router.NewRouter(r)

	return s
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
