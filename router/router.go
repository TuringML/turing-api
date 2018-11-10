package router

import (
	"github.com/gin-gonic/gin"
	"github.com/turing-ml/turing-api/controllers"
)

func Initialize(r *gin.Engine) {
	r.GET("/", controllers.APIEndpoints)
	r.GET("/ping", controllers.Ping)
}
