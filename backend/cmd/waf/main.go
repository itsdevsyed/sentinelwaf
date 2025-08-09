package main

import (
    "github.com/gin-gonic/gin"
    "sentinelwaf/backend/internal/config"
    "sentinelwaf/backend/internal/middleware"

)

func main() {
    cfg := config.LoadConfig()

    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })


    protected := r.Group("/waf")
    protected.Use(middleware.JWTAuthMiddleware(cfg.JWTSecret))
    protected.GET("/check", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Request allowed through WAF"})
    })

    r.Run(":" + cfg.Port)
}
