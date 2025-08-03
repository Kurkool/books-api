package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"kurkool-uss/go-gin-gorm-crud/controllers"
	"kurkool-uss/go-gin-gorm-crud/initializers"
	"kurkool-uss/go-gin-gorm-crud/repository"
)

// init runs before main() to initialize the application.
func init() {
	// Load environment variables from .env file (if it exists)
	initializers.LoadEnvVars()
	// Connect to the database and run migrations
	initializers.ConnectToDB()
}

// main is the entry point of the application.
func main() {
	// Create instances of the repository and controller
	bookRepo := repository.NewGormBookRepository(initializers.DB)
	bookController := controllers.BooksController{Repo: bookRepo}

	// Initialize Gin router
	r := gin.Default()

	// --- API Routes ---
	v1 := r.Group("/api/v1")
	{
		bookRoutes := v1.Group("/books")
		{
			// Note: We now call the methods on the controller instance
			bookRoutes.POST("/", bookController.CreateBook)
			bookRoutes.GET("/", bookController.GetAllBooks)
			bookRoutes.GET("/:id", bookController.GetBookByID)
			bookRoutes.PUT("/:id", bookController.UpdateBook)
			bookRoutes.DELETE("/:id", bookController.DeleteBook)
		}
	}

	// Start the server
	log.Println("Starting server on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
