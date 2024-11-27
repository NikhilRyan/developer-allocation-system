package middleware

import (
    "time"

    "github.com/gin-gonic/gin"
    "developer-allocation-system/pkg/utils"
)

// LoggingMiddleware logs each incoming request and its response time.
func LoggingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        startTime := time.Now()

        // Process request
        c.Next()

        // Calculate latency
        latency := time.Since(startTime)

        // Get status code
        statusCode := c.Writer.Status()

        // Log details
        utils.GetLogger().Infof("%s %s %d %s", c.Request.Method, c.Request.URL.Path, statusCode, latency)
    }
}
