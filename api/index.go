package api

import "github.com/gin-gonic/gin"

func SetUp(router *gin.RouterGroup) {
	setUpMovieRoutes(router.Group("/movie"))
}
