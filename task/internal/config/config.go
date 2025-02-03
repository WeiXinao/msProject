package config

import (
	lx "github.com/WeiXinao/msProject/pkg/logx"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisConfig redis.RedisConf
	LogConfig   lx.LogConfig
	MySQLConfig MySQLConfig
	Jwt         Jwt
	AESKey      string
	UserRpcClient    zrpc.RpcClientConf
	ProjectRpcClient zrpc.RpcClientConf
	KqConsumerConf kq.KqConf
}

type Jwt struct {
	AccessExp  string
	RefreshExp string
	AtKey      string
	RtKey      string
}

type MySQLConfig struct {
	DriverName string
	Dsn        string
}

