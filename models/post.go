package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID     string  `gorm:"primaryKey"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Body   string  `json:"body"`
	Reply  []Reply `json:"replies"`
}
