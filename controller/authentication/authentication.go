package authentication

import (
	"net/http"
)

// Login is the login request.
func Login(w http.ResponseWriter, r *http.Request) {
	// get the basic auth credentials
	_, _, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// check the credentials
}
