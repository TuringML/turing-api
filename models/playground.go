package models

import (
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
}

// GetAllPlaygrounds will return all the playgrounds from the database
func GetAllPlaygrounds(db *mgo.Database) ([]Playground, error) {
	return nil, nil
}

// GetPlaygrounds returns all the playgrounds from the database of a given userID
func GetPlaygrounds(db *mgo.Database, userID string) ([]Playground, error) {
	return nil, nil
}

// GetPlayground return a single playground object given the ID in input
func GetPlayground(db *mgo.Database, ID string) (*User, error) {
	return nil, nil
}

// CreatePlayground creates a new playground in the database
func CreatePlayground(db *mgo.Database, userID, pipelineSchema string) error {
	return nil
}

// UpdatePlayground updates the playground's properties in the database
func UpdatePlayground(db *mgo.Database, ID, pipelineSchema string) error {
	return nil
}

// DeletePlayground deletes the specified playground from the database
func DeletePlayground(db *mgo.Database, ID string) error {
	return nil
}
