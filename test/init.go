package test

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/config"
	val "github.com/go-playground/validator/v10"
	fib "github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	l "log"
	"os"
)

var db *sqlx.DB
var validator *val.Validate
var fiber *fib.App

var testEmail = "raymondananda12@gmail.com"

func init() {
	err := os.Chdir("..")
	if err != nil {
		l.Fatalf("could not change directory: %v", err)
	}

	config.NewENV()
	log := config.NewLogrusTest()
	db = config.NewSQLX(log)
	validator = config.NewValidator()
	fiber = config.NewFiber()

	config.App(&config.AppConfig{
		App:       fiber,
		DB:        db,
		Log:       log,
		Validator: validator,
	})
}
