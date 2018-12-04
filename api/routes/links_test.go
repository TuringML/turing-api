package routes

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLinks(t *testing.T) {
	router.GET("/playgrounds/:playground_id/nodes/:node_id/links", GetLinks)

	code, body, err := MockRequest(http.MethodGet, "/playgrounds/1/nodes/1/links", nil)
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

func TestGetLink(t *testing.T) {
	router.GET("/playgrounds/:playground_id/nodes/:node_id/links/:link_id", GetLink)

	code, body, err := MockRequest(http.MethodGet, "/playgrounds/1/nodes/1/links/1", nil)
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

func TestCreateLink(t *testing.T) {
	router.POST("/playgrounds/:playground_id/nodes/:node_id/links", CreateLink)

	/*
	   FromNodeID:  uint(1),
	   FromFieldID: uint(1),
	   ToNodeID:    uint(2),
	   ToFieldID:   uint(4),
	*/

	reqBody := `{
		"from_node_id": 1,
		"from_field_id": 1,
		"to_node_id": 2,
		"to_field_id": 4
	}`

	code, body, err := MockRequest(http.MethodPost, "/playgrounds/1/nodes/1/links", strings.NewReader(reqBody))
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

func TestUpdateLink(t *testing.T) {
	router.PUT("/playgrounds/:playground_id/nodes/:node_id/links/:link_id", UpdateLink)

	reqBody := `{
		"from_node_id": 1,
		"from_field_id": 3,
		"to_node_id": 2,
		"to_field_id": 4
	}`

	code, body, err := MockRequest(http.MethodPut, "/playgrounds/1/nodes/1/links/1", strings.NewReader(reqBody))
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

func TestDeleteLink(t *testing.T) {
	router.DELETE("/playgrounds/:playground_id/nodes/:node_id/links/:link_id", DeleteLink)

	code, body, err := MockRequest(http.MethodDelete, "/playgrounds/1/nodes/1/links/1", nil)
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
