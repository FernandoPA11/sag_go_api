package routes

import (
	"SAG_GO_API/pkg/handlers/permission"

	"github.com/gorilla/mux"
)

func PermissionRouteHandlers(router *mux.Router) {

	router.HandleFunc("/permissions", permission.GetPermissions).Methods("GET")

	router.HandleFunc("/permissions/{id}", permission.GetPermissionByID).Methods("GET")

	router.HandleFunc("/permissions", permission.AddPermission).Methods("POST")
}
