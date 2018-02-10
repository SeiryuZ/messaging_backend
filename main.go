package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

func main() {
	log.Println("Starting Server .....")

	log.Println("Initializing DB .....")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	log.Println("Migrate DB .....")
	db.AutoMigrate(&User{})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
