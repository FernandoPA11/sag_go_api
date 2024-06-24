package common

import (
	"SAG_GO_API/core/db"
	"SAG_GO_API/pkg/resources"
	"SAG_GO_API/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAnimalTypes(w http.ResponseWriter, r *http.Request) {
	var animalTypes []resources.AnimalType
	db.DB.Unscoped().Find(&animalTypes)
	if len(animalTypes) == 0 {
		utils.Respond(w, http.StatusNotFound, "No animal types found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Animal types found", animalTypes)
}

func GetAnimalTypeByID(w http.ResponseWriter, r *http.Request) {
	animalTypeID := mux.Vars(r)["id"]
	var animalType resources.AnimalType
	db.DB.Unscoped().First(&animalType, animalTypeID)
	if animalType.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Animal type not found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Animal type found", animalType)
}

func GetAnimalTypeByName(w http.ResponseWriter, r *http.Request) {
	animalTypeName := mux.Vars(r)["id"]
	var animalType resources.AnimalType
	db.DB.Unscoped().Where("name LIKE ?", animalTypeName).First(&animalType)
	if animalType.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Animal type not found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Animal type found", animalType)
}

func AddAnimalType(w http.ResponseWriter, r *http.Request) {
	var animalType resources.AnimalType
	createdAnimalType := db.DB.Create(&animalType)
	if createdAnimalType.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error creating animal type: "+createdAnimalType.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusCreated, "Animal type created", animalType)
}

func UpdateAnimalType(w http.ResponseWriter, r *http.Request) {
	animalTypeID := mux.Vars(r)["id"]
	var animalType resources.AnimalType
	db.DB.Unscoped().First(&animalType, animalTypeID)
	if animalType.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Animal type not found", nil)
		return
	}
	updatedAnimalType := db.DB.Save(&animalType)
	if updatedAnimalType.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error updating animal type: "+updatedAnimalType.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Animal type updated", animalType)
}

func DeleteAnimalType(w http.ResponseWriter, r *http.Request) {
	animalTypeID := mux.Vars(r)["id"]
	var animalType resources.AnimalType
	db.DB.Unscoped().First(&animalType, animalTypeID)
	if animalType.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Animal type not found", nil)
		return
	}
	db.DB.Unscoped().Delete(&animalType)
	utils.Respond(w, http.StatusOK, "Animal type deleted", nil)
}

func DisableAnimalType(w http.ResponseWriter, r *http.Request) {
	animalTypeID := mux.Vars(r)["id"]
	var animalType resources.AnimalType
	db.DB.First(&animalType, animalTypeID)
	if animalType.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Animal type not found", nil)
		return
	}
	db.DB.Delete(&animalType)
	utils.Respond(w, http.StatusOK, "Animal type disabled", nil)
}

func EnableAnimalType(w http.ResponseWriter, r *http.Request) {
	animalTypeID := mux.Vars(r)["id"]
	var animalType resources.AnimalType
	db.DB.Unscoped().First(&animalType, animalTypeID)
	if animalType.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Animal type not found", nil)
		return
	}
	db.DB.Model(&animalType).Update("deleted_at", nil)
	utils.Respond(w, http.StatusOK, "Animal type enabled", nil)
}

func GetDisabledAnimalTypes(w http.ResponseWriter, r *http.Request) {
	var animalTypes []resources.AnimalType
	db.DB.Unscoped().Where("deleted_at IS NOT NULL").Find(&animalTypes)
	if len(animalTypes) == 0 {
		utils.Respond(w, http.StatusNotFound, "No disabled animal types found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Disabled animal types found", animalTypes)
}
