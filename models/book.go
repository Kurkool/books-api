package models

import "gorm.io/gorm"

// Book represents the model for a book in the database.
// It includes gorm.Model to get fields like ID, CreatedAt, UpdatedAt, and DeletedAt.
type Book struct {
	gorm.Model
	ID     int    `json:"id"` // ID is included for JSON serialization, but gorm.Model already has ID
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}
