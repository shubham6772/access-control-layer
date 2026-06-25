package handlers

import (
	"access-control-layer/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ActuatorHandler struct {
	actuatorService services.ActuatorService
}

func NewActuatorHandler(actuatorService services.ActuatorService) *ActuatorHandler {
	return &ActuatorHandler{
		actuatorService: actuatorService,
	}
}

func (h *ActuatorHandler) Info(c *gin.Context) {
	resp := h.actuatorService.Info()
	c.JSON(http.StatusOK, resp)
}

func (h *ActuatorHandler) Health(c *gin.Context) {
	resp := h.actuatorService.Health()
	c.JSON(http.StatusOK, resp)
}

func (h *ActuatorHandler) Ready(c *gin.Context) {
	resp := h.actuatorService.Ready()

	// Kubernetes expects 503 when NOT_READY
	if resp.Status != "READY" {
		c.JSON(http.StatusServiceUnavailable, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}
