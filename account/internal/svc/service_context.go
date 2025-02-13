package svc

import (
	"time"

	"github.com/WeiXinao/msProject/account/internal/config"
	"github.com/WeiXinao/msProject/account/internal/repo"
	"github.com/WeiXinao/msProject/account/internal/repo/dao"
	"github.com/WeiXinao/msProject/pkg/encrypts"
	"github.com/WeiXinao/msProject/pkg/jwtx"
	"github.com/WeiXinao/msProject/user/loginservice"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/zrpc"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config    config.Config
	Jwter     jwtx.Jwter
	Encrypter encrypts.Encrypter
	AccoutRepo repo.AccountRepo
	UserClient     loginservice.LoginService
}

func NewServiceContext(c config.Config) *ServiceContext {
	encrypter := encrypts.NewEncrypter(c.AESKey)
	jwter := InitJwter(c)
	db := InitDB(c)

	userClient := loginservice.NewLoginService(zrpc.MustNewClient(c.UserRpcClient))

	accountDao := dao.NewAccountXormDao(db)
	accountRepo := repo.NewAccountRepo(accountDao)
	return &ServiceContext{
		Config:    c,
		Jwter:     jwter,
		Encrypter: encrypter,
		AccoutRepo: accountRepo,
		UserClient: userClient,	
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

func InitDB(c config.Config) *xorm.Engine {
	mySQLConfig := c.MySQLConfig
	engine, err := xorm.NewEngine(mySQLConfig.DriverName, mySQLConfig.Dsn)
	if err != nil {
		panic(err)
	}
	return engine
}
