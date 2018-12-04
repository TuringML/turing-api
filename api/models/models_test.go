package models

import (
	"os"
	"testing"

	"github.com/turing-ml/turing-api/pkg/database"
	testfixtures "gopkg.in/testfixtures.v2"

	_ "github.com/go-sql-driver/mysql"
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
	// access as root and create testing db
	testDb = database.OpenConnection("root", "root", "192.168.99.100", "", false)

	testDb = testDb.Exec("CREATE DATABASE IF NOT EXISTS testing")
	testDb = testDb.Exec("USE testing")

	// run migration here
	runMigration()

	fixtures, err := testfixtures.NewFolder(testDb.DB(), &testfixtures.MySQL{}, "../fixtures")
	if err != nil {
		panic(err)
	}

	err = fixtures.Load()
	if err != nil {
		panic(err)
	}
}

func tearDown() {
	testDb.Exec("DROP DATABASE testing")
}

func runMigration() {

	testDb.AutoMigrate(
		&Node{}, &Playground{},
		&Link{}, &Field{},
		&Configuration{},
	)

	// BUG in AutoMigrate. Forced to run the foreign key manually
	// These lines will lead to an error when starting the APIs but I can safely ignore it
	testDb.Model(&Field{}).AddForeignKey("node_id", "nodes(id)", "CASCADE", "CASCADE")
	testDb.Model(&Link{}).AddForeignKey("from_node_id", "nodes(id)", "CASCADE", "CASCADE")
	testDb.Model(&Link{}).AddForeignKey("from_field_id", "fields(id)", "CASCADE", "CASCADE")
	testDb.Model(&Link{}).AddForeignKey("to_node_id", "nodes(id)", "CASCADE", "CASCADE")
	testDb.Model(&Link{}).AddForeignKey("to_field_id", "fields(id)", "CASCADE", "CASCADE")
	testDb.Model(&Node{}).AddForeignKey("playground_id", "playgrounds(id)", "CASCADE", "CASCADE")
}
