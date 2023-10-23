package app

import (
	"github.com/gin-gonic/gin"
)

func createApp() *gin.Engine {
	app := gin.New()
	app.Use(gin.Logger())

	return app
}

func Kernel() {
	engine := createApp()
	gin.SetMode(gin.DebugMode)

	appRouterEngine(engine)
	appTemplateEngine(engine)

	engine.Run()
}
