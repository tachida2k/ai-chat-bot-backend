package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func (r *RedisClientWrapper) SetSession(sessionID string, walletAddress string, ttl time.Duration) error {
	key := fmt.Sprintf("session:%s", sessionID)
	return r.Client.Set(context.Background(), key, walletAddress, ttl).Err()
}

func (r *RedisClientWrapper) GetSession(sessionID string) (string, error) {
	key := fmt.Sprintf("session:%s", sessionID)
	val, err := r.Client.Get(context.Background(), key).Result()

	// Key does not exist
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisClientWrapper) GetWalletAddress(sessionID string) (string, error) {
	key := fmt.Sprintf("session:%s", sessionID)
	return r.Client.Get(context.Background(), key).Result()
}
