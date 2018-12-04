package routes

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPlaygrounds(t *testing.T) {
	router.GET("/playgrounds", GetPlaygrounds)

	code, body, err := MockRequest(http.MethodGet, "/playgrounds", nil)
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

func TestGetPlayground(t *testing.T) {
	router.GET("/playgrounds/:playground_id", GetPlayground)

	code, body, err := MockRequest(http.MethodGet, "/playgrounds/1", nil)
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

func TestCreatePlayground(t *testing.T) {
	router.POST("/playgrounds", CreatePlayground)

	reqBody := `{
		"name":"new playground from test"
	}`

	code, body, err := MockRequest(http.MethodPost, "/playgrounds", strings.NewReader(reqBody))
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

func TestUpdatePlayground(t *testing.T) {
	router.PUT("/playgrounds/:playground_id", UpdatePlayground)

	reqBody := `{
		"name":"new playground from update test"
	}`

	code, body, err := MockRequest(http.MethodPut, "/playgrounds/1", strings.NewReader(reqBody))
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

func TestDeletePlayground(t *testing.T) {
	router.DELETE("/playgrounds/:playground_id", DeletePlayground)

	code, body, err := MockRequest(http.MethodDelete, "/playgrounds/3", nil)
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
