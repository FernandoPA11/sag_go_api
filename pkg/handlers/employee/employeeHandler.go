package employee

import (
	"SAG_GO_API/core/db"
	"SAG_GO_API/pkg/resources"
	"SAG_GO_API/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	var employees []resources.Employee
	db.DB.Find(&employees)
	if len(employees) == 0 {
		utils.Respond(w, http.StatusNotFound, "No employees found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Employees found", employees)
}

func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	employeeID := mux.Vars(r)["id"]
	var employee resources.Employee
	db.DB.First(&employee, employeeID)
	if employee.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Employee not found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Employee found", employee)
}

func GetEmployeesByRole(w http.ResponseWriter, r *http.Request) {
	roleID := mux.Vars(r)["roleID"]
	var employees []resources.Employee
	db.DB.Where("role_id = ?", roleID).Find(&employees)
	if len(employees) == 0 {
		utils.Respond(w, http.StatusNotFound, "No employees found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Employees found", employees)
}

func GetEmployeesByRanch(w http.ResponseWriter, r *http.Request) {
	ranchID := mux.Vars(r)["ranchID"]
	var employees []resources.Employee
	db.DB.Where("ranch_id = ?", ranchID).Find(&employees)
	if len(employees) == 0 {
		utils.Respond(w, http.StatusNotFound, "No employees found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Employees found", employees)
}

func GetEmployeesByCURP(w http.ResponseWriter, r *http.Request) {
	curp := mux.Vars(r)["curp"]
	var employee resources.Employee
	curpString := "%" + curp + "%"
	db.DB.Where("curp LIKE ?", curpString).Find(&employee)
	if employee.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Employee not found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Employee found", employee)
}

func GetEmployeesBySalary(w http.ResponseWriter, r *http.Request) {
	var salary salaryRange
	json.NewDecoder(r.Body).Decode(&salary)
	var employees []resources.Employee
	db.DB.Where("salary < ? AND salary > ?", salary.Max, salary.Min).Find(&employees)
	if len(employees) == 0 {
		utils.Respond(w, http.StatusNotFound, "No employees found on this salary range", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Employees found", employees)
}

func AddEmployee(w http.ResponseWriter, r *http.Request) {
	var employee resources.Employee
	json.NewDecoder(r.Body).Decode(&employee)

	createdEmployee := db.DB.Create(&employee)
	if createdEmployee.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error creating employee: "+createdEmployee.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusCreated, "Employee created: "+employee.FirstName, employee)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	employeeID := mux.Vars(r)["id"]
	var employee resources.Employee
	db.DB.Unscoped().First(&employee, employeeID)
	if employee.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Employee not found", nil)
		return
	}
	json.NewDecoder(r.Body).Decode(&employee)

	updatedEmployee := db.DB.Save(&employee)
	if updatedEmployee.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error updating employee: "+updatedEmployee.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Employee updated", employee)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	employeeID := mux.Vars(r)["id"]
	var employee resources.Employee
	db.DB.Unscoped().First(&employee, employeeID)
	if employee.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Employee not found", nil)
		return
	}
	db.DB.Unscoped().Delete(&employee)
	utils.Respond(w, http.StatusOK, "Employee deleted", nil)
}

func DisableEmployee(w http.ResponseWriter, r *http.Request) {
	employeeID := mux.Vars(r)["id"]
	var employee resources.Employee
	db.DB.First(&employee, employeeID)
	if employee.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Employee not found", nil)
		return
	}
	db.DB.Delete(&employee)
	utils.Respond(w, http.StatusOK, "Employee disabled", nil)
}

func EnableEmployee(w http.ResponseWriter, r *http.Request) {
	employeeID := mux.Vars(r)["id"]
	var employee resources.Employee
	db.DB.Unscoped().First(&employee, employeeID)
	if employee.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Employee not found", nil)
		return
	}
	db.DB.Model(&employee).Unscoped().Update("deleted_at", nil)
	utils.Respond(w, http.StatusOK, "Employee enabled", nil)
}

func GetEmployeesDisebled(w http.ResponseWriter, r *http.Request) {
	var employees []resources.Employee
	db.DB.Unscoped().Find(&employees).Where("deleted_at IS NOT NULL")
	if len(employees) == 0 {
		utils.Respond(w, http.StatusNotFound, "No employees found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Employees found", employees)
}

type salaryRange struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}
