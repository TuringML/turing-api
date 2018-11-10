package db

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/serenize/snaker"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Connect() *gorm.DB {
	dir := filepath.Dir("db/database.db")
	db, err := gorm.Open("sqlite3", dir+"/database.db")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	db.LogMode(false)

	if gin.IsDebugging() {
		db.LogMode(true)
	}

	if os.Getenv("AUTOMIGRATE") == "1" {
		db.AutoMigrate(

		)
	}

	return db
}

func GetInstance(c *gin.Context) *gorm.DB {
	return c.MustGet("DB").(*gorm.DB)
}

func (s *Parameter) SetPreloads(db *gorm.DB) *gorm.DB {
	if s.Preloads == "" {
		return db
	}

	for _, preload := range strings.Split(s.Preloads, ",") {
		var a []string

		for _, s := range strings.Split(preload, ".") {
			a = append(a, snaker.SnakeToCamel(s))
		}

		db = db.Preload(strings.Join(a, "."))
	}

	return db
}
