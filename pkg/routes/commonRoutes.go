package routes

import (
	"SAG_GO_API/pkg/handlers/common"

	"github.com/gorilla/mux"
)

func SpecieRouteHandlers(router *mux.Router) {

	router.HandleFunc("/species", common.GetSpecies).Methods("GET")

	router.HandleFunc("/species/{id}", common.GetSpecieByID).Methods("GET")

	router.HandleFunc("/species", common.AddSpecie).Methods("POST")

	router.HandleFunc("/species/{id}", common.UpdateSpecie).Methods("PUT")

	router.HandleFunc("/species/{id}", common.DeleteSpecie).Methods("DELETE")

	router.HandleFunc("/species/{id}/disable", common.DisableSpecie).Methods("PUT")

	router.HandleFunc("/species/{id}/enable", common.EnableSpecie).Methods("PUT")

	router.HandleFunc("/species/disabled", common.GetDisabledSpecies).Methods("GET")
}

func BreedRouteHandlers(router *mux.Router) {

	router.HandleFunc("/breeds", common.GetBreeds).Methods("GET")

	router.HandleFunc("/breeds/{id}", common.GetBreedByID).Methods("GET")

	router.HandleFunc("/breeds", common.AddBreed).Methods("POST")

	router.HandleFunc("/breeds/{id}", common.UpdateBreed).Methods("PUT")

	router.HandleFunc("/breeds/{id}", common.DeleteBreed).Methods("DELETE")

	router.HandleFunc("/breeds/{id}/disable", common.DisableBreed).Methods("PUT")

	router.HandleFunc("/breeds/{id}/enable", common.EnableBreed).Methods("PUT")

	router.HandleFunc("/breeds/disabled", common.GetDisabledBreeds).Methods("GET")
}

func CorralRouteHandlers(router *mux.Router) {

	router.HandleFunc("/corrals", common.GetCorrals).Methods("GET")

	router.HandleFunc("/corrals/{id}", common.GetCorralByID).Methods("GET")

	router.HandleFunc("/corrals", common.AddCorral).Methods("POST")

	router.HandleFunc("/corrals/{id}", common.UpdateCorral).Methods("PUT")

	router.HandleFunc("/corrals/{id}", common.DeleteCorral).Methods("DELETE")

	router.HandleFunc("/corrals/{id}/disable", common.DisableCorral).Methods("PUT")

	router.HandleFunc("/corrals/{id}/enable", common.EnableCorral).Methods("PUT")

	router.HandleFunc("/corrals/disabled", common.GetDisabledCorrals).Methods("GET")
}

func AnimalTypeRouteHandlers(router *mux.Router) {

	router.HandleFunc("/animal-types", common.GetAnimalTypes).Methods("GET")

	router.HandleFunc("/animal-types/{id}", common.GetAnimalTypeByID).Methods("GET")

	router.HandleFunc("/animal-types", common.AddAnimalType).Methods("POST")

	router.HandleFunc("/animal-types/{id}", common.UpdateAnimalType).Methods("PUT")

	router.HandleFunc("/animal-types/{id}", common.DeleteAnimalType).Methods("DELETE")

	router.HandleFunc("/animal-types/{id}/disable", common.DisableAnimalType).Methods("PUT")

	router.HandleFunc("/animal-types/{id}/enable", common.EnableAnimalType).Methods("PUT")

	router.HandleFunc("/animal-types/disabled", common.GetDisabledAnimalTypes).Methods("GET")
}

func RanchRouteHandlers(router *mux.Router) {

	router.HandleFunc("/ranches", common.GetRanches).Methods("GET")

	router.HandleFunc("/ranches/{id}", common.GetRanchByID).Methods("GET")

	router.HandleFunc("/ranches", common.AddRanch).Methods("POST")

	router.HandleFunc("/ranches/{id}", common.UpdateRanch).Methods("PUT")

	router.HandleFunc("/ranches/{id}", common.DeleteRanch).Methods("DELETE")

	router.HandleFunc("/ranches/{id}/disable", common.DisableRanch).Methods("PUT")

	router.HandleFunc("/ranches/{id}/enable", common.EnableRanch).Methods("PUT")

	router.HandleFunc("/ranches/disabled", common.GetDisabledRanches).Methods("GET")
}
