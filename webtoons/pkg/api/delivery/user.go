package delivery

import (
    "net/http"
    "github.com/webtoons/pkg/domain"
    "github.com/webtoons/pkg/usecase"
    "github.com/gin-gonic/gin"
)

type AuthHandler struct {
    AuthUC usecase.AuthUseCase
}

func NewAuthHandler(authUC usecase.AuthUseCase) *AuthHandler {
    return &AuthHandler{
        AuthUC: authUC,
    }
}
// Register endpoint for user signup
func (h *AuthHandler) Register(c *gin.Context) {
    var user domain.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    err := h.AuthUC.RegisterUser(user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Login endpoint for user authentication
func (h *AuthHandler) Login(c *gin.Context) {
    var loginCredentials struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&loginCredentials); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    token, err := h.AuthUC.LoginUser(loginCredentials.Username, loginCredentials.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
