package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gamorvi/restapi2/app/models"
	u "github.com/gamorvi/restapi2/utils"
	"github.com/gorilla/mux"
)

// Get one user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	user := models.GetUser(id)
	if user == nil {
		u.Respond(w, u.Message(false, "User not found"))
		return
	}

	resp := u.Message(true, "success")
	resp["data"] = user
	u.Respond(w, resp)
	return
}

// Get all the users in the users table
func GetUsers(w http.ResponseWriter, r *http.Request) {
	resp := u.Message(true, "success")
	users := models.GetUsers()
	if users == nil {
		u.Respond(w, u.Message(false, "No users found"))
		return
	}
	resp["data"] = users
	u.Respond(w, resp)
	return
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	defer r.Body.Close()

	resp := user.Create()
	u.Respond(w, resp)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var user models.User
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	err = models.GetUserForUpdateOrDelete(id, &user)
	if err != nil {
		u.Respond(w, u.Message(false, "User not found"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	user.ID = uint(id)
	user.UpdatedAt = time.Now().Local()
	defer r.Body.Close()

	// Update user here
	err = models.UpdateUser(&user)
	if err != nil {
		u.Respond(w, u.Message(false, "Could not update the record"))
		return
	}
	resp := u.Message(true, "Updated successfully")
	resp["data"] = user
	u.Respond(w, resp)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user models.User
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(w, u.Message(false, "There was an error in your request"))
		return
	}

	err = models.GetUserForUpdateOrDelete(id, &user)
	if err != nil {
		u.Respond(w, u.Message(false, "User not found"))
		return
	}

	err = models.DeleteUser(&user)
	if err != nil {
		u.Respond(w, u.Message(false, "Could not delete the record"))
		return
	}
	u.Respond(w, u.Message(true, "User has been deleted successfully"))
	return
}
