package models

import (
	"bytes"
	"database/sql/driver"
	"errors"

	"github.com/jinzhu/gorm"
)

// JSON is the type for representing json object in the database
type JSON []byte

// Value gets the driver value of the json object
func (j JSON) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}

// Scan scans the values inside the value parameter
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("Invalid Scan Source")
	}
	*j = append((*j)[0:0], s...)
	return nil
}

// MarshalJSON returns the byte array representation of the json object
func (j JSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

// UnmarshalJSON unmarshall the JSON into a byte array
func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("null point exception")
	}
	*j = append((*j)[0:0], data...)
	return nil
}

// IsNull checks if the json object is empty
func (j JSON) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}

// Equals checks the equality between two json objects
func (j JSON) Equals(j1 JSON) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}

// Configuration will hold the config info of a field and the reference of the credentials in Vault (the secret name only)
type Configuration struct {
	gorm.Model
	Blob JSON `sql:"type:json" json:"blob"`
}

// GetConfiguration will retrieve the configration for the given ID from Vault
func GetConfiguration(db *gorm.DB, ID int) {

}

// CreateConfiguration creates a new configration in Vault and the DB
func CreateConfiguration(db *gorm.DB) {

}

// UpdateConfiguration updates the configuration in Vault and DB for the given ID
func UpdateConfiguration(db *gorm.DB, ID int) {

}

// DeleteConfiguration deletes the configuration object in Vault and in the DB for the given ID
func DeleteConfiguration(db *gorm.DB, ID int) {

}
