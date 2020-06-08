package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/im6/vp3/controllers"
)

// Load initialize server
func Load(r *gin.Engine) *gin.Engine {
	// Middlewares.
	r.Use(gin.Recovery())

	// 404 Handler.
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	r.LoadHTMLGlob("views/*.tmpl")

	// view handler
	r.GET("/", controllers.LatestPage)
	r.GET("/latest", controllers.LatestPage)
	r.GET("/popular", controllers.PopularPage)
	r.GET("/color/:id", controllers.OneColorPage)

	// api handler
	r.POST("/like/:id", controllers.LikeColor)
	r.DELETE("/like/:id", controllers.UnlikeColor)

	return r
}
