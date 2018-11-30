package models

import (
	"github.com/jinzhu/gorm"
)

// Node represents a generic node in the playground canvas
type Node struct {
	gorm.Model
	Playground      Playground    `gorm:"foreignkey:PlaygroundID;auto_preload"`
	PlaygroundID    uint          `json:"playground_id"`
	Active          bool          `json:"active"`
	Type            string        `json:"type"`
	Name            string        `json:"name"`
	X               float64       `json:"x"`
	Y               float64       `json:"y"`
	Configuration   Configuration `gorm:"foreignkey:ConfigurationID;auto_preload"`
	ConfigurationID uint          `json:"configuration_id"`
}

// GetNodes returns all the nodes from a playground
func GetNodes(db *gorm.DB, playgroundID int) ([]Node, error) {
	return nil, nil
}

// GetNode returns a single node given an ID
func GetNode(db *gorm.DB, playgroundID, ID int) (*Node, error) {
	return nil, nil
}

// CreateNode creates a new Node in the playground
func CreateNode(db *gorm.DB, playgroundID int, n Node) (*Node, error) {
	return nil, nil
}

// UpdateNode updates the information of the node specified by the ID
func UpdateNode(db *gorm.DB, playgroundID, ID int, n Node) error {
	return nil
}

// DeleteNode deletes the node specified by the ID
func DeleteNode(db *gorm.DB, playgroundID, ID int) error {
	return nil
}
