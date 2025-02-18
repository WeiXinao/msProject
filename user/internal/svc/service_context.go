package svc

import (
	"time"

	userv1 "github.com/WeiXinao/msProject/api/proto/gen/user/v1"
	"github.com/WeiXinao/msProject/pkg/cachex"
	"github.com/WeiXinao/msProject/pkg/encrypts"
	"github.com/WeiXinao/msProject/pkg/grpcx/interceptor"
	"github.com/WeiXinao/msProject/pkg/jwtx"
	"github.com/WeiXinao/msProject/user/internal/config"
	"github.com/WeiXinao/msProject/user/internal/repo"
	"github.com/WeiXinao/msProject/user/internal/repo/cache"
	"github.com/WeiXinao/msProject/user/internal/repo/dao"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"google.golang.org/grpc"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config    config.Config
	UserRepo  repo.UserRepo
	Jwter     jwtx.Jwter
	Encrypter encrypts.Encrypter
	Interceptors []grpc.UnaryServerInterceptor
}

func NewServiceContext(c config.Config) *ServiceContext {
	encrypter := encrypts.NewEncrypter(c.AESKey)
	jwter := InitJwter(c)
	rcli := redis.MustNewRedis(c.RedisConfig)
	ca := cachex.NewRedisCache(rcli)
	db := InitDB(c)
	interceptors := InitInterceptors(ca)

	rExp, err := time.ParseDuration(c.Jwt.AccessExp)
	if err != nil {
		logx.Error(err)
		panic(err)
	}
	userCache := &cache.UserCache{}
	userDao := dao.NewXormUserDao(db)
	userRepo := repo.NewUserRepo(userDao, ca, userCache, rExp)
	return &ServiceContext{
		UserRepo:  userRepo,
		Config:    c,
		Jwter:     jwter,
		Encrypter: encrypter,
		Interceptors: interceptors,
	}
}


func InitInterceptors(cache cachex.Cache) []grpc.UnaryServerInterceptor {
	cacheInterceptor := interceptor.NewUniformCacheInterceptorBuilder(cache).
	AddPatternRespMap(userv1.LoginService_MyOrgList_FullMethodName, 
		&userv1.MyOrgListResponse{}).
	AddPatternRespMap(userv1.LoginService_MemberInfo_FullMethodName,
		&userv1.MemberInfoResponse{}).
	Build()

	return []grpc.UnaryServerInterceptor{cacheInterceptor}
}

func InitJwter(c config.Config) jwtx.Jwter {
	aExp, err := time.ParseDuration(c.Jwt.AccessExp)
	if err != nil {
		logx.Error(err)
		return nil
	}
	rExp, err := time.ParseDuration(c.Jwt.AccessExp)
	if err != nil {
		logx.Error(err)
		return nil
	}
	return jwtx.NewJwtToken(c.Jwt.AtKey, c.Jwt.RtKey, aExp, rExp)
}

func InitDB(c config.Config) *xorm.Engine {
	mySQLConfig := c.MySQLConfig
	engine, err := xorm.NewEngine(mySQLConfig.DriverName, mySQLConfig.Dsn)
	if err != nil {
		logx.Error(err)
		return nil
	}
	err = engine.Sync(
		new(dao.Member),
		new(dao.Organization),
	)
	if err != nil {
		logx.Error(err)
		return nil
	}
	return engine
}
