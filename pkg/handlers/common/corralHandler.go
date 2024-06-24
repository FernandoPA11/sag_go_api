package common

import (
	"SAG_GO_API/core/db"
	"SAG_GO_API/pkg/resources"
	"SAG_GO_API/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCorrals(w http.ResponseWriter, r *http.Request) {
	var corrals []resources.Corral
	db.DB.Unscoped().Find(&corrals)
	if len(corrals) == 0 {
		utils.Respond(w, http.StatusNotFound, "No corrals found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Corrals found", corrals)
}

func GetCorralByID(w http.ResponseWriter, r *http.Request) {
	corralID := mux.Vars(r)["id"]
	var corral resources.Corral
	db.DB.Unscoped().First(&corral, corralID)
	if corral.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Corral not found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Corral found", corral)
}

func AddCorral(w http.ResponseWriter, r *http.Request) {
	var corral resources.Corral
	createdCorral := db.DB.Create(&corral)
	if createdCorral.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error creating corral: "+createdCorral.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusCreated, "Corral created", corral)
}

func UpdateCorral(w http.ResponseWriter, r *http.Request) {
	corralID := mux.Vars(r)["id"]
	var corral resources.Corral
	db.DB.Unscoped().First(&corral, corralID)
	if corral.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Corral not found", nil)
		return
	}
	updatedCorral := db.DB.Save(&corral)
	if updatedCorral.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error updating corral: "+updatedCorral.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Corral updated", corral)
}

func DeleteCorral(w http.ResponseWriter, r *http.Request) {
	corralID := mux.Vars(r)["id"]
	var corral resources.Corral
	db.DB.Unscoped().First(&corral, corralID)
	if corral.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Corral not found", nil)
		return
	}
	db.DB.Unscoped().Delete(&corral)
	utils.Respond(w, http.StatusOK, "Corral deleted", nil)
}

func DisableCorral(w http.ResponseWriter, r *http.Request) {
	corralID := mux.Vars(r)["id"]
	var corral resources.Corral
	db.DB.Unscoped().First(&corral, corralID)
	if corral.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Corral not found", nil)
		return
	}
	db.DB.Delete(&corral)
	utils.Respond(w, http.StatusOK, "Corral disabled", nil)
}

func EnableCorral(w http.ResponseWriter, r *http.Request) {
	corralID := mux.Vars(r)["id"]
	var corral resources.Corral
	db.DB.Unscoped().First(&corral, corralID)
	if corral.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Corral not found", nil)
		return
	}
	db.DB.Model(&corral).Update("deleted_at", nil)
	utils.Respond(w, http.StatusOK, "Corral enabled", nil)
}

func GetDisabledCorrals(w http.ResponseWriter, r *http.Request) {
	var corrals []resources.Corral
	db.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&corrals)
	if len(corrals) == 0 {
		utils.Respond(w, http.StatusNotFound, "No disabled corrals found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Disabled corrals found", corrals)
}
