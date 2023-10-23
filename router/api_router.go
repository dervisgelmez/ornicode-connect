package router

import (
	"github.com/gin-gonic/gin"
	"ornicode-connect/middleware"
	controller "ornicode-connect/src/controller/api"
)

// ListenApiRouter Api router list
func ListenApiRouter(r *gin.Engine) {

	apiRequest := r.Group("/api")
	apiRequest.Use(middleware.CheckApiRequirements)
	{
		authRequest := apiRequest.Group("/auth")
		{
			authRequest.POST("/login", controller.Login)
			authRequest.POST("/register", controller.Register)
		}

		oauthRequest := apiRequest.Group("/oauth")
		{
			oauthRequest.POST("/token/email", controller.GetOAuthTokenWithEmail)
			oauthRequest.POST("/token/phone", controller.GetOAuthTokenWithPhoneNumber)
		}

		// This routes required authenticate
		apiRequest.Use(middleware.CheckApiAuthenticate)
		{
			apiRequest.GET("/validate", controller.Validate)
		}
	}
}
