package models

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	collectionPlaygrounds = "playgrounds"
)

// Playground will contain the definition of the pipeline and the author
type Playground struct {
	ID             bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	UserID         string        `json:"userId" bson:"user_id"`
	PipelineSchema string        `json:"pipelineSchema" bson:"pipeline_schema"`
	CreatedAt      time.Time     `json:"createdAt" bson:"created_at"`
	UpdatedAt      time.Time     `json:"updatedAt" bson:"updated_at"`
}

// GetAllPlaygrounds will return all the playgrounds from the database
func GetAllPlaygrounds(db mgo.Database) ([]Playground, error) {
	var playgrounds []Playground
	err := db.C(collectionPlaygrounds).Find(nil).All(&playgrounds)
	if err != nil {
		return nil, err
	}
	return playgrounds, nil
}

// GetPlaygrounds returns all the playgrounds from the database of a given userID
func GetPlaygrounds(db mgo.Database, userID string) ([]Playground, error) {
	var playgrounds []Playground
	err := db.C(collectionPlaygrounds).Find(bson.M{"user_id": userID}).All(&playgrounds)
	if err != nil {
		return nil, err
	}
	return playgrounds, nil
}

// GetPlayground return a single playground object given the ID in input
func GetPlayground(db mgo.Database, userID, ID string) (*Playground, error) {
	var playground Playground
	err := db.C(collectionPlaygrounds).Find(bson.M{"_id": bson.ObjectIdHex(ID), "user_id": userID}).One(&playground)
	if err != nil {
		return nil, err
	}
	return &playground, nil
}

// CreatePlayground creates a new playground in the database
func CreatePlayground(db mgo.Database, userID, pipelineSchema string) (*Playground, error) {
	p := &Playground{
		UserID:         userID,
		PipelineSchema: pipelineSchema,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err := db.C(collectionPlaygrounds).Insert(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// UpdatePlayground updates the playground's properties in the database
func UpdatePlayground(db mgo.Database, userID, ID, pipelineSchema string) error {
	err := db.C(collectionPlaygrounds).Update(
		bson.M{"_id": bson.ObjectIdHex(ID), "user_id": userID},
		bson.M{"pipeline_schema": pipelineSchema, "updated_at": time.Now()})
	if err != nil {
		return err
	}
	return nil
}

// DeletePlayground deletes the specified playground from the database
func DeletePlayground(db mgo.Database, userID, ID string) error {
	err := db.C(collectionPlaygrounds).RemoveId(bson.ObjectIdHex(ID))
	if err != nil {
		return err
	}
	return nil
}
