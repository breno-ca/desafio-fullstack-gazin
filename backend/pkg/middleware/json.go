package middleware

import "net/http"

const (
	// Symbolic Constants
	header = "Content-Type"
	value  = "application/json"
)

func Json(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(header, value)

		next.ServeHTTP(w, r)
	})
}
