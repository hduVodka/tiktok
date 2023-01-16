package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/config"
	"tiktok/controller"
	"tiktok/log"
	"tiktok/models"
)

func main() {
	log.Init()
	config.Init()
	models.Init()

	e := gin.Default()
	controller.InitRouter(e)
	// run message websocket server
	err := e.Run(":" + config.Conf.GetString("server.port"))
	if err != nil {
		log.Fatal(err)
		return
	}
}
