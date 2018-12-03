package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLink(t *testing.T) {
	l, err := CreateLink(testDb, &Link{
		FromNodeID:  uint(1),
		FromFieldID: uint(1),
		ToNodeID:    uint(2),
		ToFieldID:   uint(4),
	})
	if err != nil {
		t.Errorf("TestCreateLink failed with error %v", err)
		return
	}
	assert.Equal(t, uint(4), l.ToFieldID)
	assert.Equal(t, uint(1), l.FromFieldID)
	assert.Equal(t, uint(2), l.ToNodeID)
	assert.Equal(t, uint(1), l.FromNodeID)
}

func TestGetLink(t *testing.T) {
	l, err := GetLink(testDb, 1)
	if err != nil {
		t.Errorf("TestGetLink failed with error %v", err)
	}
	assert.Equal(t, uint(4), l.ToFieldID)
	assert.Equal(t, uint(1), l.FromFieldID)
	assert.Equal(t, uint(2), l.ToNodeID)
	assert.Equal(t, uint(1), l.FromNodeID)
}

func TestGetLinkFail(t *testing.T) {
	l, err := GetLink(testDb, 100)
	if l != nil {
		t.Errorf("TestGetLinkFail failed. Expected nil, got %v", l)
	}
	assert.NotNil(t, err)
}

func TestGetLinks(t *testing.T) {
	ls, err := GetLinks(testDb, 1)
	if err != nil {
		t.Errorf("TestGetLinks failed with error %v", err)
	}
	assert.Equal(t, 1, len(ls))
}

func TestUpdateLink(t *testing.T) {
	l, err := GetLink(testDb, 1)
	if err != nil {
		t.Errorf("TestUpdateLink failed. Expected nil, got %v", err)
	}

	l.FromFieldID = uint(3)

	err = UpdateLink(testDb, l)
	if err != nil {
		t.Errorf("TestUpdateLink failed. Expected nil, got %v", err)
	}

	assert.Equal(t, uint(4), l.ToFieldID)
	assert.Equal(t, uint(3), l.FromFieldID)
	assert.Equal(t, uint(2), l.ToNodeID)
	assert.Equal(t, uint(1), l.FromNodeID)
}

func TestDeleteLink(t *testing.T) {
	err := DeleteLink(testDb, 1)
	if err != nil {
		t.Errorf("TestDeleteLink failed. Expected nil, got %v", err)
	}
}
