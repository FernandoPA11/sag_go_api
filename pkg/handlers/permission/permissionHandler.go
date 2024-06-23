package permission

import (
	"SAG_GO_API/core/db"
	"SAG_GO_API/pkg/resources"
	"SAG_GO_API/pkg/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPermissions(w http.ResponseWriter, r *http.Request) {
	var permissions []resources.Permission
	db.DB.Find(&permissions)
	if len(permissions) == 0 {
		utils.Respond(w, http.StatusNotFound, "No permissions found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Permissions retrieved", permissions)
}

func GetPermissionByID(w http.ResponseWriter, r *http.Request) {
	permissionID := r.URL.Query().Get("id")
	var permission resources.Permission
	db.DB.First(&permission, permissionID)
	if permission.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Permission not found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Permission found", permission)
}

func AddPermission(w http.ResponseWriter, r *http.Request) {
	var permission resources.Permission
	json.NewDecoder(r.Body).Decode(&permission)

	createdPermission := db.DB.Create(&permission)
	if createdPermission.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error creating permission: "+createdPermission.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusCreated, "Permission created: "+*permission.Name, permission)
}

func UpdatePermission(w http.ResponseWriter, r *http.Request) {
	permissionID := r.URL.Query().Get("id")
	var permission resources.Permission
	db.DB.Unscoped().First(&permission, permissionID)
	if permission.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Permission not found", nil)
		return
	}
	json.NewDecoder(r.Body).Decode(&permission)

	updatedPermission := db.DB.Save(&permission)
	if updatedPermission.Error != nil {
		utils.Respond(w, http.StatusBadRequest, "Error updating permission: "+updatedPermission.Error.Error(), nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Permission updated", permission)
}

func DeletePermission(w http.ResponseWriter, r *http.Request) {
	permissionID := mux.Vars(r)["id"]
	var permission resources.Permission
	db.DB.First(&permission, permissionID)
	if permission.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Permission not found", nil)
		return
	}
	db.DB.Unscoped().Delete(&permission)
	utils.Respond(w, http.StatusOK, "Permission deleted", nil)
}

func DisablePermission(w http.ResponseWriter, r *http.Request) {
	permissionID := mux.Vars(r)["id"]
	var permission resources.Permission
	db.DB.First(&permission, permissionID)
	if permission.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Permission not found", nil)
		return
	}
	db.DB.Delete(&permission)
	utils.Respond(w, http.StatusOK, "Permission disabled", nil)
}

func EnablePermission(w http.ResponseWriter, r *http.Request) {
	permissionID := mux.Vars(r)["id"]
	var permission resources.Permission
	db.DB.First(&permission, permissionID)
	if permission.ID == 0 {
		utils.Respond(w, http.StatusNotFound, "Permission not found", nil)
		return
	}
	db.DB.Model(&permission).Unscoped().Update("deleted_at", nil)
	utils.Respond(w, http.StatusOK, "Permission enabled", nil)
}

func GetPermissionsDisabled(w http.ResponseWriter, r *http.Request) {
	var permissions []resources.Permission
	db.DB.Unscoped().Find(&permissions).Where("deleted_at IS NOT NULL")
	if len(permissions) == 0 {
		utils.Respond(w, http.StatusNotFound, "No permissions found", nil)
		return
	}
	utils.Respond(w, http.StatusOK, "Permissions retrieved", permissions)
}
