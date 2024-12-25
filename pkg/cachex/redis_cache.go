package cachex

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"time"
)

type RedisCache struct {
	client *redis.Redis
}

func NewRedisCache(client *redis.Redis) *RedisCache {
	return &RedisCache{
		client: client,
	}
}

func (r *RedisCache) Put(ctx context.Context, key string, val any, expire time.Duration) error {
	var value string
	switch v := val.(type) {
	case string:
		value = v
	case []byte:
		value = string(v)
	default:
		bytes, err := json.Marshal(val)
		if err != nil {
			return err
		}
		value = string(bytes)
	}
	return r.client.SetexCtx(ctx, key, value, int(expire.Seconds()))
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return r.client.GetCtx(ctx, key)
}
