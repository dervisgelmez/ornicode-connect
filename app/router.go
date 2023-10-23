package app

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"ornicode-connect/router"
)

func appRouterEngine(engine *gin.Engine) {
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.ListenApiRouter(engine)
	router.ListenWebRouter(engine)
}
