package svc

import (
	"github.com/WeiXinao/msProject/pkg/cachex"
	"github.com/WeiXinao/msProject/pkg/encrypts"
	"github.com/WeiXinao/msProject/pkg/jwtx"
	"github.com/WeiXinao/msProject/project/internal/config"
	"github.com/WeiXinao/msProject/project/internal/repo"
	"github.com/WeiXinao/msProject/project/internal/repo/dao"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"time"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config      config.Config
	Jwter       jwtx.Jwter
	Encrypter   encrypts.Encrypter
	ProjectRepo repo.ProjectRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	encrypter := encrypts.NewEncrypter(c.AESKey)
	jwter := InitJwter(c)
	rcli := redis.MustNewRedis(c.RedisConfig)
	ca := cachex.NewRedisCache(rcli)
	db := InitDB(c)

	projectDao, err := dao.NewProjectXormDao(db)
	if err != nil {
		panic(err)
	}
	projectRepo := repo.NewProjectRepo(ca, projectDao)
	return &ServiceContext{
		Config:      c,
		Jwter:       jwter,
		Encrypter:   encrypter,
		ProjectRepo: projectRepo,
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
	return engine
}
