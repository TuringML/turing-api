package models

import (
	"os"
	"testing"

	"github.com/turing-ml/turing-api/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var testDb *mgo.Database
var user1 *User
var playground1 *Playground

func TestMain(m *testing.M) {
	tearUp()
	c := m.Run()
	tearDown()
	os.Exit(c)
}

func tearUp() {
	testDb = db.Open("192.168.99.100", "turing", "turing", "testing")

	// populate some user data
	testDb.C(collectionUsers).Insert(bson.M{"name": "test-user1"})
	testDb.C(collectionUsers).Insert(bson.M{"name": "test-user2"})
	testDb.C(collectionUsers).Insert(bson.M{"name": "test-user3"})

	// get user id for testing
	testDb.C(collectionUsers).Find(bson.M{"name": "test-user1"}).One(&user1)

	// populate some playground data
	testDb.C(collectionPlaygrounds).Insert(bson.M{"user_id": user1.ID.Hex(), "pipeline_schema": `{ "datasource": "s3"}`})
	testDb.C(collectionPlaygrounds).Insert(bson.M{"user_id": user1.ID.Hex(), "pipeline_schema": `{ "datasource": "s3", "storage":"druid"}`})
	testDb.C(collectionPlaygrounds).Insert(bson.M{"user_id": user1.ID.Hex(), "pipeline_schema": `{ "datasource": "s3", "storage":"druid"}`})

	// get playground id for testing
	var ps []Playground
	testDb.C(collectionPlaygrounds).Find(bson.M{"user_id": user1.ID.Hex()}).All(&ps)
	playground1 = &ps[0]
}

func tearDown() {
	testDb.C(collectionUsers).RemoveAll(nil)
	testDb.C(collectionPlaygrounds).RemoveAll(nil)
	testDb.Session.Close()
}
