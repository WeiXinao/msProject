package svc

import (
	"time"

	"github.com/WeiXinao/msProject/pkg/cachex"
	"github.com/WeiXinao/msProject/pkg/encrypts"
	"github.com/WeiXinao/msProject/pkg/jwtx"
	"github.com/WeiXinao/msProject/project/internal/config"
	"github.com/WeiXinao/msProject/project/internal/repo"
	"github.com/WeiXinao/msProject/project/internal/repo/dao"
	"github.com/WeiXinao/msProject/task/taskservice"
	"github.com/WeiXinao/msProject/user/loginservice"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config      config.Config
	Jwter       jwtx.Jwter
	Encrypter   encrypts.Encrypter
	ProjectRepo repo.ProjectRepo
	TaskClient taskservice.TaskService
	UserClient     loginservice.LoginService
}

func NewServiceContext(c config.Config) *ServiceContext {
	encrypter := encrypts.NewEncrypter(c.AESKey)
	jwter := InitJwter(c)
	rcli := redis.MustNewRedis(c.RedisConfig)
	ca := cachex.NewRedisCache(rcli)
	db := InitDB(c)

	taskClient := taskservice.NewTaskService(zrpc.MustNewClient(c.TaskRpcClient))
	userClient := loginservice.NewLoginService(zrpc.MustNewClient(c.UserRpcClient))

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
		TaskClient: taskClient,
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
