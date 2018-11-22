package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionUsers = "users"
)

// User is the struct used to deal with the users in the database
type User struct {
	ID   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

// GetUsers returns all the users from the database
func GetUsers(db mgo.Database) ([]User, error) {
	var users []User
	err := db.C(collectionUsers).Find(nil).All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// GetUser return a single user object given the ID in input
func GetUser(db mgo.Database, ID string) (*User, error) {
	var user User
	err := db.C(collectionUsers).FindId(bson.ObjectIdHex(ID)).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user in the database
func CreateUser(db mgo.Database, name string) error {
	u := &User{Name: name}
	err := db.C(collectionUsers).Insert(u)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser updates the user's properties in the database
func UpdateUser(db mgo.Database, ID, name string) error {
	err := db.C(collectionUsers).UpdateId(bson.ObjectId(ID), name)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes the specified user from the database
func DeleteUser(db mgo.Database, ID string) error {
	err := db.C(collectionUsers).RemoveId(bson.ObjectId(ID))
	if err != nil {
		return err
	}
	return nil
}
