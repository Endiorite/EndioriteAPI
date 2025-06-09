package main

import (
	"EndioriteAPI/config"
	"EndioriteAPI/middleware"
	"EndioriteAPI/routes"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()
	r.Use(middleware.CustomRateLimiterMiddleware())
	routes.SetupRoutes(r)
	port := config.GetEnv("PORT", "8080")
	r.Run(":" + port)
}
