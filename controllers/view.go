package controllers

import (
	"encoding/json"
	"fmt"
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

	err := store.RedisDB.Set("helloworldkey", "hello123", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := store.RedisDB.Get("helloworldkey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	colorJSON, err := json.Marshal(colors)
	if err != nil {
		panic(err)
	}
	cxt.HTML(http.StatusOK, "main", gin.H{
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
	cxt.HTML(http.StatusOK, "main", gin.H{
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
	cxt.HTML(http.StatusOK, "one-color", gin.H{
		"data":      template.JS(colorJSON),
		"assetName": "bundle3",
		"version":   version,
	})
}

// CreatePage view
func CreatePage(cxt *gin.Context) {
	cxt.HTML(http.StatusOK, "create", gin.H{
		"assetName":    "bundle2",
		"version":      version,
		"defaultValue": cxt.Query("c"),
	})
}
