package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=fer password=fer123 dbname=sag port=5433 sslmode=disable"
var DB *gorm.DB

func DBconnect() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatalf("Failed to connect to database! %v", error)
	} else {
		println("Connected to database!")
	}
}
