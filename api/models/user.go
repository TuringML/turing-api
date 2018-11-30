package models

import (
	"github.com/jinzhu/gorm"
)

// User is the struct used to deal with the users in the database
type User struct {
	gorm.Model
	Name string `json:"name"`
}

// GetUsers returns all the users from the database
func GetUsers(db *gorm.DB) ([]User, error) {
	var users []User
	return users, nil
}

// GetUser return a single user object given the ID in input
func GetUser(db *gorm.DB, ID string) (*User, error) {
	var user User
	return &user, nil
}
