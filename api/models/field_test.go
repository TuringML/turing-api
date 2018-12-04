package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateField(t *testing.T) {
	f, err := CreateField(testDb, &Field{
		NodeID:       uint(1),
		Active:       false,
		KeyID:        "test",
		KeyName:      "Test property",
		KeyPrimary:   false,
		ValueClass:   DimensionClass,
		ValueType:    "string",
		ValueExample: "sloth",
	})
	if err != nil {
		t.Errorf("TestCreateField failed with error %v", err)
		return
	}
	assert.Equal(t, uint(1), f.NodeID)
	assert.Equal(t, false, f.Active)
	assert.Equal(t, "test", f.KeyID)
	assert.Equal(t, "Test property", f.KeyName)
	assert.Equal(t, false, f.KeyPrimary)
	assert.Equal(t, DimensionClass, f.ValueClass)
	assert.Equal(t, "string", f.ValueType)
	assert.Equal(t, "sloth", f.ValueExample)
}

func TestGetField(t *testing.T) {
	f, err := GetField(testDb, 1)
	if err != nil {
		t.Errorf("TestGetField failed with error %v", err)
	}
	assert.Equal(t, uint(1), f.NodeID)
	assert.Equal(t, true, f.Active)
	assert.Equal(t, "ip_address", f.KeyID)
	assert.Equal(t, "IP Address", f.KeyName)
	assert.Equal(t, false, f.KeyPrimary)
	assert.Equal(t, DimensionClass, f.ValueClass)
	assert.Equal(t, "string", f.ValueType)
	assert.Equal(t, "127.0.0.1", f.ValueExample)
}

func TestGetFieldFail(t *testing.T) {
	f, err := GetField(testDb, 100)
	if f != nil {
		t.Errorf("TestGetFieldFail failed. Expected nil, got %v", f)
	}
	assert.NotNil(t, err)
}

func TestGetFields(t *testing.T) {
	fs, err := GetFields(testDb, 1)
	if err != nil {
		t.Errorf("TestGetFields failed with error %v", err)
	}
	assert.Equal(t, 3, len(fs))
}

func TestUpdateField(t *testing.T) {
	f, err := GetField(testDb, 1)
	if err != nil {
		t.Errorf("TestUpdateField failed. Expected nil, got %v", err)
	}

	// update field
	f.Active = false

	err = UpdateField(testDb, f)
	if err != nil {
		t.Errorf("TestUpdateField failed. Expected nil, got %v", err)
	}
}

func TestDeleteField(t *testing.T) {
	err := DeleteField(testDb, 4)
	if err != nil {
		t.Errorf("TestDeleteField failed. Expected nil, got %v", err)
	}
}
