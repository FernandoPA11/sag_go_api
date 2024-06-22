package main

import (
	"SAG_GO_API/core/db"
	"SAG_GO_API/pkg/resources"
	"SAG_GO_API/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	db.DBconnect()
	migration()

	router := mux.NewRouter()

	AddRoutes(router)

	http.ListenAndServe(":3000", router)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func AddRoutes(router *mux.Router) {
	routes.UserRouteHandlers(router)
}

func migration() {
	db.DB.AutoMigrate(&resources.User{})
	db.DB.AutoMigrate(&resources.Employee{})
	db.DB.AutoMigrate(&resources.DocumentEmployee{})
	db.DB.AutoMigrate(&resources.Role{})
	db.DB.AutoMigrate(&resources.Permission{})
	db.DB.AutoMigrate(&resources.RolePermission{})
	db.DB.AutoMigrate(&resources.UserPermission{})
}
