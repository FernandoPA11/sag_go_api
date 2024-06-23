package db

import (
	"SAG_GO_API/core/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ConfigDB = config.Cfg.Database
var DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", ConfigDB.Host, ConfigDB.User, ConfigDB.Password, ConfigDB.DBName, ConfigDB.Port)
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
