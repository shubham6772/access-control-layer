package services

import (
	actuatordto "access-control-layer/internal/dto/actuator"
	"access-control-layer/internal/utils/redis"
	"context"
)

type ActuatorService interface {
	Info() actuatordto.InfoResponse
	Ready() actuatordto.ReadyResponse
	Health() actuatordto.HealthResponse
}

type HealthService struct {
	redis *redis.Client
}

func NewHealthService(redis *redis.Client) *HealthService {
	return &HealthService{
		redis: redis,
	}
}

func (h *HealthService) Health() actuatordto.HealthResponse {
	return actuatordto.HealthResponse{
		Status: "UP",
	}
}

func (h *HealthService) Ready() actuatordto.ReadyResponse {
	resp := actuatordto.ReadyResponse{
		Status: "READY",
		Redis:  "UP",
	}

	ctx := context.Background()

	if h.redis != nil {
		if err := h.redis.Ping(ctx); err != nil {
			resp.Status = "NOT_READY"
			resp.Redis = "DOWN"
		}
	}

	return resp
}

func (h *HealthService) Info() actuatordto.InfoResponse {
	return actuatordto.InfoResponse{
		Service: "access-control-layer",
		Version: "1.0.0",
	}
}
