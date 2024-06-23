package common

import (
	"SAG_GO_API/core/db"
	"SAG_GO_API/pkg/resources"
	"SAG_GO_API/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRanchs(w http.ResponseWriter, r *http.Request) {
	var ranchs []resources.Ranch
	db.DB.Find(&ranchs)
	if len(ranchs) == 0 {
		utils.Respond(w, http.StatusNotFound, "No ranchs found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Ranchs found", ranchs)
}

func GetRanchByID(w http.ResponseWriter, r *http.Request) {
	ranchID := mux.Vars(r)["id"]
	var ranch resources.Ranch
	db.DB.First(&ranch, ranchID)
	if ranch.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Ranch not found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Ranch found", ranch)
}

func AddRanch(w http.ResponseWriter, r *http.Request) {
	var ranch resources.Ranch
	json.NewDecoder(r.Body).Decode(&ranch)

	createdRanch := db.DB.Create(&ranch)
	if createdRanch.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error creating ranch: "+createdRanch.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusCreated, "Ranch created: "+ranch.Name, ranch)
}

func UpdateRanch(w http.ResponseWriter, r *http.Request) {
	ranchID := mux.Vars(r)["id"]
	var ranch resources.Ranch
	db.DB.Unscoped().First(&ranch, ranchID)
	if ranch.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Ranch not found", nil)
		return
	}
	json.NewDecoder(r.Body).Decode(&ranch)

	updatedRanch := db.DB.Save(&ranch)
	if updatedRanch.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error updating ranch: "+updatedRanch.Error.Error(), nil)
		return
	}
}

func DeleteRanch(w http.ResponseWriter, r *http.Request) {
	ranchID := mux.Vars(r)["id"]
	var ranch resources.Ranch
	db.DB.Unscoped().First(&ranch, ranchID)
	if ranch.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Ranch not found", nil)
		return
	}
	db.DB.Unscoped().Delete(&ranch)
	utils.Respond(w, http.StatusOK, "Ranch deleted", nil)
}

func DiseableRanch(w http.ResponseWriter, r *http.Request) {
	ranchID := mux.Vars(r)["id"]
	var ranch resources.Ranch
	db.DB.First(&ranch, ranchID)
	if ranch.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Ranch not found", nil)
		return
	}
	db.DB.Delete(&ranch)
	utils.Respond(w, http.StatusOK, "Ranch disabled", nil)
}

func EnableRanch(w http.ResponseWriter, r *http.Request) {
	ranchID := mux.Vars(r)["id"]
	var ranch resources.Ranch
	db.DB.Unscoped().First(&ranch, ranchID)
	if ranch.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Ranch not found", nil)
		return
	}
	db.DB.Model(&ranch).Unscoped().Update("deleted_at", nil)
	utils.Respond(w, http.StatusOK, "Ranch enabled", nil)
}

func GetDisabledRanchs(w http.ResponseWriter, r *http.Request) {
	var ranchs []resources.Ranch
	db.DB.Unscoped().Find(&ranchs)
	if len(ranchs) == 0 {
		utils.Respond(w, http.StatusNotFound, "No ranchs found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Ranchs found", ranchs)
}
