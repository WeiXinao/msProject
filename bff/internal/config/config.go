package config

import (
	lx "github.com/WeiXinao/msProject/pkg/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	LogConfig        lx.LogConfig
	Jwt              Jwt
	StaticPath       string
	UserRpcClient    zrpc.RpcClientConf
	ProjectRpcClient zrpc.RpcClientConf
	TaskRpcClient    zrpc.RpcClientConf
	FileRpcClient    zrpc.RpcClientConf
	AccountRpcClient zrpc.RpcClientConf
}

type Jwt struct {
	AccessExp  string
	RefreshExp string
	AtKey      string
	RtKey      string
}
