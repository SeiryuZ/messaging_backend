package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seiryuz/messaging_backend/api"
	"github.com/seiryuz/messaging_backend/models"
)

func main() {
	log.Println("Starting Server .....")
	models.InitDB()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	api.SetupRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
