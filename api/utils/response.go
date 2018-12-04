package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type respError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type resp struct {
	Data json.RawMessage `json:"data"`
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
	b, err := json.Marshal(payload)
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	b, err = json.Marshal(&resp{Data: b})
	if err != nil {
		ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	r := string(b)
	c.JSON(code, r)
}
