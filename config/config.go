package config

import (
	"fmt"
	"github.com/shakyapeiris/e-commerce-scraper/logging"
	"github.com/shakyapeiris/e-commerce-scraper/models"
	"os"
	"strconv"
)

var DB models.DBConfig

type Error struct {
	reason string
}

func (e Error) Error() string {
	return fmt.Sprintf("config error: %s", e.reason)
}

func init() {
	P, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		logging.Fatalln(err)
		logging.Fatalln(Error{reason: "DB_PORT must be an integer"})
		panic(Error{reason: "DB_PORT must be an integer"})
	}
	DB = models.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     int32(P),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}
}
