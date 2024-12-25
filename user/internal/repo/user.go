package repo

import (
	"context"
	"github.com/WeiXinao/msProject/pkg/cachex"
	"github.com/WeiXinao/msProject/user/internal/repo/cache"
	"time"
)

type UserRepo interface {
	CacheCaptcha(ctx context.Context, mobile, captcha string, expire time.Duration) error
}

type userRepo struct {
	cache     cachex.Cache
	userCache *cache.UserCache
}

func (u *userRepo) CacheCaptcha(ctx context.Context, mobile, captcha string, expire time.Duration) error {
	return u.cache.Put(ctx, u.userCache.RegisterCaptchaKey(mobile), captcha, expire)
}

func NewUserRepo(cache cachex.Cache, userCache *cache.UserCache) *userRepo {
	return &userRepo{
		cache:     cache,
		userCache: userCache,
	}
}
