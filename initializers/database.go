package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql" // Correct: Using the GORM MySQL driver
	"gorm.io/gorm"
	"kurkool-uss/go-gin-gorm-crud/models"
)

// DB is the global database connection instance.
var DB *gorm.DB

// ConnectToDB initializes the database connection and runs auto-migrations.
func ConnectToDB() {
	var err error
	// Construct the MySQL database connection string (DSN) from environment variables
	// user:pass@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	//log.Println("Connecting to database with DSN:", dsn)
	// Open a connection to the database using the MySQL driver
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) // Correct: Using mysql.Open()
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	log.Println("Database connection successful.")

	// Auto-migrate the schema
	// This will create the 'books' table if it doesn't exist.
	err = DB.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrated successfully.")
}
