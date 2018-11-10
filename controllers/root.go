package controllers

import (
	"fmt"
	"github.com/turing-ml/turing-api/pkg/version"
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIEndpoints(c *gin.Context) {
	reqScheme := "http"

	if c.Request.TLS != nil {
		reqScheme = "https"
	}

	reqHost := c.Request.Host
	baseURL := fmt.Sprintf("%s://%s", reqScheme, reqHost)

	resources := map[string]string{
		"companies_url": baseURL + "/api/companies",
	}

	c.JSON(http.StatusOK, resources)
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, version.LongVersion())
}
