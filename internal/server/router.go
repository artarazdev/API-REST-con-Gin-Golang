package server

import (
	"github.com/artarazdev/API-REST-con-Gin-Golang/internal/health"
	"github.com/artarazdev/API-REST-con-Gin-Golang/internal/config"

	"github.com/gin-gonic/gin"
)

func NewRouter(cfg config.Config) *gin.Engine {
	router := gin.Default()

	router.GET("/health", health.Handler)

	return router
}
