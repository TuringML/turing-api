package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNode(t *testing.T) {
	n, err := CreateNode(testDb, &Node{
		Active:       true,
		Name:         "new name",
		XCoordinate:  0.0,
		YCoordinate:  10.5,
		PlaygroundID: 2,
		Type:         CollectorType,
	})
	if err != nil {
		t.Errorf("TestCreateNode failed with error %v", err)
	}
	assert.Equal(t, "new name", n.Name)
	assert.Equal(t, 0.0, n.XCoordinate)
	assert.Equal(t, 10.5, n.YCoordinate)
	assert.Equal(t, uint(2), n.PlaygroundID)
	assert.Equal(t, CollectorType, n.Type)
}

func TestGetNode(t *testing.T) {
	n, err := GetNode(testDb, 1)
	if err != nil {
		t.Errorf("TestGetNode failed with error %v", err)
	}
	assert.Equal(t, "s3 collector", n.Name)
	assert.Equal(t, 10.0, n.XCoordinate)
	assert.Equal(t, 15.0, n.YCoordinate)
	assert.Equal(t, uint(1), n.PlaygroundID)
	assert.Equal(t, CollectorType, n.Type)
}

func TestGetNodeFail(t *testing.T) {
	n, err := GetNode(testDb, 100)
	if n != nil {
		t.Errorf("TestGetNodeFail failed. Expected nil, got %v", n)
	}
	assert.NotNil(t, err)
}

func TestGetNodes(t *testing.T) {
	ns, err := GetNodes(testDb, 2)
	if err != nil {
		t.Errorf("TestGetNodes failed with error %v", err)
	}
	assert.Equal(t, 1, len(ns))
}

func TestUpdateNode(t *testing.T) {
	name := "this is a new name"

	n, err := GetNode(testDb, 1)
	if err != nil {
		t.Errorf("TestUpdateNode failed. Expected nil, got %v", err)
	}

	// update here
	n.Name = name

	err = UpdateNode(testDb, n)
	if err != nil {
		t.Errorf("TestUpdateNode failed. Expected nil, got %v", err)
	}
}

func TestDeleteNode(t *testing.T) {
	err := DeleteNode(testDb, 15)
	if err != nil {
		t.Errorf("TestDeleteNode failed. Expected nil, got %v", err)
	}
}
