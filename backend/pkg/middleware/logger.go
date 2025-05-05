package middleware

import (
	"fmt"
	"log"
	"net/http"
)

type writer struct {
	http.ResponseWriter
	status int
}

func (w *writer) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// | Method | Status | URI
const pattern string = "| %-7s | %d | %s "

// logger - log middleware
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writer := &writer{ResponseWriter: w}

		next.ServeHTTP(writer, r)

		artifact := fmt.Sprintf(pattern, r.Method, writer.status, r.URL)

		log.Println(artifact)
	})
}
