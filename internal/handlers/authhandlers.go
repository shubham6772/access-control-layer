package handlers

import (
	"log"
	"net/http"

	"access-control-layer/internal/services"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Validate(c *gin.Context) {

	token, err := c.Cookie("playback_token")

	if err != nil{
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Printf("Token: %v", token)
	
	claims, err := h.authService.Validate(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Header("X-User-Id", claims.UserID)
	c.Header("X-Video-Id", claims.VideoID)
	c.Header("X-Session-Id", claims.SessionID)

	c.Status(http.StatusOK)
}
