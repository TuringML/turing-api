package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePlaground(t *testing.T) {
	uid := "auth0|789"
	n := "test playground"
	p, err := CreatePlayground(testDb, &Playground{
		Name:   n,
		UserID: uid,
	})
	if err != nil {
		t.Errorf("TestCreatePlaground failed with error %v", err)
	}
	assert.Equal(t, n, p.Name)
	assert.Equal(t, uid, p.UserID)
}

func TestGetPlayground(t *testing.T) {
	uid := "auth0|123"
	n := "playground 1"
	p, err := GetPlayground(testDb, 1)
	if err != nil {
		t.Errorf("TestGetPlayground failed with error %v", err)
	}
	assert.Equal(t, n, p.Name)
	assert.Equal(t, uint(1), p.ID)
	assert.Equal(t, uid, p.UserID)
}

func TestGetPlaygroundFail(t *testing.T) {
	p, err := GetPlayground(testDb, 100)
	if p != nil {
		t.Errorf("TestGetPlaygroundFail failed. Expected nil, got %v", p)
	}
	assert.NotNil(t, err)
}

func TestGetPlaygrounds(t *testing.T) {
	ps, err := GetPlaygrounds(testDb, "auth0|123")
	if err != nil {
		t.Errorf("TestGetPlaygrounds failed with error %v", err)
	}
	assert.Equal(t, 1, len(ps))
}

func TestGetAllPlaygrounds(t *testing.T) {
	ps, err := GetAllPlaygrounds(testDb)
	if err != nil {
		t.Errorf("TestGetAllPlaygrounds failed with error %v", err)
	}
	assert.Equal(t, 4, len(ps))
}

func TestUpdatePlayground(t *testing.T) {
	n := "this is a new name"

	p, err := GetPlayground(testDb, 1)
	if err != nil {
		t.Errorf("TestUpdatePlayground failed. Expected nil, got %v", err)
	}

	// update here
	p.Name = n

	err = UpdatePlayground(testDb, p)
	if err != nil {
		t.Errorf("TestUpdatePlayground failed. Expected nil, got %v", err)
	}
}

func TestDeletePlayground(t *testing.T) {
	err := DeletePlayground(testDb, 3)
	if err != nil {
		t.Errorf("TestDeletePlayground failed. Expected nil, got %v", err)
	}
}
