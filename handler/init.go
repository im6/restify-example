package handler

import (
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func registerHandler(r *gin.Engine) {
	// 404 Handler.
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "not found (404)")
	})

	// view handler
	r.GET("/", latestPage)
	r.GET("/latest", latestPage)
	r.GET("/popular", popularPage)
	r.GET("/color/:id", oneColorPage)
	r.GET("/new", createPage)
	r.GET("/signin", signInPage)

	// api handler
	r.POST("/create", handleCreateColor)
	r.POST("/like/:id", handleLikeColor)
	r.DELETE("/like/:id", handleUnlikeColor)
}

func createRender() multitemplate.Renderer {
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
	r.AddFromFiles(
		"create",
		"templates/create.tmpl",
		"templates/shared/head.tmpl",
		"templates/shared/header.tmpl",
	)
	r.AddFromFiles(
		"signin",
		"templates/signin.tmpl",
		"templates/shared/head.tmpl",
		"templates/shared/header.tmpl",
	)
	r.AddFromFiles(
		"error",
		"templates/error.tmpl",
	)
	return r
}

func Initialize(r *gin.Engine) *gin.Engine {
	r.Use(gin.Recovery())
	r.HTMLRender = createRender()
	registerHandler(r)
	return r
}
