package main

import (
	"github.com/gin-gonic/gin"
	"github.com/im6/vp3/config"
	"github.com/im6/vp3/store"
)

const port = ":3000"

func main() {
	r := gin.Default()

	config.InitTemplate(r)
	config.InitRoute(r)
	
	err := store.InitMySQLConn()
	if err != nil {
		panic("mysql connect error")
	}

	r.Run(port)
}
