package svc

import (
	"time"

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
	UserClient     loginservice.LoginService
	ProjectClient  projectservice.ProjectService
	TaskClient     taskservice.TaskService
	FileClient     file.File
	AuthMiddleware rest.Middleware
	StaticPath     string
}

func NewServiceContext(c config.Config) *ServiceContext {
	jwter := InitJwter(c)
	userClient := loginservice.NewLoginService(zrpc.MustNewClient(c.UserRpcClient))
	projectClient := projectservice.NewProjectService(zrpc.MustNewClient(c.ProjectRpcClient))
	taskClient := taskservice.NewTaskService(zrpc.MustNewClient(c.TaskRpcClient))
	fileClient := file.NewFile(zrpc.MustNewClient(c.FileRpcClient))
	authMiddleware := middleware.NewAuthMiddlewareBuilder(jwter, userClient)
	StaticPath := c.StaticPath
	return &ServiceContext{
		Config:         c,
		UserClient:     userClient,
		ProjectClient:  projectClient,
		TaskClient:     taskClient,
		FileClient:     fileClient,
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
