package config

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

// InitTemplate initialize server
func InitTemplate(r *gin.Engine) *gin.Engine {
	r.SetHTMLTemplate(template.Must(
		template.ParseFiles(
			"templates/shared/head.tmpl",
			"templates/shared/header.tmpl",
			"templates/shared/help-banner.tmpl",
			"templates/main.tmpl",
		)))
	return r
}
