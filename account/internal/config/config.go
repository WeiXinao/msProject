package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	lx "github.com/WeiXinao/msProject/pkg/logx"
)

type Config struct {
	zrpc.RpcServerConf
	RedisConfig redis.RedisConf
	LogConfig   lx.LogConfig
	DB DB
	MySQLConfig MySQLConfig
	Jwt         Jwt
	AESKey      string
	UserRpcClient    zrpc.RpcClientConf
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

type DB struct {
	Separation bool
	Master MySQLConfig
	Slaves []MySQLConfig
}