package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/im6/vp3/models"
)

type NewColor struct {
	Color []string `json:"color"`
}

func handleLikeColor(ctx *gin.Context) {
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

func handleUnlikeColor(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"error": false,
	})
}

func handleCreateColor(ctx *gin.Context) {
	var newColor NewColor
	err := ctx.BindJSON(&newColor)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, gin.H{
		"error":  false,
		"origin": newColor,
	})
}
