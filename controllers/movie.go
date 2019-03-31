package controllers

import (
	"gin-web/config"
	"gin-web/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindAllMovies(ctx *gin.Context) {
	var movieList []models.MovieList
	config.GetDBConnection().Find(&movieList)
	ctx.JSON(http.StatusOK, movieList)
}

func FindAllMoviesByType(ctx *gin.Context) {
	var movieList []models.MovieList
	var movieType = ctx.Param("type")
	config.GetDBConnection().Where(&models.MovieList{MovieType: movieType}).Find(&movieList)
	ctx.JSON(http.StatusOK, movieList)
}
