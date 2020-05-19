package main

import "github.com/gin-gonic/gin"

func HomePage(cxt *gin.Context) {
	cxt.JSON(200, gin.H{
		"message": "latest colors",
	})
}

func OneColorPage(cxt *gin.Context) {
	id := cxt.Param("id")
	cxt.JSON(200, gin.H{
		"message": "latest color: " + id,
	})
}

func main() {
	r := gin.Default()
	r.GET("/", HomePage)
	r.GET("/color/:id", OneColorPage)
	r.Run()
}
