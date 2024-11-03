package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"net/http"
	"octopusbi.com/e-commerce-platform/api"
	"octopusbi.com/e-commerce-platform/logging"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, api.Response{
		Message:    "pong",
		Success:    true,
		StatusCode: http.StatusOK,
	})
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		logging.Error("Error loading .env file")
	}
	r := gin.Default()
	r.GET("/ping", ping)

	logging.Info("Starting application...")
	err := r.Run()
	if err != nil {
		logging.Error(err.Error())
	}
}
