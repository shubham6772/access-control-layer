package config

import (
	"fmt"
	"os"
	"time"
)

type JwtConfig struct {
	SecretKey string
}

type RedisConfig struct {
	Addr         string
	Password     string
	DB           int
	PoolSize     int
	MinIdleConns int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type GlobalConfig struct {
	GIN_MODE string
	PORT     string
}

const (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
	REDIS_HOST     = "REDIS_HOST"
	REDIS_PASS     = "REDIS_PASSWORD"
	REDIS_DB       = "REDIS_DB"
	GIN_MODE       = "GIN_MODE"
	PORT           = "PORT"
)

func LoadGlobalConfig() *GlobalConfig {
	return &GlobalConfig{
		GIN_MODE: getENVValue(GIN_MODE),
		PORT:     getENVValue(PORT),
	}
}

func LoadJWTConfig() *JwtConfig {
	return &JwtConfig{
		SecretKey: getENVValue(JWT_SECRET_KEY),
	}
}

func LoadRedisConfig() *RedisConfig {
	return &RedisConfig{
		Addr:         getENVValue(REDIS_HOST),
		Password:     getENVValue(REDIS_PASS),
		DB:           0,
		PoolSize:     10,
		MinIdleConns: 2,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
}

func getENVValue(envKey string) string {
	envValue := os.Getenv(envKey)

	if envValue == "" {
		panic(fmt.Sprintf("missing required env: %s", envKey))
	}

	return envValue
}
