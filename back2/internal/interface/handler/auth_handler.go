package handler

import (
	"back2/internal/usecase"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	staffUseCase *usecase.StaffUseCase
}

func NewAuthHandler(staffUseCase *usecase.StaffUseCase) *AuthHandler {
	return &AuthHandler{staffUseCase: staffUseCase}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginInput struct {
		StaffID  string `json:"staff_id" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	staff, err := h.staffUseCase.Authenticate(loginInput.StaffID, loginInput.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// JWTトークンの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"staff_id": staff.StaffID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	secretKey := os.Getenv("JWT_SECRET_KEY") // 環境変数から秘密鍵を取得
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":    tokenString,
		"staff_id": staff.StaffID,
		"message":  "Logged in successfully",
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "ログアウトしました"})
}
