package main

import (
	"SAG_GO_API/core/config"
	"SAG_GO_API/core/db"
	"SAG_GO_API/pkg/resources"
	"SAG_GO_API/pkg/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Load config yaml file
	config.LoadConfig()
	//Connect to database
	db.DBconnect()
	//run gorm migration
	migration()
	//Initialize router
	router := mux.NewRouter()
	//Add routes
	AddRoutes(router)
	//Set home route
	serverPort := config.Cfg.Server.Port
	http.ListenAndServe(fmt.Sprintf(":%d", serverPort), router)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func AddRoutes(router *mux.Router) {
	routes.UserRouteHandlers(router)
	routes.RoleRouteHandlers(router)
	routes.PermissionRouteHandlers(router)
	routes.EmployeeRouteHandlers(router)
	routes.RanchRouteHandlers(router)
	routes.HerdRouteHandlers(router)
	routes.BreedRouteHandlers(router)
	routes.CorralRouteHandlers(router)
	routes.SpecieRouteHandlers(router)
}

func migration() {
	db.DB.AutoMigrate(&resources.User{})
	db.DB.AutoMigrate(&resources.Employee{})
	db.DB.AutoMigrate(&resources.EmployeeDocument{})
	db.DB.AutoMigrate(&resources.Role{})
	db.DB.AutoMigrate(&resources.Permission{})
	db.DB.AutoMigrate(&resources.RolePermission{})
	db.DB.AutoMigrate(&resources.herd{})
	db.DB.AutoMigrate(&resources.Breed{})
	db.DB.AutoMigrate(&resources.Corral{})
	db.DB.AutoMigrate(&resources.Picture{})
}
