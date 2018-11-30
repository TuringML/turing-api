package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/turing-ml/turing-api/pkg/vault"
)

// Vault is a middleware to inject the Vault client's object to each handler
func Vault(v *vault.Vault) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("VAULT", v)
		c.Next()
	}
}
