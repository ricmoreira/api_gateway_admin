package middleware

import (
	"time"

	"gopkg.in/gin-contrib/cors.v1"
	"github.com/gin-gonic/gin"
)

// Cors returns a cors handler
func Cors() gin.HandlerFunc{
	return cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:4200", "http://127.0.0.1:4200", "http://localhost:8089", "http://localhost:8069", "http://localhost:8080", "http://localhost:8099"},
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Accept",
			"Accept-Encoding",
			"Accept-Language",
			"access-control-allow-origin",
			"Access-Control-Request-Headers",
			"Access-Control-Request-Method",
			"authorization",
			"Cache-Control",
			"Connection",
			"Content-Type",
			"Host",
			"If-Modified-Since",
			"Keep-Alive",
			"Key",
			"Origin",
			"Pragma",
			"User-Agent",
			"X-Custom-Header"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           48 * time.Hour,
	})
}
