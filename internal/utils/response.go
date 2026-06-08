package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	RequestID string      `json:"request_id"`
	Data      interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {

	rid, _ := c.Get("X-Request-ID")

	c.JSON(http.StatusOK, Response{
		Code:      0,
		Message:   "ok",
		RequestID: rid.(string),
		Data:      data,
	})
}

func Error(c *gin.Context, code int, msg string) {

	rid, _ := c.Get("X-Request-ID")

	c.JSON(http.StatusOK, Response{
		Code:      code,
		Message:   msg,
		RequestID: rid.(string),
	})
}
