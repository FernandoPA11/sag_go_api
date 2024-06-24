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
	utils.Respond(w, http.StatusOK, "Herd found", herd)
}

func GetAnimalByID(w http.ResponseWriter, r *http.Request) {
	herdID := mux.Vars(r)["id"]
	var herd resources.Herd
	db.DB.First(&herd, herdID)
	if herd.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Animal not found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "herd found", herd)
}

func AddAnimal(w http.ResponseWriter, r *http.Request) {
	var herd resources.Herd
	json.NewDecoder(r.Body).Decode(&herd)

	createdHerd := db.DB.Create(&herd)
	if createdHerd.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error creating animal: "+createdHerd.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusCreated, "Animal created: "+herd.TagNumber, herd)
}

func UpdateAnimal(w http.ResponseWriter, r *http.Request) {
	var herd resources.Herd
	json.NewDecoder(r.Body).Decode(&herd)

	db.DB.Unscoped().First(&herd, herd.ID)
	if herd.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Animal not found", nil)
		return
	}

	updatedHerd := db.DB.Save(&herd)
	if updatedHerd.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error updating herd: "+updatedHerd.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Animal updated: "+herd.TagNumber, herd)
}

func DeleteAnimal(w http.ResponseWriter, r *http.Request) {
	herdID := mux.Vars(r)["id"]
	var herd resources.Herd
	db.DB.Unscoped().First(&herd, herdID)
	if herd.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Animal not found", nil)
		return
	}
	db.DB.Unscoped().Delete(&herd)
	utils.Respond(w, http.StatusOK, "Animal deleted", nil)
}

func DisableAnimal(w http.ResponseWriter, r *http.Request) {
	herdID := mux.Vars(r)["id"]
	var herd resources.Herd
	db.DB.First(&herd, herdID)
	if herd.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Animal already disabled or not found", nil)
		return
	}
	db.DB.Delete(&herd)
	utils.Respond(w, http.StatusOK, "Animal disabled", nil)
}

func EnableAnimal(w http.ResponseWriter, r *http.Request) {
	herdID := mux.Vars(r)["id"]
	var herd resources.Herd
	db.DB.Unscoped().First(&herd, herdID)
	if herd.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Animal not found", nil)
		return
	}
	db.DB.Model(&herd).Unscoped().Update("deleted_at", nil)
	utils.Respond(w, http.StatusOK, "Animal enabled", nil)
}

func GetDisabledAnimals(w http.ResponseWriter, r *http.Request) {
	var herd []resources.Herd
	db.DB.Unscoped().Find(&herd).Where("deleted_at IS NOT NULL")
	if len(herd) == 0 {
		utils.Respond(w, http.StatusNotFound, "No animal has been disabled", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Animals found", herd)
}
