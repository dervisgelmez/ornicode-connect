package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckApiRequirements(c *gin.Context) {
	contentType := c.GetHeader("Content-Type")
	if contentType != "application/json" {
		c.Status(http.StatusBadRequest)
	}
}
