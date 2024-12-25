package svc

import (
	"github.com/WeiXinao/msProject/bff/internal/config"
	"github.com/WeiXinao/msProject/user/loginservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	UserClient loginservice.LoginService
}

func NewServiceContext(c config.Config) *ServiceContext {
	userClient := loginservice.NewLoginService(zrpc.MustNewClient(c.UserRpcClient))
	return &ServiceContext{
		Config:     c,
		UserClient: userClient,
	}
}
