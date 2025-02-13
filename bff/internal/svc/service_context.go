package svc

import (
	"time"

	"github.com/WeiXinao/msProject/account/account"
	"github.com/WeiXinao/msProject/bff/internal/config"
	"github.com/WeiXinao/msProject/bff/internal/middleware"
	"github.com/WeiXinao/msProject/file/file"
	"github.com/WeiXinao/msProject/pkg/jwtx"
	"github.com/WeiXinao/msProject/project/projectservice"
	"github.com/WeiXinao/msProject/task/taskservice"
	"github.com/WeiXinao/msProject/user/loginservice"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	StaticPath     string
	UserClient     loginservice.LoginService
	ProjectClient  projectservice.ProjectService
	TaskClient     taskservice.TaskService
	FileClient     file.File
	AccountClient account.Account
}

func NewServiceContext(c config.Config) *ServiceContext {
	jwter := InitJwter(c)
	StaticPath := c.StaticPath
	userClient := loginservice.NewLoginService(zrpc.MustNewClient(c.UserRpcClient))
	authMiddleware := middleware.NewAuthMiddlewareBuilder(jwter, userClient)
	projectClient := projectservice.NewProjectService(zrpc.MustNewClient(c.ProjectRpcClient))
	taskClient := taskservice.NewTaskService(zrpc.MustNewClient(c.TaskRpcClient))
	fileClient := file.NewFile(zrpc.MustNewClient(c.FileRpcClient))
	accountClient := account.NewAccount(zrpc.MustNewClient(c.AccountRpcClient))
	return &ServiceContext{
		Config:         c,
		UserClient:     userClient,
		ProjectClient:  projectClient,
		TaskClient:     taskClient,
		FileClient:     fileClient,
		AccountClient: accountClient,
		AuthMiddleware: authMiddleware.Build,
		StaticPath:     StaticPath,
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
