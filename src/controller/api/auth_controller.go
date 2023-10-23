package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "ornicode-connect/docs"
	"ornicode-connect/src/services"
	types "ornicode-connect/src/types/request"
)

func Login(c *gin.Context) {
	var request types.LoginRequestType
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func Register(c *gin.Context) {
	var request types.RegisterRequestType
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func GetOAuthTokenWithEmail(c *gin.Context) {
	var request types.ValidateEmail2FAType
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func GetOAuthTokenWithPhoneNumber(c *gin.Context) {
	var request types.GetOAuthTokenWithPhoneNumberType
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenResponse := services.UserService{}.Generate2FATokenForPhoneNumber(request)

	c.JSON(http.StatusOK, tokenResponse)
}

func Validate(c *gin.Context) {
}
