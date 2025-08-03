package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"kurkool-uss/go-gin-gorm-crud/domain"
	"kurkool-uss/go-gin-gorm-crud/models"
)

// BooksController holds the repository dependency.
type BooksController struct {
	Repo domain.BookRepository
}

// CreateBook handles the creation of a new book.
func (ctrl *BooksController) CreateBook(c *gin.Context) {
	var body struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	book := models.Book{Title: body.Title, Author: body.Author}
	if err := ctrl.Repo.Create(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"book": book})
}

// GetAllBooks retrieves all books.
func (ctrl *BooksController) GetAllBooks(c *gin.Context) {
	books, err := ctrl.Repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"books": books})
}

// GetBookByID retrieves a single book by its ID.
func (ctrl *BooksController) GetBookByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	book, err := ctrl.Repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": book})
}

// UpdateBook updates an existing book's information.
func (ctrl *BooksController) UpdateBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// First, check if the book exists
	book, err := ctrl.Repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var body struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	// Update the book's fields
	book.Title = body.Title
	book.Author = body.Author

	if err := ctrl.Repo.Update(book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": book})
}

// DeleteBook removes a book.
func (ctrl *BooksController) DeleteBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	rowsAffected, err := ctrl.Repo.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
