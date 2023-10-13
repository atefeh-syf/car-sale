package cache

import (
	"fmt"
	"time"

	"github.com/atefeh-syf/car-sale/config"
	"github.com/go-redis/redis/v7"
)

var redisClient *redis.Client 

func InItRedis (cfg *config.Config) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:  fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB : 0,
		DialTimeout: cfg.Redis.DialTimeout * time.Second,
		ReadTimeout: cfg.Redis.ReadTimeout,
		WriteTimeout: cfg.Redis.WriteTimeout,
		PoolSize: cfg.Redis.PoolSize,
		PoolTimeout: cfg.Redis.PoolTimeout,
		IdleTimeout: 500 * time.Millisecond,
		IdleCheckFrequency: cfg.Redis.IdleCheckFrequency * time.Millisecond,	
	})
}

func GetRedis() * redis.Client {
	return redisClient
}

func CloseRedis() {
	redisClient.Close()
}