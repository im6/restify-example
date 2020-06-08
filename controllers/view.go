package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/im6/vp3/models"
	"github.com/im6/vp3/store"
)

// LatestPage view
func LatestPage(cxt *gin.Context) {
	colors := []models.Color{}
	store.DB.Table("colorpk_color").Order("createdate desc").Where("display = 0").Find(&colors)
	colorJSON, err := json.Marshal(colors)
	if err != nil {
		panic(err)
	}
	cxt.HTML(http.StatusOK, "main.tmpl", gin.H{
		"data": template.JS(colorJSON),
	})
}

// PopularPage view
func PopularPage(cxt *gin.Context) {
	colors := []models.Color{}
	store.DB.Table("colorpk_color").Order("`like` desc").Where("display = 0").Find(&colors)
	colorJSON, err := json.Marshal(colors)
	if err != nil {
		panic(err)
	}
	cxt.HTML(http.StatusOK, "main.tmpl", gin.H{
		"data": template.JS(colorJSON),
	})
}

// OneColorPage view
func OneColorPage(cxt *gin.Context) {
	id := cxt.Param("id")
	cxt.JSON(200, gin.H{
		"message": "one color: " + id,
	})
}
