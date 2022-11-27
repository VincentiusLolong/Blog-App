package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var err error = godotenv.Load()

func AllEnv(s string) string {
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(s)
}
