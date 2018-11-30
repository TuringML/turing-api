package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	user, err := GetUser(testDb, "")
	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestGetUsers(t *testing.T) {
	users, err := GetUsers(testDb)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))
}
