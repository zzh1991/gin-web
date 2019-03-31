package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUp(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
