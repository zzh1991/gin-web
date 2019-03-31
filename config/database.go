package config

import (
	"fmt"
	"gin-web/controllers"
	"gin-web/models"
	"gin-web/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func ConnectDatabase() {

	dbHost := utils.DefaultGetEnv("DB_HOST", "localhost")
	dbPort := utils.DefaultGetEnv("DB_PORT", "5432")
	dbName := utils.DefaultGetEnv("DB_NAME", "mydb")
	dbUser := utils.DefaultGetEnv("DB_USER", "postgres")
	dbPass := utils.DefaultGetEnv("DB_PASS", "postgres")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName, dbPass,
	)

	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connect database PostgreSQL successfully")
	}

	// Pass db connection to package controllers and models
	models.SetUpDBConnection(db)
	controllers.SetUpDBConnection(db)

	// Store this db connection for package config
	db.SingularTable(true)
	setUpDBConnection(db)
}

func setUpDBConnection(DB *gorm.DB) {
	db = DB
}

// GetDBConnection : get db connection from package config
func GetDBConnection() *gorm.DB {
	return db
}
