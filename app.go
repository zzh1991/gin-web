package main

import (
	"gin-web/api"
	"gin-web/config"
	"gin-web/routes"
	"gin-web/utils"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func configRouter() {
	routes.SetUp(router)

	//	api router
	apiRouter := router.Group("/api")
	api.SetUp(apiRouter)
}

func main() {
	config.LoadEnv()
	config.ConnectDatabase()

	configRouter()

	router.Run(":" + utils.DefaultGetEnv("PORT", "8080"))
}
