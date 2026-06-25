package services

import (
	"access-control-layer/internal/utils/jwt"
	"access-control-layer/internal/utils/redis"
	"errors"
	"log"
)

type authService struct {
	jwtClient   *jwt.JwtClient
	redisClient *redis.Client
}

type AuthService interface {
	Validate(token string) (*jwt.Claims, error)
}

func NewAuthService(jwtClient *jwt.JwtClient, redisClient *redis.Client) *authService {
	return &authService{
		jwtClient:   jwtClient,
		redisClient: redisClient,
	}
}

func (auth *authService) Validate(token string) (*jwt.Claims, error) {

	claims, err := auth.jwtClient.Validate(token)
	if err != nil {
		return nil, err
	}

	log.Println("UserID:", claims.UserID)
	log.Println("VideoID:", claims.VideoID)
	log.Println("SessionID:", claims.SessionID)

	key := "session:" + claims.SessionID

	log.Println("Redis Key:", key)

	count, err := auth.redisClient.Rdb.Exists(
		auth.redisClient.Ctx,
		key,
	).Result()

	val, err := auth.redisClient.Rdb.Get(
		auth.redisClient.Ctx,
		"session:sess123",
	).Result()

	log.Println("VAL:", val)
	log.Println("ERR:", err)

	if err != nil {
		return nil, err
	}

	log.Println("Exists Count:", count)

	if count == 0 {
		return nil, errors.New("session not found")
	}

	return claims, nil
}
