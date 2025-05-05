package middleware

import (
	"net/http"

	"backend/config"
)

type CORS struct {
	AllowOrigin string
	Origin      string

	AllowMethods string
	Methods      string

	AllowHeaders string
	Headers      string

	AllowMaxAge string
	MaxAge      string
}

func NewCORS(cfg config.Config) CORS {
	return CORS{
		AllowOrigin: "Access-Control-Allow-Origin",
		Origin:      cfg.CorsAllowOrigin,

		AllowMethods: "Access-Control-Allow-Methods",
		Methods:      "GET, POST, PUT, DELETE, OPTIONS",

		AllowHeaders: "Access-Control-Allow-Headers",
		Headers:      "Accept, Authorization, Content-Type",

		AllowMaxAge: "Access-Control-Max-Age",
		MaxAge:      "18000",
	}
}

func (c *CORS) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(c.AllowOrigin, c.Origin)
		w.Header().Set(c.AllowMethods, c.Methods)
		w.Header().Set(c.AllowHeaders, c.Headers)

		if r.Method == "OPTIONS" {
			w.Header().Set(c.AllowMaxAge, c.MaxAge)
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
