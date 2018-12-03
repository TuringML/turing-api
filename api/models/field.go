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
func GetFields(db *gorm.DB, nodeID int) ([]Field, error) {
	var fields []Field
	if err := db.Where("node_id = ?", nodeID).Find(&fields).Error; err != nil {
		return nil, err
	}
	return fields, nil
}

// GetField returns a single field from given an ID of a single node
func GetField(db *gorm.DB, ID int) (*Field, error) {
	var field Field
	if err := db.Where("id = ?", ID).Find(&field).Error; err != nil {
		return nil, err
	}
	return &field, nil
}

// CreateField creates a new Field in the node
func CreateField(db *gorm.DB, f Field) (*Field, error) {
	if err := db.Create(&f).Error; err != nil {
		return nil, err
	}
	return &f, nil
}

// UpdateField updates the information of the field specified by the ID
func UpdateField(db *gorm.DB, f Field) error {
	if err := db.Save(&f).Error; err != nil {
		return err
	}
	return nil
}

// DeleteField deletes the field specified by the ID
func DeleteField(db *gorm.DB, ID int) error {
	if err := db.Where("id = ?", ID).Delete(Field{}).Error; err != nil {
		return err
	}
	return nil
}
