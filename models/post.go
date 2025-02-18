package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	UserID   uint   `json:"user_id" gorm:"not null"`
	IsEdited bool   `json:"is_edited" gorm:"default:false"`
}
