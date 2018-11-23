package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePlaground(t *testing.T) {
	pl, err := CreatePlayground(*testDb, user1.ID.Hex(), `{ "datasource": "s3"}`)
	assert.Nil(t, err)
	assert.NotNil(t, pl)
}

func TestGetPlayground(t *testing.T) {
	pl, err := GetPlayground(*testDb, user1.ID.Hex(), playground1.ID.Hex())
	assert.Nil(t, err)
	assert.NotNil(t, pl)
}

func TestGetPlaygroundFail(t *testing.T) {
	pl, err := GetPlayground(*testDb, user1.ID.Hex(), "507f191e810c19729de860ea")
	assert.NotNil(t, err)
	assert.Nil(t, pl)
}

func TestGetPlaygrounds(t *testing.T) {
	pls, err := GetPlaygrounds(*testDb, user1.ID.Hex())
	assert.Nil(t, err)
	assert.Equal(t, 4, len(pls))
}

func TestGetAllPlaygrounds(t *testing.T) {
	pls, err := GetAllPlaygrounds(*testDb)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(pls))
}

func TestUpdatePlayground(t *testing.T) {
	err := UpdatePlayground(*testDb, user1.ID.Hex(), playground1.ID.Hex(), `{ "datasource": "mysql"}`)
	assert.Nil(t, err)
}

func TestUpdatePlaygroundFail(t *testing.T) {
	err := UpdatePlayground(*testDb, user1.ID.Hex(), "507f191e810c19729de860ea", `{ "datasource": "mysql"}`)
	assert.NotNil(t, err)
}

func TestDeletePlaygroundFail(t *testing.T) {
	err := DeletePlayground(*testDb, user1.ID.Hex(), "507f191e810c19729de860ea")
	assert.NotNil(t, err)
}

func TestDeletePlayground(t *testing.T) {
	err := DeletePlayground(*testDb, user1.ID.Hex(), playground1.ID.Hex())
	assert.Nil(t, err)
}
