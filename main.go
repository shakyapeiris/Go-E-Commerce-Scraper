package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shakyapeiris/e-commerce-scraper/api"
	"github.com/shakyapeiris/e-commerce-scraper/logging"
	"io"
	"net/http"
)

func ping(c *gin.Context) {
	logging.Info("Ping called")
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

	gin.DefaultWriter = io.Discard
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(logging.LogRequestMiddleware())
	r.GET("/ping", ping)

	logging.Info("Starting application...")
	err := r.Run()
	if err != nil {
		logging.Error(err.Error())
	}
}
