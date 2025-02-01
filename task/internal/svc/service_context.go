package svc

import (
	"time"

	"github.com/WeiXinao/msProject/pkg/encrypts"
	"github.com/WeiXinao/msProject/pkg/jwtx"
	"github.com/WeiXinao/msProject/task/internal/config"
	"github.com/WeiXinao/msProject/task/internal/repo"
	"github.com/WeiXinao/msProject/task/internal/repo/dao"
	"github.com/WeiXinao/msProject/user/loginservice"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/zrpc"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Jwter       jwtx.Jwter
	Encrypter   encrypts.Encrypter
	TaskRepo repo.TaskRepo
	UserClient     loginservice.LoginService
}

func NewServiceContext(c config.Config) *ServiceContext {
	encrypter := encrypts.NewEncrypter(c.AESKey)
	jwter := InitJwter(c)
	// rcli := redis.MustNewRedis(c.RedisConfig)
	// ca := cachex.NewRedisCache(rcli)
	db := InitDB(c)
	dao := dao.NewTaskXormDao(db)
	taskRepo := repo.NewTaskRepo(dao)

	userClient := loginservice.NewLoginService(zrpc.MustNewClient(c.UserRpcClient))

	return &ServiceContext{
		Config:      c,
		Jwter:       jwter,
		Encrypter:   encrypter,
		TaskRepo: taskRepo,
		UserClient: userClient,
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
