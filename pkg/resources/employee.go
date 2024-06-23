package resources

import "gorm.io/gorm"

type Employee struct {
	gorm.Model

	FirstName string             `gorm:"not null;default:null" json:"first_name"`
	LastName  string             `gorm:"not null;default:null" json:"last_name"`
	Phone     string             `gorm:"not null;default:null" json:"phone"`
	Address   string             `json:"address"`
	Curp      string             `gorm:"unique;not null;required;default:null" json:"curp"`
	Salary    float64            `gorm:"not null;default:null" json:"salary"`
	Documents []EmployeeDocument `json:"documents"`
}

type EmployeeDocument struct {
	gorm.Model

	Name       string `gorm:"unique;not null" json:"name"`
	File       string `gorm:"not null;default:null" json:"file"`
	Extension  string `gorm:"not null;default:null" json:"extension"`
	EmployeeID int    `gorm:"not null" json:"employee_id"`

	Employee Employee `gorm:"foreignKey:EmployeeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
