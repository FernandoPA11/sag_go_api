package routes

import (
	"SAG_GO_API/pkg/handlers/role"

	"github.com/gorilla/mux"
)

func RoleRouteHandlers(router *mux.Router) {

	router.HandleFunc("/roles", role.GetRoles).Methods("GET")

	router.HandleFunc("/roles/{id}", role.GetRoleByID).Methods("GET")

	router.HandleFunc("/roles", role.AddRoleWithPermissions).Methods("POST")

	router.HandleFunc("/roles/{id}", role.UpdateRole).Methods("PUT")
}
