package utils

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type respError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// ResponseError will return the error message
func ResponseError(c *gin.Context, status int, err error) {
	er := respError{
		Code:    status,
		Message: err.Error(),
	}
	c.JSON(status, er)
}

// Response will return a response with a specific object in it
func Response(c *gin.Context, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	c.JSON(code, response)
}
