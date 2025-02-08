package svc

import (
	"time"

	"github.com/WeiXinao/msProject/file/internal/config"
	"github.com/WeiXinao/msProject/file/internal/repo"
	"github.com/WeiXinao/msProject/file/internal/repo/dao"
	"github.com/WeiXinao/msProject/pkg/encrypts"
	"github.com/WeiXinao/msProject/pkg/jwtx"
	"xorm.io/xorm"
	_ "github.com/go-sql-driver/mysql"
)

type ServiceContext struct {
	Config config.Config
	Jwter       jwtx.Jwter
	Encrypter   encrypts.Encrypter
	FileRepo repo.FileRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	encrypter := encrypts.NewEncrypter(c.AESKey)
	jwter := InitJwter(c)
	db := InitDB(c)
	dao := dao.NewXormFileDao(db)
	fileRepo := repo.NewFileRepo(dao)

	return &ServiceContext{
		Config: c,
		Jwter:       jwter,
		Encrypter:   encrypter,
		FileRepo: fileRepo,
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
