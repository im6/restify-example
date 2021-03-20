package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/im6/vp3/models"
)

// LikeColor api
func LikeColor(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": true,
		})
		return
	}
	
	if err := models.IncrementColorStar(id); err != nil {
		ctx.JSON(500, gin.H{
			"error": true,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"error": false,
	})
}

// UnlikeColor api
func UnlikeColor(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"error": false,
	})
}

type NewColor struct {
	Color []string `json:"color"`
}

// Create Color
func CreateColor(ctx *gin.Context) {
	var newColor NewColor 
  err := ctx.BindJSON(&newColor)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, gin.H{
		"error": false,
		"origin": newColor,
	})
}
