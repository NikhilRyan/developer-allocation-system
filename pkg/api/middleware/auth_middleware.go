package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "developer-allocation-system/pkg/services"
    "developer-allocation-system/pkg/utils"
)

// AuthMiddleware checks for valid JWT tokens.
func AuthMiddleware(authService services.AuthService) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            utils.RespondWithError(c, http.StatusUnauthorized, "Authorization header missing")
            c.Abort()
            return
        }

        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || parts[0] != "Bearer" {
            utils.RespondWithError(c, http.StatusUnauthorized, "Invalid authorization header format")
            c.Abort()
            return
        }

        token := parts[1]
        userID, err := authService.VerifyToken(token)
        if err != nil {
            utils.RespondWithError(c, http.StatusUnauthorized, "Invalid or expired token")
            c.Abort()
            return
        }

        // Set user ID in context for later use
        c.Set("userID", userID)
        c.Next()
    }
}
