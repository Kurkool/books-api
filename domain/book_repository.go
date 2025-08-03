package domain

import "kurkool-uss/go-gin-gorm-crud/models"

// BookRepository defines the interface for database operations for books.
// This allows us to use a real database in production and a mock database in tests.
type BookRepository interface {
	Create(book *models.Book) error
	FindAll() ([]models.Book, error)
	FindByID(id uint) (*models.Book, error)
	Update(book *models.Book) error
	Delete(id uint) (int64, error)
}
