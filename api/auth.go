package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/seiryuz/messaging_backend/models"
)

func Login(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)

	if err == nil {
		user, err = user.Login()
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"session_key": user.SessionKey})
		return
	}

	c.JSON(400, gin.H{"error": "Invalid login form"})
	return
}

func Register(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)

	if err == nil {
		err = user.Register()
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"message": "Registration successful",
		})
		return
	}

	c.JSON(400, gin.H{"error": "Invalid registration form"})
	return
}

func TestToken(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	message := fmt.Sprintf("Token for user %s OK", user.Username)
	c.JSON(200, gin.H{
		"message": message,
	})
}
