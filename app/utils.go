package app

import (
	"net/http"

	"github.com/turing-ml/turing-api/pkg/version"
)

// Info returns the version of the api currently running
func Info(w http.ResponseWriter, r *http.Request) {
	Response(w, http.StatusOK, map[string]string{"version": version.LongVersion()})
}
