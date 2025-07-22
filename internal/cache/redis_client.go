package cache

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/otosei-ai/otosei-ai-backend/internal/config"
	"github.com/redis/go-redis/v9"
)

type RedisClientWrapper struct {
	Client *redis.Client
}

var Rdb *redis.Client
var ctx = context.Background()

func InitRedis(cfg *config.Config) *RedisClientWrapper {
	dbNum := 0
	if cfg.RedisDB != "" {
		var err error
		dbNum, err = strconv.Atoi(cfg.RedisDB)
		if err != nil {
			log.Fatalf("Invalid RedisDB value: %v", err)
		}
	}

	Rdb := redis.NewClient(&redis.Options{
		Addr:         cfg.RedisAddr,
		Username:     cfg.RedisUsername,
		Password:     cfg.RedisPassword,
		DB:           dbNum,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		MinIdleConns: 2,
		PoolSize:     10,
		PoolTimeout:  4 * time.Second,
	})

	if err := Rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis")
	return &RedisClientWrapper{Client: Rdb}
}
