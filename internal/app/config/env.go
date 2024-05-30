package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func NewENV() {
	err := godotenv.Load()
	env := os.Getenv("ENV")

	if err != nil && env == "" {
		log.Fatalf("error loading .env file: %v", err)
	}
}
