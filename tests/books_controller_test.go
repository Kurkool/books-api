package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"kurkool-uss/go-gin-gorm-crud/controllers"
	"kurkool-uss/go-gin-gorm-crud/models"
)

// MockBookRepository is a mock implementation of the BookRepository interface.
type MockBookRepository struct {
	mock.Mock
}

// Implement the interface methods for the mock
func (m *MockBookRepository) Create(book *models.Book) error {
	args := m.Called(book)
	// When creating, we can assign an ID to simulate the database
	if book.ID == 0 {
		book.ID = 1
	}
	return args.Error(0)
}

func (m *MockBookRepository) FindAll() ([]models.Book, error) {
	args := m.Called()
	return args.Get(0).([]models.Book), args.Error(1)
}

func (m *MockBookRepository) FindByID(id uint) (*models.Book, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Book), args.Error(1)
}

func (m *MockBookRepository) Update(book *models.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockBookRepository) Delete(id uint) (int64, error) {
	args := m.Called(id)
	return args.Get(0).(int64), args.Error(1)
}

// TestCreateBook tests the book creation endpoint using a mock repository.
func TestCreateBook(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockRepo := new(MockBookRepository)
	bookController := controllers.BooksController{Repo: mockRepo}

	// Define the book to be created
	bookToCreate := &models.Book{Title: "The Phoenix Project", Author: "Gene Kim"}

	// Set up the mock expectation
	// We expect the 'Create' method to be called with our book object.
	// We tell it to return 'nil' for the error, simulating a successful creation.
	mockRepo.On("Create", mock.AnythingOfType("*models.Book")).Return(nil)

	// Setup the HTTP request
	router := gin.Default()
	router.POST("/api/v1/books", bookController.CreateBook)
	body, _ := json.Marshal(bookToCreate)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)
	mockRepo.AssertExpectations(t) // Verify that the mock's expectations were met
}

// TestGetAllBooks tests fetching all books using a mock.
func TestGetAllBooks(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockRepo := new(MockBookRepository)
	bookController := controllers.BooksController{Repo: mockRepo}

	// Define what the mock should return
	expectedBooks := []models.Book{{Title: "Test Book", Author: "Test Author"}}
	mockRepo.On("FindAll").Return(expectedBooks, nil)

	// Setup the HTTP request
	router := gin.Default()
	router.GET("/api/v1/books", bookController.GetAllBooks)
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/books", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Book")
	mockRepo.AssertExpectations(t)
}

// TestGetBookByID tests fetching a book by its ID.
func TestGetBookByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(MockBookRepository)
	bookController := controllers.BooksController{Repo: mockRepo}

	expectedBook := &models.Book{ID: 1, Title: "Go in Action", Author: "William Kennedy"}
	mockRepo.On("FindByID", uint(1)).Return(expectedBook, nil)

	router := gin.Default()
	router.GET("/api/v1/books/:id", bookController.GetBookByID)
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/books/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Go in Action")
	mockRepo.AssertExpectations(t)
}

// TestUpdateBook tests updating a book.
func TestUpdateBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(MockBookRepository)
	bookController := controllers.BooksController{Repo: mockRepo}

	// The book that exists before update
	existingBook := &models.Book{ID: 1, Title: "Old Title", Author: "Old Author"}
	updatedBook := &models.Book{ID: 1, Title: "Updated Title", Author: "Updated Author"}

	// Expect FindByID to be called to fetch the existing book
	mockRepo.On("FindByID", uint(1)).Return(existingBook, nil)
	// Expect Update to be called with the updated book
	mockRepo.On("Update", mock.AnythingOfType("*models.Book")).Return(nil)

	router := gin.Default()
	router.PUT("/api/v1/books/:id", bookController.UpdateBook)
	body, _ := json.Marshal(updatedBook)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/books/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

// TestDeleteBook tests deleting a book.
func TestDeleteBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(MockBookRepository)
	bookController := controllers.BooksController{Repo: mockRepo}

	mockRepo.On("Delete", uint(1)).Return(int64(1), nil)

	router := gin.Default()
	router.DELETE("/api/v1/books/:id", bookController.DeleteBook)
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/books/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}
