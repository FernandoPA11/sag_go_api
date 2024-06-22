package utils

import "net/mail"

func IsEmail(email string) bool {
	emailAddress, err := mail.ParseAddress(email)
	return err == nil && emailAddress.Address == email
}

func IsValidPassword(password string) bool {
	return len(password) > 5
}
