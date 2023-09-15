package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates a new router.
func NewRouter(router *mux.Router) {
	// prefix all routes with /api
	api := router.PathPrefix("/api").Subrouter()

	// Add routes.
	api.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		w.Header().Set("Content-Type", "application/json")

		// return a json response
		if err := json.NewEncoder(w).Encode(map[string]string{"message": "api health check"}); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
		}
	}).Methods(http.MethodGet)
}
