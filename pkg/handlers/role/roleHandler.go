package role

import (
	"SAG_GO_API/core/db"
	"SAG_GO_API/pkg/resources"
	"SAG_GO_API/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRoles(w http.ResponseWriter, r *http.Request) {
	var roles []resources.Role
	db.DB.Preload("Permissions").Find(&roles)
	if len(roles) == 0 {
		utils.Respond(w, http.StatusNotFound, "No roles found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Roles found", roles)
}

func GetRoleByID(w http.ResponseWriter, r *http.Request) {
	roleID := mux.Vars(r)["id"]
	var role resources.Role
	db.DB.First(&role, roleID)
	if role.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Role not found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Role found", role)
}

func AddRole(w http.ResponseWriter, r *http.Request) {
	var role resources.Role
	json.NewDecoder(r.Body).Decode(&role)

	createdRole := db.DB.Create(&role)
	if createdRole.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error creating role: "+createdRole.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusCreated, "Role created: "+role.Name, role)
}

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	roleID := mux.Vars(r)["id"]
	var role resources.Role
	db.DB.Unscoped().First(&role, roleID)
	if role.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Role not found", nil)
		return
	}
	json.NewDecoder(r.Body).Decode(&role)

	updatedRole := db.DB.Save(&role)
	if updatedRole.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error updating role: "+updatedRole.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Role updated: "+role.Name, role)
}

func DeleteRole(w http.ResponseWriter, r *http.Request) {
	roleID := mux.Vars(r)["id"]
	var role resources.Role
	db.DB.Unscoped().First(&role, roleID)
	if role.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Role not found", nil)
		return
	}
	db.DB.Delete(&role)
	utils.Respond(w, http.StatusOK, "Role deleted: "+role.Name, nil)
}

func DisableRole(w http.ResponseWriter, r *http.Request) {
	roleID := mux.Vars(r)["id"]
	var role resources.Role
	db.DB.First(&role, roleID)
	if role.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Role not found", nil)
		return
	}
	db.DB.Delete(&role)
	utils.Respond(w, http.StatusOK, "Role disabled: "+role.Name, nil)
}

func EnableRole(w http.ResponseWriter, r *http.Request) {
	roleID := mux.Vars(r)["id"]
	var role resources.Role
	db.DB.Unscoped().First(&role, roleID)
	if role.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Role not found", nil)
		return
	}
	db.DB.Model(&role).Update("deleted_at", nil)
	utils.Respond(w, http.StatusOK, "Role enabled: "+role.Name, nil)
}

func AddPermissionToRole(w http.ResponseWriter, r *http.Request) {
	roleID := mux.Vars(r)["id"]
	permissionID := mux.Vars(r)["permissionID"]
	var role resources.Role
	db.DB.First(&role, roleID)
	if role.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Role not found", nil)
		return
	}
	var permission resources.Permission
	db.DB.First(&permission, permissionID)
	if permission.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Permission not found", nil)
		return
	}
	db.DB.Model(&role).Association("Permissions").Append(&permission)
	utils.Respond(w, http.StatusOK, "Permission added to role", role)
}

func AddPermissionsToRole(w http.ResponseWriter, r *http.Request) {
	roleID := mux.Vars(r)["id"]
	var permissions []resources.Permission
	json.NewDecoder(r.Body).Decode(&permissions)
	var role resources.Role
	db.DB.First(&role, roleID)
	if role.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Role not found", nil)
		return
	}
	for _, permission := range permissions {
		db.DB.Model(&role).Association("Permissions").Append(&permission)
	}
	utils.Respond(w, http.StatusOK, "Permissions added to role", role)
}

func AddRoleWithPermissions(w http.ResponseWriter, r *http.Request) {
	// Define la estructura esperada en la solicitud
	var roleData resources.RoleWithPermissionIDs

	// Decodifica el JSON de la solicitud
	if err := json.NewDecoder(r.Body).Decode(&roleData); err != nil {
		utils.Respond(w, http.StatusBadRequest, "Invalid request payload", nil)
		return
	}

	// Crear el rol

	role := resources.Role{
		Name: roleData.Name,
	}

	// Iniciar una transacción
	tx := db.DB.Begin()
	if err := tx.Error; err != nil {
		utils.Respond(w, http.StatusInternalServerError, "Failed to start transaction", nil)
		return
	}

	if err := tx.Create(&role).Error; err != nil {
		tx.Rollback()
		utils.Respond(w, http.StatusInternalServerError, "Failed to create role: "+err.Error(), nil)
		return
	}

	// Verificar y asociar los permisos existentes al rol
	for _, permissionID := range roleData.PermissionIDs {
		var permission resources.Permission
		if err := tx.First(&permission, permissionID).Error; err != nil {
			tx.Rollback()
			utils.Respond(w, http.StatusNotFound, "Permission not found: "+err.Error(), nil)
			return
		}

		// Crear la relación en la tabla `role_permissions`
		rolePermission := resources.RolePermission{
			RoleID:       int(role.ID),
			PermissionID: int(permission.ID),
		}
		if err := tx.Create(&rolePermission).Error; err != nil {
			tx.Rollback()
			utils.Respond(w, http.StatusInternalServerError, "Failed to associate permission with role: "+err.Error(), nil)
			return
		}
	}

	// Commit de la transacción
	if err := tx.Commit().Error; err != nil {
		utils.Respond(w, http.StatusInternalServerError, "Failed to commit transaction: "+err.Error(), nil)
		return
	}

	// Responder con el rol creado
	utils.Respond(w, http.StatusOK, "Role created with permissions", role)
}

func GetRolesDisabled(w http.ResponseWriter, r *http.Request) {
	var roles []resources.Role
	db.DB.Unscoped().Find(&roles).Where("deleted_at IS NOT NULL")
	if len(roles) == 0 {
		utils.Respond(w, http.StatusNotFound, "No roles found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Roles found", roles)
}
