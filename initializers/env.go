package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnvVars loads environment variables from a .env file.
// This is primarily for local development outside of Docker.
func LoadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		// This is not a fatal error because in production/Docker,
		// env vars will be set directly.
		log.Println("No .env file found, using environment variables from system")
	}
}
