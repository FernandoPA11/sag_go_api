package user

import (
	"SAG_GO_API/core/db"
	"SAG_GO_API/pkg/resources"
	"SAG_GO_API/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []resources.User
	db.DB.Find(&users)
	if len(users) == 0 {
		utils.Respond(w, http.StatusNotFound, "No users found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Users found", users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]
	var user resources.User
	db.DB.First(&user, userID)
	if user.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "User not found", nil)
		return
	}
	json.NewEncoder(w).Encode(&user)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user resources.User
	json.NewDecoder(r.Body).Decode(&user)
	//User validation
	ValidateUser(w, user)

	createdUser := db.DB.Create(&user)
	if createdUser.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error creating user: "+createdUser.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusCreated, "User created: "+user.Username, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]
	var user resources.User
	db.DB.Unscoped().First(&user, userID)
	if user.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "User not found", nil)
		return
	}
	json.NewDecoder(r.Body).Decode(&user)

	//User validation
	ValidateUser(w, user)

	updatedUser := db.DB.Save(&user)
	if updatedUser.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error updating user: "+updatedUser.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusOK, "User updated: "+user.Username, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]
	var user resources.User
	db.DB.Unscoped().First(&user, userID)
	if user.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "User not found", nil)
		return
	}
	db.DB.Unscoped().Delete(&user)
	utils.Respond(w, http.StatusOK, "User deleted: "+user.Username, nil)
}

func DisableUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]
	var user resources.User
	db.DB.First(&user, userID)
	if user.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "User not found", nil)
		return
	}
	db.DB.Delete(&user)
	utils.Respond(w, http.StatusOK, "User disabled: "+user.Username, nil)
}

func EnableUser(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["id"]
	var user resources.User
	db.DB.Unscoped().First(&user, userID)
	if user.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "User not found", nil)
		return
	}
	db.DB.Model(&user).UpdateColumn("deleted_at", nil)
	utils.Respond(w, http.StatusOK, "User enabled: "+user.Username, nil)
}

func GetUsersDisabled(w http.ResponseWriter, r *http.Request) {
	var users []resources.User
	db.DB.Unscoped().Find(&users).Where("deleted_at IS NOT NULL")
	if len(users) == 0 {
		utils.Respond(w, http.StatusNotFound, "No users found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Users found", users)
}

func ValidateUser(w http.ResponseWriter, user resources.User) {
	if !utils.IsEmail(user.Email) {
		utils.Respond(w, http.StatusBadRequest, "Invalid email", nil)
		return
	}
	if !utils.IsValidPassword(user.Password) {
		utils.Respond(w, http.StatusBadRequest, "Password must be at least 6 characters", nil)
		return
	}
}
