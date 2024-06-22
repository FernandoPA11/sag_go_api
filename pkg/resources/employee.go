package resources

import "gorm.io/gorm"

type Employee struct {
	gorm.Model

	FirstName string             `gorm:"not null"`
	LastName  string             `gorm:"not null"`
	Phone     string             `gorm:"not null"`
	Address   string             `gorm:"not null"`
	Curp      string             `gorm:"unique;not null"`
	Salary    float64            `gorm:"not null" json:"salary"`
	Documents []EmployeeDocument `json:"documents"`
}

type EmployeeDocument struct {
	gorm.Model

	Name       string `gorm:"unique;not null" json:"name"`
	File       string `gorm:"not null" json:"file"`
	Extension  string `gorm:"not null" json:"extension"`
	EmployeeID int    `gorm:"not null" json:"employee_id"`

	Employee Employee `gorm:"foreignKey:EmployeeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
