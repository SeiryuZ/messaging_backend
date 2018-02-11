package api

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/seiryuz/messaging_backend/models"
)

func handleError(statusCode int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(statusCode, resp)
	c.AbortWithStatus(statusCode)
}

// Middleware that handle Token Authentication
// See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization
// Expecting a header with "Authorization: Token <credentials>"
// This middleware will set user object to the context if token is valid
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authorizationTokens := strings.Split(c.GetHeader("Authorization"), " ")

		if len(authorizationTokens) != 2 || authorizationTokens[0] != "Token" {
			handleError(401, "API token malformed", c)
			return
		}

		user, err := models.CheckToken(authorizationTokens[1])
		if err != nil {
			handleError(401, "Invalid API token", c)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func SetupRoutes(router *gin.Engine) {

	API := router.Group("/api")
	{
		// These should not be protected
		API.POST("auth/register", Register)
		API.POST("/auth/login", Login)

		// These are protected
		API.Use(TokenAuthMiddleware())
		API.POST("/auth/test-token", TestToken)
	}

}
