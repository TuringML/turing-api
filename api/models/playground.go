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
	if err := db.Find(&playgrounds).Error; err != nil {
		return nil, err
	}
	return playgrounds, nil
}

// GetPlaygrounds returns all the playgrounds from the database of a given userID
func GetPlaygrounds(db *gorm.DB, userID string) ([]Playground, error) {
	var playgrounds []Playground
	if err := db.Where("user_id = ?", userID).Find(&playgrounds).Error; err != nil {
		return nil, err
	}
	return playgrounds, nil
}

// GetPlayground return a single playground object given the ID in input
func GetPlayground(db *gorm.DB, userID string, ID int) (*Playground, error) {
	var playground Playground
	if err := db.Where("id = ? AND user_id = ?", ID, userID).First(&playground).Error; err != nil {
		return nil, err
	}
	return &playground, nil
}

// CreatePlayground creates a new playground in the database
func CreatePlayground(db *gorm.DB, userID string, p Playground) (*Playground, error) {
	if err := db.Create(p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

// UpdatePlayground updates the playground's properties in the database
func UpdatePlayground(db *gorm.DB, userID string, ID int, p Playground) error {
	if err := db.Where("id = ? AND user_id = ?", ID, userID).Save(p).Error; err != nil {
		return err
	}
	return nil
}

// DeletePlayground deletes the specified playground from the database
func DeletePlayground(db *gorm.DB, userID string, ID int) error {
	if err := db.Where("id = ? AND user_id = ?", ID, userID).Delete(Playground{}).Error; err != nil {
		return err
	}
	return nil
}
