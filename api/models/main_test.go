package models

import (
	"os"
	"testing"

	"github.com/turing-ml/turing-api/pkg/database"

	"github.com/jinzhu/gorm"
)

var testDb *gorm.DB

func TestMain(m *testing.M) {
	tearUp()
	c := m.Run()
	tearDown()
	os.Exit(c)
}

func tearUp() {
	testDb = database.OpenConnection("turing", "turing", "192.168.99.100", "testing")
}

func tearDown() {

}
