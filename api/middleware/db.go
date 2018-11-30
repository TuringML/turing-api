package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// DB is a middleware to inject the DB object to each handler
func DB(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}
