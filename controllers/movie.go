package controllers

import (
	"gin-web/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindAllMovies(ctx *gin.Context) {
	var movieList []models.MovieList
	db.Find(&movieList)
	ctx.JSON(http.StatusOK, movieList)
}

func FindAllMoviesByType(ctx *gin.Context) {
	var movieList []models.MovieList
	var movieType = ctx.Param("type")
	//db.Where("movie_type = ?", movieType).Find(&movieList)
	db.Where(&models.MovieList{MovieType: movieType}).Find(&movieList)
	ctx.JSON(http.StatusOK, movieList)
}
