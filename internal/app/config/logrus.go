package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func NewLogrus() *logrus.Logger {
	//
	log := logrus.New()

	logLevelInt, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
	if err != nil {
		fmt.Println("logrus")
		log.Fatal(err)
	}

	log.SetLevel(logrus.Level(logLevelInt))
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}
