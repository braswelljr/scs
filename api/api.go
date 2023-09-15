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

	r.Use(server.LoggingMiddlewareFunc(logger))

	// Add logging middleware.
	r.Use(mux.MiddlewareFunc(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Debug().Str("method", r.Method).Str("path", r.URL.Path).Msg("request")
			next.ServeHTTP(w, r)
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
