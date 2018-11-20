package api

import (
	"github.com/gorilla/mux"
	"github.com/turing-ml/turing-api/middlewares"
)

type Api struct {
	AuthSecret string
}

func NewApi(secret string) *Api {
	return &Api{AuthSecret: secret}
}

func (a *Api) NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	// set up global middlewares
	r.Use(middlewares.Logging)

	r.HandleFunc("/info", Info).Methods("GET")

	return r
}