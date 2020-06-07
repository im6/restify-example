package controllers

import "github.com/gin-gonic/gin"

// LikeColor api
func LikeColor(cxt *gin.Context) {
	id := cxt.Param("id")
	cxt.JSON(200, gin.H{
		"message": "like: " + id,
	})
}

// UnlikeColor api
func UnlikeColor(cxt *gin.Context) {
	id := cxt.Param("id")
	cxt.JSON(200, gin.H{
		"message": "unlike: " + id,
	})
}
