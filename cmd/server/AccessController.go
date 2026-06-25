package main

import (
	"access-control-layer/internal/config"
	"access-control-layer/internal/handlers"
	"access-control-layer/internal/routes"
	"access-control-layer/internal/services"
	"access-control-layer/internal/utils/jwt"
	"access-control-layer/internal/utils/logger"
	"access-control-layer/internal/utils/redis"
	"fmt"
	"log"

	// "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	logWriter, err := logger.Setup()
	if err != nil {
		panic(err)
	}
	defer logWriter.Close()

	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	globalConf := config.LoadGlobalConfig()
	redisConf := config.LoadRedisConfig()
	redisClient := redis.NewClient(redisConf)

	jwtConf := config.LoadJWTConfig()
	jwtClient := jwt.NewJwtClient(jwtConf)

	authService := services.NewAuthService(jwtClient, redisClient)
	actuatorService := services.NewHealthService(redisClient)

	authHandler := handlers.NewAuthHandler(authService)
	actuatorHandler := handlers.NewActuatorHandler(actuatorService)

	router := &routes.Router{
		AuthHandler:     authHandler,
		ActuatorHandler: actuatorHandler,
	}

	// gin.SetMode(globalConf.GIN_MODE)
	ginEngine := routes.Setup(router)
	PORT := fmt.Sprintf(":%v", globalConf.PORT)

	if err := ginEngine.Run(PORT); err != nil {
		panic(err)
	}
}
