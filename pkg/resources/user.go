package resources

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID       int
	Username string
	Password string
	Email    string
}
