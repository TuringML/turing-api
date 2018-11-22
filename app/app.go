package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turing-ml/turing-api/db"
	"github.com/turing-ml/turing-api/middlewares"
	"gopkg.in/mgo.v2"
)

// App is the struct that encapsulate all the services necessary to run the application
type App struct {
	Router *mux.Router
	DB     *mgo.Database
}

// NewApp instantiates a new application
func NewApp(secret, dbURL, dbUsername, dbPassword, dbName string) (*App, error) {
	app := &App{
		Router: mux.NewRouter().StrictSlash(true),
		DB:     db.Open(dbURL, dbUsername, dbPassword, dbName),
	}
	app.InitializeRoutes()
	return app, nil
}

// Serve will start listening on the assigned address for requests
func (a *App) Serve(addr string) error {
	err := http.ListenAndServe(addr, a.Router)
	if err != nil {
		return err
	}
	return nil
}

// InitializeRoutes creates a new Gorilla/Mux router object with the endpoints
func (a *App) InitializeRoutes() {

	// set up global middlewares
	a.Router.Use(middlewares.Logging)

	a.Router.HandleFunc("/info", Info).Methods("GET")
	a.Router.HandleFunc("/users", a.GetUsers).Methods("GET")
	a.Router.HandleFunc("/users", a.CreateUser).Methods("POST")
	a.Router.HandleFunc("/users/{id}", a.GetUser).Methods("GET")
}
