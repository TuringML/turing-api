package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// OpenConnection will open up the connection with the MySQL database
func OpenConnection(dbUsername, dbPassword, dbHost, dbName string, debug bool) *gorm.DB {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", dbUsername, dbPassword, dbHost, dbName)

	db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}

	db.LogMode(debug)
	db.DB().SetMaxIdleConns(10)

	return db
}
