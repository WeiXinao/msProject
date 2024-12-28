package config

import (
	lx "github.com/WeiXinao/msProject/pkg/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	RedisConfig redis.RedisConf
	LogConfig   lx.LogConfig
	MySQLConfig MySQLConfig
}

type MySQLConfig struct {
	DriverName string
	Dsn        string
}
