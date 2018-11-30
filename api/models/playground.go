package models

import (
	"github.com/jinzhu/gorm"
)

// Playground will contain the definition of the pipeline and the author
type Playground struct {
	gorm.Model
	User   User   `gorm:"foreignkey:UserID;auto_preload"`
	UserID string `json:"user_id"`
	Name   string `json:"name"`
}

// GetAllPlaygrounds will return all the playgrounds from the database
func GetAllPlaygrounds(db *gorm.DB) ([]Playground, error) {
	var playgrounds []Playground

	return playgrounds, nil
}

// GetPlaygrounds returns all the playgrounds from the database of a given userID
func GetPlaygrounds(db *gorm.DB, userID string) ([]Playground, error) {
	var playgrounds []Playground

	return playgrounds, nil
}

// GetPlayground return a single playground object given the ID in input
func GetPlayground(db *gorm.DB, userID, ID string) (*Playground, error) {
	var playground Playground

	return &playground, nil
}

// CreatePlayground creates a new playground in the database
func CreatePlayground(db *gorm.DB, userID, pipelineSchema string) (*Playground, error) {
	p := &Playground{
		UserID: userID,
	}

	return p, nil
}

// UpdatePlayground updates the playground's properties in the database
func UpdatePlayground(db *gorm.DB, userID, ID, pipelineSchema string) error {

	return nil
}

// DeletePlayground deletes the specified playground from the database
func DeletePlayground(db *gorm.DB, userID, ID string) error {

	return nil
}
