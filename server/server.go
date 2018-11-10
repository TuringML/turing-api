package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/turing-ml/turing-api/middleware"
	"github.com/turing-ml/turing-api/router"
)

func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.RedirectTrailingSlash = true

	r.Use(middleware.SetDB(db))

	router.Initialize(r)
	return r
}
