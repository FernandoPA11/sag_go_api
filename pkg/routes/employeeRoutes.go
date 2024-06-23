package routes

import (
	"SAG_GO_API/pkg/handlers/employee"

	"github.com/gorilla/mux"
)

func EmployeeRouteHandlers(router *mux.Router) {

	router.HandleFunc("/employees", employee.GetEmployees).Methods("GET")

	router.HandleFunc("/employees/{id}", employee.GetEmployeeByID).Methods("GET")

	router.HandleFunc("/employees", employee.AddEmployee).Methods("POST")

	router.HandleFunc("/employees/{id}", employee.UpdateEmployee).Methods("PUT")

	router.HandleFunc("/employees/{id}", employee.DeleteEmployee).Methods("DELETE")

	router.HandleFunc("/employees/{id}/disable", employee.DisableEmployee).Methods("PUT")

	router.HandleFunc("/employees/{id}/enable", employee.EnableEmployee).Methods("PUT")
}
