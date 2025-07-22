package cache

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

func (r *RedisClientWrapper) SetNonce(walletAddress, nonce string, ttl time.Duration) error {
	key := fmt.Sprintf("nonce:%s", strings.ToLower(walletAddress))
	return r.Client.Set(context.Background(), key, nonce, ttl).Err()
}

func (r *RedisClientWrapper) GetNonce(walletAddress string) (string, error) {
	key := fmt.Sprintf("nonce:%s", strings.ToLower(walletAddress))
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
