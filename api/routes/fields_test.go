package routes

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFields(t *testing.T) {
	router.GET("/playgrounds/:playground_id/nodes/:node_id/fields", GetFields)

	code, body, err := MockRequest(http.MethodGet, "/playgrounds/1/nodes/1/fields", nil)
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

func TestGetField(t *testing.T) {
	router.GET("/playgrounds/:playground_id/nodes/:node_id/fields/:field_id", GetField)

	code, body, err := MockRequest(http.MethodGet, "/playgrounds/1/nodes/1/fields/1", nil)
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

func TestCreateField(t *testing.T) {
	router.POST("/playgrounds/:playground_id/nodes/:node_id/fields", CreateField)

	reqBody := `{
		"active": false,
		"key_id": "test",
		"key_name": "Test Property",
		"key_primary": true,
		"value_class": "Metric",
		"value_type": "double",
		"value_example": "12.00"
	}`

	code, body, err := MockRequest(http.MethodPost, "/playgrounds/1/nodes/1/fields", strings.NewReader(reqBody))
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

func TestUpdateField(t *testing.T) {
	router.PUT("/playgrounds/:playground_id/nodes/:node_id/fields/:field_id", UpdateField)

	reqBody := `{
		"active": true,
		"key_id": "test",
		"key_name": "Test Property",
		"key_primary": true,
		"value_class": "Metric",
		"value_type": "string",
		"value_example": "12.00"
	}`

	code, body, err := MockRequest(http.MethodPut, "/playgrounds/1/nodes/1/fields/1", strings.NewReader(reqBody))
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

func TestDeleteField(t *testing.T) {
	router.DELETE("/playgrounds/:playground_id/nodes/:node_id/fields/:field_id", DeleteField)

	code, body, err := MockRequest(http.MethodDelete, "/playgrounds/1/nodes/1/fields/1", nil)
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
