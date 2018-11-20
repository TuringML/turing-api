package api

import (
	"github.com/turing-ml/turing-api/pkg/version"
	"net/http"
)

func Info(w http.ResponseWriter, r *http.Request) {
	Response(w, http.StatusOK, map[string]string{"version": version.LongVersion()})
}
