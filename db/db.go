package db

import (
	"github.com/rs/zerolog/log"

	mgo "gopkg.in/mgo.v2"
)

// Open creates a new connection to the selected database
func Open(url, username, password, name string) *mgo.Database {
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal().Err(err).Msg("could not connect to MongoDB server")
	}
	db := session.DB(name)
	err = db.Login(username, password)
	if err != nil {
		log.Fatal().Err(err).Msg("could not connect to MongoDB database")
	}
	return db
}
