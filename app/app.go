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

	// TODO: finish this
	// a.Router.HandleFunc("/playgrounds", GetAllPlaygrounds).Methods("GET")

	u := a.Router.PathPrefix("/users").Subrouter()
	u.HandleFunc("/", a.GetUsers).Methods("GET")
	u.HandleFunc("/", a.CreateUser).Methods("POST")
	u.HandleFunc("/{id}", a.GetUser).Methods("GET")
	u.HandleFunc("/{id}", a.UpdateUser).Methods("PUT")
	u.HandleFunc("/{id}", a.DeleteUser).Methods("DELETE")

	// TODO: finish these
	// u.HandleFunc("/{id}/playgrounds", a.CreatePlaygrounds).Methods("POST")
	// u.HandleFunc("/{id}/playgrounds/{playgroundId}", a.GetPlayground).Methods("GET")
	// u.HandleFunc("/{id}/playgrounds/{playgroundId}", a.UpdatePlayground).Methods("PUT")
	// u.HandleFunc("/{id}/playgrounds/{playgroundId}", a.DeletePlayground).Methods("DELETE")
}
