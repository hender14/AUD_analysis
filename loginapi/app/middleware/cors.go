package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
)

// CORS Middleware
func CORS() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Writer.Header().Set("Content-Type", "application/json")
    c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_ADDRESS"))
    c.Writer.Header().Set("Access-Control-Max-Age", "86400")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
    c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

    if c.Request.Method == "OPTIONS" {
      c.AbortWithStatus(204)
    } else {
      c.Next()
    }
  }
}