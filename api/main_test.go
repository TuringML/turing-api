package api

import (
	"os"
	"testing"
	"time"

	"github.com/turing-ml/turing-api/api/models"

	"github.com/bouk/monkey"
	"github.com/jinzhu/gorm"
)

func TestMain(m *testing.M) {

	db, err := gorm.Open("mysql", "root:root@tcp(192.168.99.100:3306)/")
	if err != nil {
		panic(err)
	}

	db.Exec("CREATE DATABASE IF NOT EXISTS turing_test")
	db.Exec("USE turing_test")

	db.AutoMigrate(&models.User{})

	// Reset Primary Key
	db.Exec("TRUNCATE TABLE users")

	os.Exit(m.Run())
}

func FreezeTime() *monkey.PatchGuard {
	freezeTime := time.Date(2018, time.June, 1, 10, 0, 0, 0, time.UTC)
	return monkey.Patch(time.Now, func() time.Time { return freezeTime })
}
