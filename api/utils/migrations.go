package utils

import (
	"github.com/jinzhu/gorm"
	"github.com/turing-ml/turing-api/api/models"
)

// RunMigration will execute all the migrations of the database
func RunMigration(db *gorm.DB) {

	db.AutoMigrate(
		&models.Node{}, &models.Playground{},
		&models.LinkGorm{}, &models.Field{},
		&models.Configuration{},
	)

	// BUG in AutoMigrate. Forced to run the foreign key manually
	// These lines will lead to an error when starting the APIs but I can safely ignore it
	db.Model(&models.Field{}).AddForeignKey("node_id", "nodes(id)", "CASCADE", "CASCADE")
	db.Model(&models.LinkGorm{}).AddForeignKey("from_node_id", "nodes(id)", "CASCADE", "CASCADE")
	db.Model(&models.LinkGorm{}).AddForeignKey("from_field_id", "fields(id)", "CASCADE", "CASCADE")
	db.Model(&models.LinkGorm{}).AddForeignKey("to_node_id", "nodes(id)", "CASCADE", "CASCADE")
	db.Model(&models.LinkGorm{}).AddForeignKey("to_field_id", "fields(id)", "CASCADE", "CASCADE")
	db.Model(&models.Node{}).AddForeignKey("playground_id", "playgrounds(id)", "CASCADE", "CASCADE")
}
