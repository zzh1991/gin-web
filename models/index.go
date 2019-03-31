package models

import "github.com/jinzhu/gorm"

var db *gorm.DB

func SetUpDBConnection(conn *gorm.DB) {
	db = conn
}

// DBTables : Export tables list
var DBTables = []interface{}{
	&MovieList{},
}
