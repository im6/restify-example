package main

import (
	"github.com/gin-gonic/gin"
	"github.com/im6/vp3/config"
)

const port = ":3000"

func main() {
	r := gin.Default()

	config.InitTemplate(r)
	config.InitRoute(r)

	r.Run(port)
}
