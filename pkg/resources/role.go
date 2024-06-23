package resources

import "gorm.io/gorm"

type Role struct {
	gorm.Model

	Name        string       `gorm:"unique;not null;default:null"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type RoleWithPermissionIDs struct {
	Name          string `json:"name"`
	PermissionIDs []int  `json:"permission_ids"`
}
