package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/im6/vp3/models"
)

// LikeColor api
func LikeColor(cxt *gin.Context) {
	idStr := cxt.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		cxt.JSON(500, gin.H{
			"error": true,
		})
		return
	}
	
	if err := models.IncrementColorStar(id); err != nil {
		cxt.JSON(500, gin.H{
			"error": true,
		})
		return
	}
	cxt.JSON(200, gin.H{
		"error": false,
	})
}

// UnlikeColor api
func UnlikeColor(cxt *gin.Context) {
	cxt.JSON(200, gin.H{
		"error": false,
	})
}
