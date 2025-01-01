package config

import (
	lx "github.com/WeiXinao/msProject/pkg/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	LogConfig        lx.LogConfig
	UserRpcClient    zrpc.RpcClientConf
	ProjectRpcClient zrpc.RpcClientConf
	Jwt              Jwt
}

type Jwt struct {
	AccessExp  string
	RefreshExp string
	AtKey      string
	RtKey      string
}
