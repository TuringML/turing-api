package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turing-ml/turing-api/models"
)

// GetAllPlaygrounds returns all the playgrounds from the database
func (a *App) GetAllPlaygrounds(w http.ResponseWriter, r *http.Request) {
	playgrounds, err := models.GetAllPlaygrounds(*a.DB)
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	Response(w, http.StatusOK, playgrounds)
}

// GetPlaygrounds returns all the playground objects of the user in input
func (a *App) GetPlaygrounds(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playground, err := models.GetPlaygrounds(*a.DB, vars["id"])
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	Response(w, http.StatusOK, playground)
}

// GetPlayground return a single playground object given the ID in input
func (a *App) GetPlayground(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playground, err := models.GetPlayground(*a.DB, vars["id"], vars["playgroundId"])
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	Response(w, http.StatusOK, playground)
}

// CreatePlayground creates a new playground in the database
func (a *App) CreatePlayground(w http.ResponseWriter, r *http.Request) {
	var p models.Playground
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	err := models.CreatePlayground(*a.DB, vars["id"], p.PipelineSchema)
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	Response(w, http.StatusOK, map[string]string{"message": "playground created"})
}

// UpdatePlayground updates the playground's properties in the database
func (a *App) UpdatePlayground(w http.ResponseWriter, r *http.Request) {
	var p models.Playground
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	err := models.UpdatePlayground(*a.DB, vars["id"], vars["playgroundId"], p.PipelineSchema)
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	Response(w, http.StatusOK, map[string]string{"message": "playground updated"})
}

// DeletePlayground deletes the specified playground from the database
func (a *App) DeletePlayground(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := models.DeletePlayground(*a.DB, vars["id"], vars["playgroundId"])
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	Response(w, http.StatusOK, map[string]string{"message": "playground deleted"})
}
