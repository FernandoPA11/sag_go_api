package herd

import (
	"SAG_GO_API/core/db"
	"SAG_GO_API/pkg/resources"
	"SAG_GO_API/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetHerd(w http.ResponseWriter, r *http.Request) {
	var herd []resources.Herd
	db.DB.Find(&herd)
	if len(herd) == 0 {
		utils.Respond(w, http.StatusNotFound, "No herd found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "herd found", herd)
}

func GetBovineByID(w http.ResponseWriter, r *http.Request) {
	herdID := mux.Vars(r)["id"]
	var herd resources.Herd
	db.DB.First(&herd, herdID)
	if herd.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "herd not found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "herd found", herd)
}

func AddBovine(w http.ResponseWriter, r *http.Request) {
	var herd resources.Herd
	json.NewDecoder(r.Body).Decode(&herd)

	createdHerd := db.DB.Create(&herd)
	if createdHerd.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error creating herd: "+createdHerd.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusCreated, "herd created: "+herd.TagNumber, herd)
}

func UpdateBovine(w http.ResponseWriter, r *http.Request) {
	var herd resources.Herd
	json.NewDecoder(r.Body).Decode(&herd)

	db.DB.Unscoped().First(&herd, herd.ID)
	if herd.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "herd not found", nil)
		return
	}

	updatedHerd := db.DB.Save(&herd)
	if updatedHerd.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error updating herd: "+updatedHerd.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusOK, "herd updated: "+herd.TagNumber, herd)
}

func DeleteBovine(w http.ResponseWriter, r *http.Request) {
	herdID := mux.Vars(r)["id"]
	var herd resources.Herd
	db.DB.Unscoped().First(&herd, herdID)
	if herd.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "herd not found", nil)
		return
	}
	db.DB.Unscoped().Delete(&herd)
	utils.Respond(w, http.StatusOK, "herd deleted", nil)
}

func DisableBovine(w http.ResponseWriter, r *http.Request) {
	herdID := mux.Vars(r)["id"]
	var herd resources.Herd
	db.DB.First(&herd, herdID)
	if herd.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "herd not found", nil)
		return
	}
	db.DB.Delete(&herd)
	utils.Respond(w, http.StatusOK, "herd disabled", nil)
}

func EnableBovine(w http.ResponseWriter, r *http.Request) {
	herdID := mux.Vars(r)["id"]
	var herd resources.Herd
	db.DB.Unscoped().First(&herd, herdID)
	if herd.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "herd not found", nil)
		return
	}
	db.DB.Model(&herd).Unscoped().Update("deleted_at", nil)
	utils.Respond(w, http.StatusOK, "herd enabled", nil)
}

func GetBovinesDisabled(w http.ResponseWriter, r *http.Request) {
	var herd []resources.Herd
	db.DB.Unscoped().Find(&herd).Where("deleted_at IS NOT NULL")
	if len(herd) == 0 {
		utils.Respond(w, http.StatusNotFound, "No herd found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "herd found", herd)
}
