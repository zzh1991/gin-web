package api

import (
	"gin-web/controllers"
	"github.com/gin-gonic/gin"
)

func setUpMovieRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.FindAllMovies)
	router.GET("/:type", controllers.FindAllMoviesByType)
}
