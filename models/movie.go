package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Artists     string `json:"artists"`
	Genres      string `json:"genres"`
	Watch       int    `json:"watch"`
	WatchURL    string `json:"watch_url"`
}
