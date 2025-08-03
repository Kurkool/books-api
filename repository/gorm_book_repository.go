package repository

import (
	"gorm.io/gorm"
	"kurkool-uss/go-gin-gorm-crud/domain"
	"kurkool-uss/go-gin-gorm-crud/models"
)

// gormBookRepository is the GORM implementation of the BookRepository interface.
type gormBookRepository struct {
	db *gorm.DB
}

// NewGormBookRepository creates a new instance of the GORM book repository.
func NewGormBookRepository(db *gorm.DB) domain.BookRepository {
	return &gormBookRepository{db: db}
}

func (r *gormBookRepository) Create(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *gormBookRepository) FindAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *gormBookRepository) FindByID(id uint) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	return &book, err
}

func (r *gormBookRepository) Update(book *models.Book) error {
	return r.db.Save(book).Error
}

func (r *gormBookRepository) Delete(id uint) (int64, error) {
	result := r.db.Delete(&models.Book{}, id)
	return result.RowsAffected, result.Error
}
