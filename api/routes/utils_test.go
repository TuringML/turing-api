package routes

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	router.GET("/version", Version)

	code, body, err := MockRequest(http.MethodGet, "/version", nil)
	if err != nil {
		t.FailNow()
	}

	// Decode body
	decoder := json.NewDecoder(body)
	var b string
	err = decoder.Decode(&b)
	if err != nil {
		t.FailNow()
	}

	assert.Equal(t, http.StatusOK, code)
	assert.Equal(t, `{"data":{"version":"TuringML APIs v0.0.1"}}`, b)
}
