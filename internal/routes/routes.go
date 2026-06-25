package routes

import (
	"github.com/gin-gonic/gin"
)

func Setup(router *Router) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")

	RegisterAuthRoutes(api, router)

	actuatorGroup := r.Group("/api/v1/actuator")
	RegisterActuatorRoutes(actuatorGroup, router)

	return r
}
