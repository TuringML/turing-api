package middleware

import (
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
)

// AuthenticationSecret is a middleware to inject the Authentication Secret
func AuthenticationSecret(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("auth-secret", secret)
		c.Next()
	}
}

// Authentication is a middleware to validate the Authorization token in the request
func Authentication(secret string) gin.HandlerFunc {
	return jwt.Auth(secret)
}
