package controllers

import "github.com/jinzhu/gorm"

var db *gorm.DB

func SetUpDBConnection(conn *gorm.DB) {
	db = conn
}
