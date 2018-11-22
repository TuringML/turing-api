package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turing-ml/turing-api/models"
)

// GetUsers returns all the users from the database
func (a *App) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetUsers(*a.DB)
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	Response(w, http.StatusOK, users)
}

// GetUser return a single user object given the ID in input
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
func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	defer r.Body.Close()

	err := models.CreateUser(*a.DB, u.Name)
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}
	Response(w, http.StatusOK, map[string]string{"message": "user created"})
}

// UpdateUser updates the user's properties in the database
func (a *App) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := models.UpdateUser(*a.DB, vars["id"], "")
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
	}
	Response(w, http.StatusOK, map[string]string{"message": "user updated"})
}

// DeleteUser deletes the specified user from the database
func (a *App) DeleteUser(w http.ResponseWriter, r *http.Request) {
	err := models.DeleteUser(*a.DB, "")
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
	}
	Response(w, http.StatusOK, map[string]string{"message": "user deleted"})
}
