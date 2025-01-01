package svc

import (
	"github.com/WeiXinao/msProject/pkg/cachex"
	"github.com/WeiXinao/msProject/pkg/encrypts"
	"github.com/WeiXinao/msProject/pkg/jwtx"
	"github.com/WeiXinao/msProject/user/internal/config"
	"github.com/WeiXinao/msProject/user/internal/repo"
	"github.com/WeiXinao/msProject/user/internal/repo/cache"
	"github.com/WeiXinao/msProject/user/internal/repo/dao"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"time"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config    config.Config
	UserRepo  repo.UserRepo
	Jwter     jwtx.Jwter
	Encrypter encrypts.Encrypter
}

func NewServiceContext(c config.Config) *ServiceContext {
	encrypter := encrypts.NewEncrypter(c.AESKey)
	jwter := InitJwter(c)
	rcli := redis.MustNewRedis(c.RedisConfig)
	ca := cachex.NewRedisCache(rcli)
	db := InitDB(c)

	userCache := &cache.UserCache{}
	userDao := dao.NewXormUserDao(db)
	userRepo := repo.NewUserRepo(userDao, ca, userCache)
	return &ServiceContext{
		UserRepo:  userRepo,
		Config:    c,
		Jwter:     jwter,
		Encrypter: encrypter,
	}
}

func InitJwter(c config.Config) jwtx.Jwter {
	aExp, err := time.ParseDuration(c.Jwt.AccessExp)
	if err != nil {
		panic(err)
	}
	rExp, err := time.ParseDuration(c.Jwt.AccessExp)
	if err != nil {
		panic(err)
	}
	return jwtx.NewJwtToken(c.Jwt.AtKey, c.Jwt.RtKey, aExp, rExp)
}

func InitDB(c config.Config) *xorm.Engine {
	mySQLConfig := c.MySQLConfig
	engine, err := xorm.NewEngine(mySQLConfig.DriverName, mySQLConfig.Dsn)
	if err != nil {
		panic(err)
	}
	err = engine.Sync(
		new(dao.Member),
		new(dao.Organization),
	)
	if err != nil {
		panic(err)
	}
	return engine
}
