package api

import "github.com/gin-gonic/gin"

func SetUpRouter(router *gin.RouterGroup) {
	setUpMovieRoutes(router.Group("/movie"))
}
