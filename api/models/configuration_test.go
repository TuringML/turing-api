package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateConfiguration(t *testing.T) {
	b := []byte("test blob")
	c, err := CreateConfiguration(testDb, &Configuration{
		Blob: b,
	})
	if err != nil {
		t.Errorf("TestCreateConfiguration failed with error %v", err)
	}
	assert.Equal(t, b, c.Blob)

}

func TestGetConfiguration(t *testing.T) {
	c, err := GetConfiguration(testDb, 1)
	if err != nil {
		t.Errorf("TestGetConfiguration failed with error %v", err)
	}
	assert.Equal(t, "test blob", string(c.Blob))
}

func TestGetConfigurationFail(t *testing.T) {
	c, err := GetConfiguration(testDb, 100)
	if c != nil {
		t.Errorf("TestGetConfigurationFail failed. Expected nil, got %v", c)
	}
	assert.NotNil(t, err)
}

func TestUpdateConfiguration(t *testing.T) {
	b := "this is a new blob"

	c, err := GetConfiguration(testDb, 1)
	if err != nil {
		t.Errorf("TestUpdateConfiguration failed. Expected nil, got %v", err)
	}

	// update here
	c.Blob = []byte(b)

	err = UpdateConfiguration(testDb, c)
	if err != nil {
		t.Errorf("TestUpdateConfiguration failed. Expected nil, got %v", err)
	}
}

func TestDeleteConfiguration(t *testing.T) {
	err := DeleteConfiguration(testDb, 3)
	if err != nil {
		t.Errorf("TestDeleteConfiguration failed. Expected nil, got %v", err)
	}
}
