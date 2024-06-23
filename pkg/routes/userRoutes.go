package routes

import (
	"SAG_GO_API/pkg/handlers/user"

	"github.com/gorilla/mux"
)

func UserRouteHandlers(router *mux.Router) {

	router.HandleFunc("/users", user.GetUsers).Methods("GET")

	router.HandleFunc("/users/{id}", user.GetUserByID).Methods("GET")

	router.HandleFunc("/users", user.AddUser).Methods("POST")

	router.HandleFunc("/users/{id}", user.UpdateUser).Methods("PUT")

	router.HandleFunc("/users/{id}", user.DeleteUser).Methods("DELETE")

	router.HandleFunc("/users/{id}/disable", user.DisableUser).Methods("PUT")

	router.HandleFunc("/users/{id}/enable", user.EnableUser).Methods("PUT")
}
