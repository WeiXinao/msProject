package svc

import (
	"github.com/WeiXinao/msProject/bff/internal/config"
	"github.com/WeiXinao/msProject/bff/internal/middlewares"
	"github.com/WeiXinao/msProject/pkg/jwtx"
	"github.com/WeiXinao/msProject/project/projectservice"
	"github.com/WeiXinao/msProject/user/loginservice"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"time"
)

type ServiceContext struct {
	Config         config.Config
	UserClient     loginservice.LoginService
	ProjectClient  projectservice.ProjectService
	AuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	jwter := InitJwter(c)
	userClient := loginservice.NewLoginService(zrpc.MustNewClient(c.UserRpcClient))
	projectClient := projectservice.NewProjectService(zrpc.MustNewClient(c.ProjectRpcClient))
	authMiddleware := middlewares.NewAuthMiddlewareBuilder(jwter)
	return &ServiceContext{
		Config:         c,
		UserClient:     userClient,
		ProjectClient:  projectClient,
		AuthMiddleware: authMiddleware.Build,
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
