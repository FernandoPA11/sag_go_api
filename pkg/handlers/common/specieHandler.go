package common

import (
	"SAG_GO_API/core/db"
	"SAG_GO_API/pkg/resources"
	"SAG_GO_API/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetSpecies(w http.ResponseWriter, r *http.Request) {
	var species []resources.Specie
	db.DB.Unscoped().Find(&species)
	if len(species) == 0 {
		utils.Respond(w, http.StatusNotFound, "No species found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Species found", species)
}

func GetSpecieByID(w http.ResponseWriter, r *http.Request) {
	specieID := mux.Vars(r)["id"]
	var specie resources.Specie
	db.DB.Unscoped().First(&specie, specieID)
	if specie.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Specie not found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Specie found", specie)
}

func AddSpecie(w http.ResponseWriter, r *http.Request) {
	var specie resources.Specie
	createdSpecie := db.DB.Create(&specie)
	if createdSpecie.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error creating specie: "+createdSpecie.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusCreated, "Specie created", specie)
}

func UpdateSpecie(w http.ResponseWriter, r *http.Request) {
	specieID := mux.Vars(r)["id"]
	var specie resources.Specie
	db.DB.Unscoped().First(&specie, specieID)
	if specie.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Specie not found", nil)
		return
	}
	updatedSpecie := db.DB.Save(&specie)
	if updatedSpecie.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error updating specie: "+updatedSpecie.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Specie updated", specie)
}

func DeleteSpecie(w http.ResponseWriter, r *http.Request) {
	specieID := mux.Vars(r)["id"]
	var specie resources.Specie
	db.DB.Unscoped().First(&specie, specieID)
	if specie.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Specie not found", nil)
		return
	}
	db.DB.Unscoped().Delete(&specie)
	utils.Respond(w, http.StatusOK, "Specie deleted", nil)
}

func DisableSpecie(w http.ResponseWriter, r *http.Request) {
	specieID := mux.Vars(r)["id"]
	var specie resources.Specie
	db.DB.Unscoped().First(&specie, specieID)
	if specie.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Specie not found", nil)
		return
	}
	db.DB.Delete(&specie)
	utils.Respond(w, http.StatusOK, "Specie disabled", nil)
}

func EnableSpecie(w http.ResponseWriter, r *http.Request) {
	specieID := mux.Vars(r)["id"]
	var specie resources.Specie
	db.DB.Unscoped().First(&specie, specieID)
	if specie.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Specie not found", nil)
		return
	}
	db.DB.Model(&specie).Update("deleted_at", nil)
	utils.Respond(w, http.StatusOK, "Specie enabled", nil)
}

func GetDisabledSpecies(w http.ResponseWriter, r *http.Request) {
	var species []resources.Specie
	db.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&species)
	if len(species) == 0 {
		utils.Respond(w, http.StatusNotFound, "No disabled species found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Disabled species found", species)
}
