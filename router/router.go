package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates a new router.
func NewRouter(router *mux.Router) {
	// set general headers
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// set headers
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("X-Content-Type-Options", "nosniff")

			// call the next handler in the chain
			next.ServeHTTP(w, r)
		})
	})

	// prefix all routes with /api
	api := router.PathPrefix("/api").Subrouter()

	// Add routes.
	api.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)

		// return a json response
		if err := json.NewEncoder(w).Encode(map[string]string{"message": "api health check"}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
		}
	}).Methods(http.MethodGet)
}
