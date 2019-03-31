package models

import "time"

type MovieList struct {
	ID         uint64 // `gorm:"primary_key"`
	MovieID    uint64
	Title      string
	Rating     float32
	MovieYear  int16
	Url        string
	ImageLarge string
	Casts      string
	Directors  string
	Genres     string
	Summary    string
	Countries  string
	Viewed     bool
	Star       bool
	UpdateTime time.Time
	MovieType  string
}
