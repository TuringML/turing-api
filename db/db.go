package db

import (
	"database/sql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/zerolog/log"
)

func Open(url, username, password, name string) (*sql.DB) {
	cs := fmt.Sprintf("%s:%s@%s/%s", username, password, url, name)
	db, err := sql.Open("mysql", cs)
	if err != nil {
		log.Panic().Err(err).Msg("could not connect to database")
	}
	return db
}
