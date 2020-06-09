package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/im6/vp3/models"
	"github.com/im6/vp3/store"
)

const version = "b12"

// LatestPage view
func LatestPage(cxt *gin.Context) {
	colors := []models.Color{}
	store.DB.Table("colorpk_color").Order("createdate desc").Where("display = 0").Find(&colors)
	colorJSON, err := json.Marshal(colors)
	if err != nil {
		panic(err)
	}
	cxt.HTML(http.StatusOK, "main.tmpl", gin.H{
		"data":      template.JS(colorJSON),
		"assetName": "bundle0",
		"version":   version,
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
		"data":      template.JS(colorJSON),
		"assetName": "bundle0",
		"version":   version,
	})
}

// OneColorPage view
func OneColorPage(cxt *gin.Context) {
	id := cxt.Param("id")
	color := models.Color{}
	store.DB.Table("colorpk_color").Where("id = " + id).Find(&color)
	colorJSON, err := json.Marshal(color)
	if err != nil {
		panic(err)
	}
	cxt.HTML(http.StatusOK, "one-color.tmpl", gin.H{
		"data":      template.JS(colorJSON),
		"assetName": "bundle3",
		"version":   version,
	})
}
