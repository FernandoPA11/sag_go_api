package resources

import "gorm.io/gorm"

// Permission represents a permission in the system.
type Permission struct {
	gorm.Model

	Name string `gorm:"unique;not null"`
}

type UserPermission struct {
	gorm.Model

	UserID       int `gorm:"not null"`
	PermissionID int `gorm:"not null"`

	User       User       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Permission Permission `gorm:"foreignKey:PermissionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type RolePermission struct {
	gorm.Model

	RoleID       int `gorm:"not null"`
	PermissionID int `gorm:"not null"`

	Role       Role       `gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Permission Permission `gorm:"foreignKey:PermissionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
