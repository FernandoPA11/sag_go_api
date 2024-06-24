package routes

import (
	"SAG_GO_API/pkg/handlers/herd"

	"github.com/gorilla/mux"
)

func HerdRouteHandlers(router *mux.Router) {

	router.HandleFunc("/herds", herd.GetHerd).Methods("GET")

	router.HandleFunc("/herds/{id}", herd.GetAnimalByID).Methods("GET")

	router.HandleFunc("/herds", herd.AddAnimal).Methods("POST")

	router.HandleFunc("/herds/{id}", herd.UpdateAnimal).Methods("PUT")

	router.HandleFunc("/herds/{id}", herd.DeleteAnimal).Methods("DELETE")

	router.HandleFunc("/herds/{id}/disable", herd.DisableAnimal).Methods("PUT")

	router.HandleFunc("/herds/{id}/enable", herd.EnableAnimal).Methods("PUT")

	router.HandleFunc("/herds/disabled", herd.GetDisabledAnimals).Methods("GET")
}
