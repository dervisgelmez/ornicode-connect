package middleware

import (
	"github.com/gin-gonic/gin"
)

func CheckApiAuthenticate(c *gin.Context) {
	c.Next()
}
