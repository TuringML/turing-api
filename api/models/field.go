package models

import (
	"database/sql/driver"
	"errors"

	"github.com/jinzhu/gorm"
)

// Class specifies the class type of the Field. It transforms it in enum in SQL
type Class string

const (
	// DimensionClass specify the dimension enum
	DimensionClass Class = "DIMENSION"
	// TimestampClass specify the timestamp enum
	TimestampClass Class = "TIMESTAMP"
	// MetricClass specify the metric enum
	MetricClass Class = "METRIC"
)

// Scan scans the values in the current class
func (c *Class) Scan(value interface{}) error {
	asBytes, ok := value.([]byte)
	if !ok {
		return errors.New("scan source is not []byte")
	}
	*c = Class(string(asBytes))
	return nil
}

// Value returns the string representation of the dirver value
func (c Class) Value() (driver.Value, error) {
	return string(c), nil
}

// Field is the information holder of the dataset
type Field struct {
	gorm.Model
	Node           Node   `gorm:"foreignkey:NodeID"`
	NodeID         uint   `json:"nodeId"`
	Active         bool   `json:"active"`
	KeyID          int    `json:"key_id"`
	KeyName        string `json:"key_name"`
	KeyPrimary     bool   `json:"key_primary"`
	ValueClass     Class  `json:"value_class" sql:"not null;type:ENUM('DIMENSION', 'TIMESTAMP', 'METRIC')"`
	ValueType      string `json:"value_type"`
	ValueExample   string `json:"value_example"`
	ValueKind      string `json:"value_kind"`
	ValueDelimiter string `json:"value_delimiter"`
}

// GetFields returns all the fields of a node
func GetFields(db *gorm.DB, nodeID int) ([]Node, error) {
	return nil, nil
}

// GetField returns a single field from given an ID of a single node
func GetField(db *gorm.DB, nodeID, ID int) (*Node, error) {
	return nil, nil
}

// CreateField creates a new Field in the node
func CreateField(db *gorm.DB, nodeID int, f Field) (*Node, error) {
	return nil, nil
}

// UpdateField updates the information of the field specified by the ID
func UpdateField(db *gorm.DB, nodeID, ID int, f Field) error {
	return nil
}

// DeleteField deletes the field specified by the ID
func DeleteField(db *gorm.DB, nodeID, ID int) error {
	return nil
}
