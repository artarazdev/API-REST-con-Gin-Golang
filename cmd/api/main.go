package main

import (
	"log"

	"github.com/artarazdev/API-REST-con-Gin-Golang/internal/health"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/health", health.Handler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
