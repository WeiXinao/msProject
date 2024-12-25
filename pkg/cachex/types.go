package cachex

import (
	"context"
	"time"
)

type Cache interface {
	Put(ctx context.Context, key string, val any, expire time.Duration) error
	Get(ctx context.Context, key string) (string, error)
}
