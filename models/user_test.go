package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	user, err := CreateUser(*testDb, "test-user")
	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestGetUser(t *testing.T) {
	id := user1.ID.Hex()
	user, err := GetUser(*testDb, id)
	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestGetUserFail(t *testing.T) {
	id := "507f191e810c19729de860ea"
	user, err := GetUser(*testDb, id)
	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestGetUsers(t *testing.T) {
	users, err := GetUsers(*testDb)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))
}

func TestUpdateUser(t *testing.T) {
	id := user1.ID.Hex()
	err := UpdateUser(*testDb, id, "test-banana")
	assert.Nil(t, err)
}

func TestUpdateUserFail(t *testing.T) {
	id := "507f191e810c19729de860ea"
	err := UpdateUser(*testDb, id, "fail")
	assert.NotNil(t, err)
}

func TestDeleteUser(t *testing.T) {
	id := user1.ID.Hex()
	err := DeleteUser(*testDb, id)
	assert.Nil(t, err)
}

func TestDeleteUserFail(t *testing.T) {
	id := "507f191e810c19729de860ea"
	err := DeleteUser(*testDb, id)
	assert.NotNil(t, err)
}
