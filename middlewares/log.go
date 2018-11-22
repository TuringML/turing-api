package middlewares

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msgf("%s %s\n", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
