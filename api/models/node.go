package models

import (
	"database/sql/driver"
	"errors"

	"github.com/jinzhu/gorm"
)

// Type specifies the Type type of the Field. It transforms it in enum in SQL
type Type string

const (
	// CollectorType specify the collector enum
	CollectorType Type = "COLLECTOR"
	// EnricherType specify the enricher enum
	EnricherType Type = "ENRICHER"
	// JoinerType specify the joiner enum
	JoinerType Type = "JOINER"
	// OperatorType specify the operator enum
	OperatorType Type = "OPERATOR"
	// IntellectorType specify the intellector enum
	IntellectorType Type = "INTELLECTOR"
	// StorerType specify the storer enum
	StorerType Type = "STORER"
)

// Scan scans the values in the current class
func (c *Type) Scan(value interface{}) error {
	asBytes, ok := value.([]byte)
	if !ok {
		return errors.New("scan source is not []byte")
	}
	*c = Type(string(asBytes))
	return nil
}

// Value returns the string representation of the dirver value
func (c Type) Value() (driver.Value, error) {
	return string(c), nil
}

// Node represents a generic node in the playground canvas
type Node struct {
	gorm.Model
	Playground      Playground `gorm:"foreignkey:PlaygroundID;auto_preload"`
	PlaygroundID    uint       `json:"playground_id"`
	Active          bool       `json:"active"`
	Type            Type       `json:"type" sql:"not null;type:ENUM('COLLECTOR', 'ENRICHER', 'JOINER', 'OPERATOR', 'INTELLECTOR', 'STORER')"`
	Name            string     `json:"name"`
	XCoordinate     float64    `json:"x"`
	YCoordinate     float64    `json:"y"`
	Configuration   Configuration
	ConfigurationID uint
}

// GetNodes returns all the nodes from a playground
func GetNodes(db *gorm.DB, playgroundID int) ([]Node, error) {
	var nodes []Node
	if err := db.Where("playground_id = ?", playgroundID).Find(&nodes).Error; err != nil {
		return nil, err
	}
	return nodes, nil
}

// GetNode returns a single node given an ID
func GetNode(db *gorm.DB, ID int) (*Node, error) {
	var node Node
	if err := db.Where("id = ?", ID).Find(&node).Error; err != nil {
		return nil, err
	}
	return &node, nil
}

// CreateNode creates a new Node in the playground
func CreateNode(db *gorm.DB, n *Node) (*Node, error) {
	if err := db.Create(&n).Error; err != nil {
		return nil, err
	}
	return n, nil
}

// UpdateNode updates the information of the node specified by the ID
func UpdateNode(db *gorm.DB, n *Node) error {
	if err := db.Save(&n).Error; err != nil {
		return err
	}
	return nil
}

// DeleteNode deletes the node specified by the ID
func DeleteNode(db *gorm.DB, ID int) error {
	if err := db.Where("id = ?", ID).Delete(Node{}).Error; err != nil {
		return err
	}
	return nil
}
