package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func HomePage(cxt *gin.Context) {
	cxt.HTML(http.StatusOK, "main.html", nil)
}

func PopularPage(cxt *gin.Context) {
	cxt.JSON(200, gin.H{
		"message": "popular colors",
	})
}

func OneColorPage(cxt *gin.Context) {
	id := cxt.Param("id")
	cxt.JSON(200, gin.H{
		"message": "one color: " + id,
	})
}

func handleLikeColor(cxt *gin.Context) {
	id := cxt.Param("id")
	cxt.JSON(200, gin.H{
		"message": "like: " + id,
	})
}

func handleUnlikeColor(cxt *gin.Context) {
	id := cxt.Param("id")
	cxt.JSON(200, gin.H{
		"message": "unlike: " + id,
	})
}

var port = ":8080"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*.html")

	r.GET("/", HomePage)
	r.GET("/popular", PopularPage)
	r.GET("/color/:id", OneColorPage)
	r.POST("/like/:id", handleLikeColor)
	r.DELETE("/like/:id", handleUnlikeColor)

	log.Printf("start to listening: %s", port)
	r.Run(port)
}
