package middleware

import (
	"net/http"
	"slices"
)

type Stack func(http.Handler) http.Handler

func Chain(middlewares ...Stack) Stack {
	return func(next http.Handler) http.Handler {
		for _, middleware := range slices.Backward(middlewares) {
			next = middleware(next)
		}
		return next
	}
}
