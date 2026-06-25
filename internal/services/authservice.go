package services

import (
	"access-control-layer/internal/utils/jwt"
	"access-control-layer/internal/utils/redis"
)

type authService struct {
	jwtClient   *jwt.JwtClient
	redisClient *redis.Client
}

type AuthService interface {
	Validate(token string) (bool, error)
}

func NewAuthService(jwtClient *jwt.JwtClient, redisClient *redis.Client) *authService {
	return &authService{
		jwtClient:   jwtClient,
		redisClient: redisClient,
	}
}

func (auth *authService) Validate(token string) (bool, error) {
	return false, nil
}
