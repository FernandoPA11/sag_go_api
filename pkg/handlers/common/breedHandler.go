package common

import (
	"SAG_GO_API/core/db"
	"SAG_GO_API/pkg/resources"
	"SAG_GO_API/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBreeds(w http.ResponseWriter, r *http.Request) {
	var breeds []resources.Breed
	db.DB.Unscoped().Find(&breeds)
	if len(breeds) == 0 {
		utils.Respond(w, http.StatusNotFound, "No breeds found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Breeds found", breeds)
}

func GetBreedByID(w http.ResponseWriter, r *http.Request) {
	breedID := mux.Vars(r)["id"]
	var breed resources.Breed
	db.DB.Unscoped().First(&breed, breedID)
	if breed.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Breed not found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Breed found", breed)
}

func AddBreed(w http.ResponseWriter, r *http.Request) {
	var breed resources.Breed
	createdBreed := db.DB.Create(&breed)
	if createdBreed.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error creating breed: "+createdBreed.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusCreated, "Breed created", breed)
}

func UpdateBreed(w http.ResponseWriter, r *http.Request) {
	breedID := mux.Vars(r)["id"]
	var breed resources.Breed
	db.DB.Unscoped().First(&breed, breedID)
	if breed.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Breed not found", nil)
		return
	}
	updatedBreed := db.DB.Save(&breed)
	if updatedBreed.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error updating breed: "+updatedBreed.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Breed updated", breed)
}

func DeleteBreed(w http.ResponseWriter, r *http.Request) {
	breedID := mux.Vars(r)["id"]
	var breed resources.Breed
	db.DB.Unscoped().First(&breed, breedID)
	if breed.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Breed not found", nil)
		return
	}
	db.DB.Unscoped().Delete(&breed)
	utils.Respond(w, http.StatusOK, "Breed deleted", nil)
}

func DisableBreed(w http.ResponseWriter, r *http.Request) {
	breedID := mux.Vars(r)["id"]
	var breed resources.Breed
	db.DB.First(&breed, breedID)
	if breed.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Breed not found", nil)
		return
	}
	db.DB.Model(&breed).Update("disabled", true)
	utils.Respond(w, http.StatusOK, "Breed disabled", nil)
}

func EnableBreed(w http.ResponseWriter, r *http.Request) {
	breedID := mux.Vars(r)["id"]
	var breed resources.Breed
	db.DB.Unscoped().First(&breed, breedID)
	if breed.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Breed not found", nil)
		return
	}
	db.DB.Model(&breed).Unscoped().Update("deleted_at", nil)
	utils.Respond(w, http.StatusOK, "Breed enabled", nil)
}

func GetDisabledBreeds(w http.ResponseWriter, r *http.Request) {
	var breeds []resources.Breed
	db.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&breeds)
	if len(breeds) == 0 {
		utils.Respond(w, http.StatusNotFound, "No disabled breeds found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Disabled breeds found", breeds)
}
