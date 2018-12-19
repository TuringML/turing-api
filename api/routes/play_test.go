package routes

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlay(t *testing.T) {
	router.PUT("/playgrounds/:playground_id/play", Play)

	code, body, err := MockRequest(http.MethodPut, "/playgrounds/5/play", nil)
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
}
