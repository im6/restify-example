package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/im6/vp3/models"
)

const version = "v0.0.9"

// LatestPage view
func LatestPage(ctx *gin.Context) {
	colors, err := models.GetColors("latest")
	colorJSON, err := json.Marshal(colors)
	if err != nil {
		panic(err)
	}
	ctx.HTML(http.StatusOK, "main", gin.H{
		"data":      template.JS(colorJSON),
		"assetName": "bundle0",
		"version":   version,
	})
}

// PopularPage view
func PopularPage(ctx *gin.Context) {
	colors, err := models.GetColors("popular")
	colorJSON, err := json.Marshal(colors)
	if err != nil {
		panic(err)
	}
	ctx.HTML(http.StatusOK, "main", gin.H{
		"data":      template.JS(colorJSON),
		"assetName": "bundle0",
		"version":   version,
	})
}

// OneColorPage view
func OneColorPage(ctx *gin.Context) {
	id := ctx.Param("id")
	color, err := models.GetOneColor(id)
	if err != nil {
		ctx.HTML(http.StatusNotFound, "error", gin.H{})
		return
	}
	colorJSON, err := json.Marshal(color)
	if err != nil {
		panic(err)
	}
	ctx.HTML(http.StatusOK, "one-color", gin.H{
		"data":      template.JS(colorJSON),
		"assetName": "bundle3",
		"version":   version,
	})
}

// CreatePage view
func CreatePage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "create", gin.H{
		"assetName":    "bundle2",
		"version":      version,
		"defaultValue": ctx.Query("c"),
	})
}

// SignInPage view
func SignInPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signin", gin.H{
		"assetName":    "bundle1",
		"version":      version,
	})
}
