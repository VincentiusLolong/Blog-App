package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func AllEnv(s string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	value, errs := os.LookupEnv(s)
	if !errs {
		log.Fatal("Env Variable Empty or Not available")
	}
	return value
}
