package routes

import "github.com/gin-gonic/gin"

func RegisterAuthRoutes(api *gin.RouterGroup, router *Router) {
	api.POST("/validate", router.AuthHandler.Validate)
}
