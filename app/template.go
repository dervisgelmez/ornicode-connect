package app

import (
	"github.com/gin-gonic/gin"
	parameters "ornicode-connect/config"
)

func appTemplateEngine(r *gin.Engine) {
	r.LoadHTMLGlob("./src/" + parameters.TemplateFolder + "/*")
}
