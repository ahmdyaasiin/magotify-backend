package main

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/config"
	"os"
	"strconv"
)

func init() {
	config.NewENV()
}

func main() {
	//
	log := config.NewLogrus()
	db := config.NewSQLX(log)
	validator := config.NewValidator()
	fiber := config.NewFiber()

	config.App(&config.AppConfig{
		App:       fiber,
		DB:        db,
		Log:       log,
		Validator: validator,
	})

	appPort, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		fmt.Println("port")
		log.Fatal(err)
	}

	err = fiber.Listen(fmt.Sprintf(":%d", appPort))
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
