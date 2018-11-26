package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turing-ml/turing-api/models"
)

// GetUsers returns all the users from the database
// @Title GetUsers
// @Description Get all the users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User array of users
// @Router /users [get]
func (a *App) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetUsers(*a.DB)
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	Response(w, http.StatusOK, users)
}

// GetUser return a single user object given the ID in input
// @Title GetUser
// @Description Get user object by ID
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Some ID"
// @Success 200 {object} models.User
// @Failure 400 {string} string "ID is mandatory for this endpoint"
// @Failure 500 {string} string "Internal Server Error"
// @Router /users/{id} [get]
func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := models.GetUser(*a.DB, vars["id"])
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	Response(w, http.StatusOK, user)
}

// CreateUser creates a new user in the database
// @Title CreateUser
// @Description Create a new user
// @Accept  json
// @Param   name        query   string     true        "Name of the user"
// @Success 200 {object}  models.User
// @Failure 500 {string} string    "Internal Server Error"
// @Resource /user
// @Router /users/ [post]
func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	defer r.Body.Close()

	user, err := models.CreateUser(*a.DB, u.Name)
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	Response(w, http.StatusOK, user)
}

// UpdateUser updates the user's properties in the database
// @Title UpdateUser
// @Description Update the user specified by the ID
// @Accept  json
// @Param   id        path   string     true        "ID of the user - the MongoDB Object ID"
// @Param   name        query   string     true        "New name of the user"
// @Success 200 {string}  string	"user updated"
// @Failure 500 {string} string    "Internal Server Error"
// @Resource /user
// @Router /users/{id} [put]
func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	err := models.UpdateUser(*a.DB, vars["id"], u.Name)
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	Response(w, http.StatusOK, map[string]string{"message": "user updated"})
}

// DeleteUser deletes the specified user from the database
// @Title DeleteUser
// @Description Delete the user specified by the ID
// @Accept  json
// @Param   id        path   string     true        "ID of the user - the MongoDB Object ID"
// @Success 200 {string}  string	"user deleted"
// @Failure 500 {string} string    "Internal Server Error"
// @Resource /user
// @Router /users/{id} [delete]
func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := models.DeleteUser(*a.DB, vars["id"])
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	Response(w, http.StatusOK, map[string]string{"message": "user deleted"})
}
