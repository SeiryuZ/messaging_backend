package models

import (
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func InitDB() {
	log.Println("Initializing DB .....")
	var err error
	db, err = gorm.Open("sqlite3", "messaging_backend.db")
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	log.Println("Migrate DB .....")
	db.AutoMigrate(&User{})
}
