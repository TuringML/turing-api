package app

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/turing-ml/turing-api/api"
	"github.com/turing-ml/turing-api/db"
	"net/http"
)

type App struct {
	Api *mux.Router
	DB  *sql.DB
}

func NewApp(secret, dbUrl, dbUsername, dbPassword, dbName string) (*App, error) {

	a := api.NewApi(secret)

	return &App{
		Api: a.NewRouter(),
		DB: db.Open(dbUrl, dbUsername, dbPassword, dbName),
	}, nil
}

func (a *App) Serve(addr string) error {
	err := http.ListenAndServe(addr, a.Api)
	if err != nil {
		return err
	}
	return nil
}
