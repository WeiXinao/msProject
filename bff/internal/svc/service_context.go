package svc

import (
	"time"

	"github.com/WeiXinao/msProject/account/account"
	"github.com/WeiXinao/msProject/bff/internal/config"
	"github.com/WeiXinao/msProject/bff/internal/middleware"
	"github.com/WeiXinao/msProject/file/file"
	"github.com/WeiXinao/msProject/pkg/encrypts"
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
	ProjectAuthMiddleware rest.Middleware
	StaticPath     string
	UserClient     loginservice.LoginService
	ProjectClient  projectservice.ProjectService
	TaskClient     taskservice.TaskService
	FileClient     file.File
	AccountClient account.Account
}

func NewServiceContext(c config.Config) *ServiceContext {
	jwter := InitJwter(c)
	encrypter := encrypts.NewEncrypter(c.AESKey)
	StaticPath := c.StaticPath
	userClient := loginservice.NewLoginService(zrpc.MustNewClient(c.UserRpcClient))
	accountClient := account.NewAccount(zrpc.MustNewClient(c.AccountRpcClient))
	projectClient := projectservice.NewProjectService(zrpc.MustNewClient(c.ProjectRpcClient))
	taskClient := taskservice.NewTaskService(zrpc.MustNewClient(c.TaskRpcClient))
	fileClient := file.NewFile(zrpc.MustNewClient(c.FileRpcClient))
	authMiddleware := middleware.NewAuthMiddlewareBuilder(jwter, userClient, accountClient).
		AddIngoreURLs(c.AuthorityIgnoreUrls...)
	projectAuthMiddleware := middleware.NewProjectAuthMiddleware(encrypter, projectClient, taskClient)
	return &ServiceContext{
		Config:         c,
		UserClient:     userClient,
		ProjectClient:  projectClient,
		TaskClient:     taskClient,
		FileClient:     fileClient,
		AccountClient: accountClient,
		AuthMiddleware: authMiddleware.Build,
		ProjectAuthMiddleware: projectAuthMiddleware.Handle,
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
