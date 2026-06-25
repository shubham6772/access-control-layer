package redis

import (
	"access-control-layer/internal/config"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	Rdb *redis.Client
	Ctx context.Context
}

func NewClient(redisConf *config.RedisConfig) *Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         redisConf.Addr,
		Password:     redisConf.Password,
		DB:           redisConf.DB,
		PoolSize:     redisConf.PoolSize,
		MinIdleConns: redisConf.MinIdleConns,
		DialTimeout:  redisConf.DialTimeout,
		ReadTimeout:  redisConf.ReadTimeout,
		WriteTimeout: redisConf.WriteTimeout,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		panic("redis connection failed: " + err.Error())
	}

	return &Client{
		Rdb: rdb,
		Ctx: context.Background(),
	}
}

func (c *Client) Get(ctx context.Context, key string) (string, error) {
	return c.Rdb.Get(ctx, key).Result()
}

func (c *Client) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	return c.Rdb.Set(ctx, key, value, ttl).Err()
}

func (c *Client) Ping(ctx context.Context) error {
	return c.Rdb.Ping(ctx).Err()
}
