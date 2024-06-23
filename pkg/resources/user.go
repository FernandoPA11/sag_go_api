package resources

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username   string `gorm:"unique;not null;default:null" json:"username"`
	Password   string `gorm:"not null" json:"password"`
	Email      string `gorm:"unique;not null" json:"email"`
	RoleID     int    `gorm:"not null" json:"role_id"`
	EmployeeID int    `gorm:"not null" json:"employee_id"`

	Role     Role     `gorm:"foreignKey:RoleID;references:ID"`
	Employee Employee `gorm:"foreignKey:EmployeeID;references:ID"`
}
