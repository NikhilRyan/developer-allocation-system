package utils

import (
    "github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, code int, message string) {
    c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func RespondWithSuccess(c *gin.Context, code int, data interface{}) {
    c.JSON(code, data)
}
