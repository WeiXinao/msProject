package svc

import (
	"github.com/WeiXinao/msProject/pkg/cachex"
	"github.com/WeiXinao/msProject/user/internal/config"
	"github.com/WeiXinao/msProject/user/internal/repo"
	"github.com/WeiXinao/msProject/user/internal/repo/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config   config.Config
	UserRepo repo.UserRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	rcli := redis.MustNewRedis(c.RedisConfig)
	ca := cachex.NewRedisCache(rcli)

	userCache := &cache.UserCache{}
	userRepo := repo.NewUserRepo(ca, userCache)
	return &ServiceContext{
		UserRepo: userRepo,
		Config:   c,
	}
}
