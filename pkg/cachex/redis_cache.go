package cachex

import (
	"context"
	"encoding/json"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
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
			logx.Errorf("[cache Put] %#v", err)
			return err
		}
		value = string(bytes)
	}
	logx.Info("[cache Put] ", value)
	err := r.client.SetexCtx(ctx, key, value, int(expire.Seconds()))
	if err != nil {
		logx.Error("[cache Put] ", err)
	}
	return err
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	s, err := r.client.GetCtx(ctx, key)
	if err != nil {
		return "", err
	}
	if len(s) == 0 {
		return "", redis.Nil
	}
	return s, nil
}
