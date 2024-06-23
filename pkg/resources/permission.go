package resources

import "gorm.io/gorm"

// Permission represents a permission in the system.
type Permission struct {
	gorm.Model

	Name *string `gorm:"unique;not null"`
}

type RolePermission struct {
	RoleID       int `gorm:"not null;default:null"`
	PermissionID int `gorm:"not null;default:null"`

	Role       Role       `gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Permission Permission `gorm:"foreignKey:PermissionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
