package main

import (
	"github.com/gin-gonic/gin"
	"github.com/im6/vp3/config"
	"github.com/im6/vp3/handler"
)

func init() {
	config.LoadConfiguration()
	config.InitMySQLConn()
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"}) // reverse proxy ip address
	handler.Initialize(r)
	r.Run(":" + config.GetPort())
}
