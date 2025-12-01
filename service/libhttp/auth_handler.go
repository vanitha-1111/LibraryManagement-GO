package libhttp

import (
	"library/service/handler"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("SUPER_SECRET_KEY")

type AuthHTTPHandler struct {
	userService *handler.UserService
}

func NewAuthHTTPHandler(service *handler.UserService) *AuthHTTPHandler {
	return &AuthHTTPHandler{userService: service}
}

func (h *AuthHTTPHandler) Login(c *gin.Context) {
	var loginReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	user, err := h.userService.Authenticate(loginReq.Username, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}
	claims := jwt.MapClaims{
		"user_id":  user.UserId,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	// SUCCESS → return token
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
