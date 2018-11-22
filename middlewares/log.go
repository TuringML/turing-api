package middlewares

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

// Logging is a middleware that will print out the incoming request
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msgf("%s %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
