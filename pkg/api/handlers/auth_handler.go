package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "developer-allocation-system/pkg/models"
    "developer-allocation-system/pkg/services"
    "developer-allocation-system/pkg/utils"
)

// AuthHandler handles authentication-related HTTP requests.
type AuthHandler struct {
    AuthService services.AuthService
}

// NewAuthHandler creates a new AuthHandler.
func NewAuthHandler(authService services.AuthService) *AuthHandler {
    return &AuthHandler{
        AuthService: authService,
    }
}

// Register handles user registration.
func (h *AuthHandler) Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid input data")
        return
    }

    err := h.AuthService.Register(&user)
    if err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Failed to register user")
        return
    }
    utils.RespondWithSuccess(c, http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login handles user login and token generation.
func (h *AuthHandler) Login(c *gin.Context) {
    var credentials models.Credentials
    if err := c.ShouldBindJSON(&credentials); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid input data")
        return
    }

    token, err := h.AuthService.Login(credentials)
    if err != nil {
        utils.RespondWithError(c, http.StatusUnauthorized, "Invalid username or password")
        return
    }
    utils.RespondWithSuccess(c, http.StatusOK, gin.H{"token": token})
}

// RefreshToken handles token refresh requests.
func (h *AuthHandler) RefreshToken(c *gin.Context) {
    var request struct {
        RefreshToken string `json:"refresh_token" binding:"required"`
    }
    if err := c.ShouldBindJSON(&request); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Invalid input data")
        return
    }

    newToken, err := h.AuthService.RefreshToken(request.RefreshToken)
    if err != nil {
        utils.RespondWithError(c, http.StatusUnauthorized, "Invalid refresh token")
        return
    }
    utils.RespondWithSuccess(c, http.StatusOK, gin.H{"token": newToken})
}
