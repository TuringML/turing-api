package models

import (
	"github.com/jinzhu/gorm"
)

// Configuration will hold the config info of a field and the reference of the credentials in Vault (the secret name only)
type Configuration struct {
	gorm.Model
	Blob []byte `sql:"type:longblob" json:"blob"`
}

// GetConfiguration will retrieve the configration for the given ID from Vault
func GetConfiguration(db *gorm.DB, ID int) (*Configuration, error) {
	var c Configuration
	if err := db.Where("id = ?", ID).Find(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

// CreateConfiguration creates a new configration in Vault and the DB
func CreateConfiguration(db *gorm.DB, c *Configuration) (*Configuration, error) {
	if err := db.Create(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
}

// UpdateConfiguration updates the configuration in Vault and DB for the given ID
func UpdateConfiguration(db *gorm.DB, c *Configuration) error {
	if err := db.Save(&c).Error; err != nil {
		return err
	}
	return nil
}

// DeleteConfiguration deletes the configuration object in Vault and in the DB for the given ID
func DeleteConfiguration(db *gorm.DB, ID int) error {
	if err := db.Where("id = ?", ID).Delete(Configuration{}).Error; err != nil {
		return err
	}
	return nil
}
