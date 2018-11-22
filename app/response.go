package app

import (
	"encoding/json"
	"net/http"
)

// ResponseError will return the error message
func ResponseError(w http.ResponseWriter, code int, message string) {
	Response(w, code, map[string]string{"error": message})
}

// Response will return a response with a specific object in it
func Response(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
