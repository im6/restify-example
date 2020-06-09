package config

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles(
		"main",
		"templates/main.tmpl",
		"templates/shared/head.tmpl",
		"templates/shared/header.tmpl",
		"templates/shared/help-banner.tmpl",
	)
	r.AddFromFiles(
		"one-color",
		"templates/one-color.tmpl",
		"templates/shared/head.tmpl",
		"templates/shared/header.tmpl",
	)
	return r
}

// InitTemplate initialize server
func InitTemplate(r *gin.Engine) *gin.Engine {
	r.HTMLRender = createMyRender()
	return r
}
