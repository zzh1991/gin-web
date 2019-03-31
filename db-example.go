package main

import (
	"fmt"
	"gin-web/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=mydb password=postgres sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	db.SingularTable(true)
	fmt.Println(db.HasTable(&models.MovieList{}))
	var movieList models.MovieList
	db.First(&movieList)
	fmt.Println(movieList)
}
