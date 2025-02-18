package main

import (
	"flag"
	"fmt"

	userv1 "github.com/WeiXinao/msProject/api/proto/gen/user/v1"
	"github.com/WeiXinao/msProject/pkg/gozerox"
	lx "github.com/WeiXinao/msProject/pkg/logx"
	"github.com/WeiXinao/msProject/user/internal/config"
	"github.com/WeiXinao/msProject/user/internal/server"
	"github.com/WeiXinao/msProject/user/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")
var distributedCentorConfigFile = flag.String("d", "etc/bootstrap.yaml", "the config file of distributed config centor")

func main() {
	flag.Parse()

	// conf.MustLoad(*configFile, &c)
	var (
		c config.Config
		nacosCfg gozerox.Cfg
	)
	conf.MustLoad(*distributedCentorConfigFile, &nacosCfg)
	cfgCentor := gozerox.MustInitNacosDistributedConfigCentor[*config.Config](nacosCfg.Nacos)
	cfgCentor.SetLocalPath(*configFile)
	cfgCentor.MustReadInConfig(&c)

	ctx := svc.NewServiceContext(c)
	cfgCentor.ListenOnConfig(func(data *config.Config) {
		ctx = svc.NewServiceContext(c)
		logx.Info("reload config...")
	})

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		userv1.RegisterLoginServiceServer(grpcServer, server.NewLoginServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(ctx.Interceptors...)
	defer s.Stop()

	// 设置 log 的 writer
	//从配置中读取日志配置，初始化日志
	writer, err := lx.NewZapWriter(&c.LogConfig)
	logx.Must(err)
	logx.SetWriter(writer)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
