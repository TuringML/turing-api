package routes

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNodes(t *testing.T) {
	router.GET("/playgrounds/:playground_id/nodes", GetNodes)

	code, body, err := MockRequest(http.MethodGet, "/playgrounds/1/nodes", nil)
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

func TestGetNode(t *testing.T) {
	router.GET("/playgrounds/:playground_id/nodes/:node_id", GetNode)

	code, body, err := MockRequest(http.MethodGet, "/playgrounds/1/nodes/1", nil)
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

func TestCreateNode(t *testing.T) {
	router.POST("/playgrounds/:playground_id/nodes", CreateNode)

	reqBody := `{
		"active":false,
		"type":"Enricher",
		"name":"test enricher",
		"x":30.87,
		"y":87.12
	}`

	code, body, err := MockRequest(http.MethodPost, "/playgrounds/1/nodes", strings.NewReader(reqBody))
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

func TestUpdateNode(t *testing.T) {
	router.PUT("/playgrounds/:playground_id/nodes/:node_id", UpdateNode)

	reqBody := `{
		"active": false,
		"type": "Enricher",
		"name": "test enricher 2",
		"x": 30.87,
		"y": 47.12
	}`

	code, body, err := MockRequest(http.MethodPut, "/playgrounds/1/nodes/1", strings.NewReader(reqBody))
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

func TestDeleteNode(t *testing.T) {
	router.DELETE("/playgrounds/:playground_id/nodes/:node_id", DeleteNode)

	code, body, err := MockRequest(http.MethodDelete, "/playgrounds/1/nodes/1", nil)
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
