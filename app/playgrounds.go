package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turing-ml/turing-api/models"
)

// GetAllPlaygrounds returns all the playgrounds from the database
// @Title GetAllPlaygrounds
// @Description Get all the playgrounds in the database
// @Accept  json
// @Success 200 {array}  models.Playground
// @Failure 500 {string} string    "Internal Server Error"
// @Resource /playground
// @Router /playgrounds [get]
func (a *App) GetAllPlaygrounds(w http.ResponseWriter, r *http.Request) {
	playgrounds, err := models.GetAllPlaygrounds(*a.DB)
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	Response(w, http.StatusOK, playgrounds)
}

// GetPlaygrounds returns all the playground objects of the user in input
// @Title GetPlaygrounds
// @Description Get all the playgrounds based on the user ID
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "The user ID"
// @Success 200 {array} models.Playground Array of playgrounds
// @Failure 500 {string} string	"Internal Server Error"
// @Router /users/{id}/playgrounds [get]
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
// @Title GetPlayground
// @Description Get a specific playground based on the ID in input of a single user
// @Accept  json
// @Produce  json
// @Param   user_id     path    string     true        "The user ID"
// @Param   playground_id     path    string     true        "The playground ID"
// @Success 200 {object} models.Playground
// @Failure 500 {string} string	"Internal Server Error"
// @Router /users/{user_id}/playgrounds/{playground_id} [get]
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
// @Title CreatePlayground
// @Description Create a new playground based on the parameters in input
// @Accept  json
// @Produce  json
// @Param   user_id     path    string     true        "The user ID"
// @Param   pipeline_schema     query    string     true        "The playground pipeline schema as JSON string"
// @Success 200 {object} models.Playground
// @Failure 500 {string} string	"Internal Server Error"
// @Router /users/{user_id}/playgrounds [post]
func (a *App) CreatePlayground(w http.ResponseWriter, r *http.Request) {
	var p models.Playground
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	playground, err := models.CreatePlayground(*a.DB, vars["id"], p.PipelineSchema)
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	Response(w, http.StatusOK, playground)
}

// UpdatePlayground updates the playground's properties in the database
// @Title UpdatePlayground
// @Description Update the playground of a specific user
// @Accept  json
// @Produce  json
// @Param   user_id     path    string     true        "The user ID"
// @Param   playground_id     path    string     true        "The playground ID"
// @Param   pipeline_schema     query    string     true        "The playground pipeline schema as JSON string to update"
// @Success 200 {string} string	"playground updated"
// @Failure 500 {string} string	"Internal Server Error"
// @Router /users/{user_id}/playgrounds/{playground_id} [put]
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
// @Title DeletePlayground
// @Description Delete the playground of a specific user
// @Accept  json
// @Produce  json
// @Param   user_id     path    string     true        "The user ID"
// @Param   playground_id     path    string     true        "The playground ID"
// @Success 200 {string} string	"playground deleted"
// @Failure 500 {string} string	"Internal Server Error"
// @Router /users/{user_id}/playgrounds/{playground_id} [delete]
func (a *App) DeletePlayground(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := models.DeletePlayground(*a.DB, vars["id"], vars["playgroundId"])
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	Response(w, http.StatusOK, map[string]string{"message": "playground deleted"})
}
