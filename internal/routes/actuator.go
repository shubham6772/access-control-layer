package routes

import "github.com/gin-gonic/gin"

func RegisterActuatorRoutes(api *gin.RouterGroup, router *Router) {
	api.GET("/info", router.ActuatorHandler.Info)
	api.GET("/ready", router.ActuatorHandler.Ready)
	api.GET("/health", router.ActuatorHandler.Health)
}
